package hs1x0

import (
	"github.com/vojapet/golang-tplink-hs100/pkg/commands"
	"time"
	"github.com/pkg/errors"
	"fmt"
)

type ICommandExecutor interface {
	Execute(aQuery commands.Query) (commands.Response, error)
	GetIp() string
	GetTimeout() time.Duration
}

type Hs110 struct {
	CommandExecutor ICommandExecutor
}

func NewHs110(ip string) *Hs110 {
	return &Hs110{
		CommandExecutor: &commands.CommandExecutor{Ip: ip, Timeout: 5 * time.Second},
	}
}

func (hs110 *Hs110) GetIp() string {
	return hs110.CommandExecutor.GetIp()
}

func (hs110 *Hs110) GetTimeout() time.Duration {
	return hs110.CommandExecutor.GetTimeout()
}

func (hs110 *Hs110) execSetRelayStateQuery(query commands.Query) error {
	resp, err := hs110.CommandExecutor.Execute(query)
	if err != nil {
		return errors.Wrap(err, "Error while executing the query.")
	}

	setRelayState, err := commands.CreateSetRelayState(resp)

	if err != nil {
		return err
	}

	if setRelayState.ErrorCode != 0 {
		return errors.New(fmt.Sprintf("Device returned error code `%d`", setRelayState.ErrorCode))
	}

	return nil
}

func (hs110 *Hs110) TurnOn() error {
	return hs110.execSetRelayStateQuery(
		commands.BuildQuery(commands.Build_System(commands.Build_SetRelayState(commands.Build_SetRelayState_On()))))
}

func (hs110 *Hs110) TurnOff() error {
	return hs110.execSetRelayStateQuery(
		commands.BuildQuery(commands.Build_System(commands.Build_SetRelayState(commands.Build_SetRelayState_Off()))))
}

func (hs110 *Hs110) IsOn() (bool, error) {
	resp, err := hs110.CommandExecutor.Execute(
		commands.BuildQuery(commands.Build_System(commands.Build_GetSysInfo())))

	if err != nil {
		return false, errors.Wrap(err, "Error while executing the command.")
	}

	sysInfo, err := commands.CreateSysInfo(resp)

	if err != nil {
		return false, err
	}

	return sysInfo.RelayState == 1, nil
}

func (hs110 *Hs110) GetName() (string, error) {
	resp, err := hs110.CommandExecutor.Execute(
		commands.BuildQuery(commands.Build_System(commands.Build_GetSysInfo())))

	if err != nil {
		return "", errors.Wrap(err, "Error while executing the command.")
	}

	sysInfo, err := commands.CreateSysInfo(resp)

	if err != nil {
		return "", err
	}

	return sysInfo.Alias, nil
}

func (hs110 *Hs110) GetCurrentPowerConsumption() (PowerConsumption, error) {
	resp, err := hs110.CommandExecutor.Execute(
		commands.BuildQuery(commands.Build_Emeter(commands.Build_GetRealtime())))

	if err != nil {
		return PowerConsumption{}, errors.Wrap(err, "Error while executing the command.")
	}

	realTime, err := commands.CreateRealTime(resp)
	if err != nil {
		return PowerConsumption{}, err
	}

	if realTime.ErrorCode != 0 {
		return PowerConsumption{}, errors.New(fmt.Sprintf("Device returned error code '%d'", realTime.ErrorCode))
	}

	return PowerConsumption{
		Current: realTime.Current_mA,
		Voltage: realTime.Voltage_mV,
		Power:   realTime.Power_mW,
		TotalEnergy: realTime.Total_Wh,
	}, nil
}

func (hs110 *Hs110) GetDayStat() ([]DailyConsumption, error) {
	var result []DailyConsumption

	resp, err := hs110.CommandExecutor.Execute(
		commands.BuildQuery(commands.Build_Emeter(commands.Build_GetDayStat(2020, 11))))

	if err != nil {
		return result, errors.Wrap(err, "Error while executing the command.")
	}

	dayStat, err := commands.CreateDayStat(resp)
	if err != nil {
		return result, err
	}

	if dayStat.ErrorCode != 0 {
		return result, errors.New(fmt.Sprintf("Device returned error code '%d'", dayStat.ErrorCode))
	}

	for _, dailyConsumption := range dayStat.DayList {
		result = append(result, DailyConsumption {
			Day: dailyConsumption.Day,
			Month: dailyConsumption.Month,
			Year: dailyConsumption.Year,
			Energy_Wh: dailyConsumption.Energy_Wh,
		})
	}

	return result, nil
}

func (hs110 *Hs110) GetMonthStat() ([]MonthlyConsumption, error) {
	var result []MonthlyConsumption

	resp, err := hs110.CommandExecutor.Execute(
		commands.BuildQuery(commands.Build_Emeter(commands.Build_GetMonthStat(2020))))

	if err != nil {
		return result, errors.Wrap(err, "Error while executing the command.")
	}

	monthStat, err := commands.CreateMonthStat(resp)
	if err != nil {
		return result, err
	}

	if monthStat.ErrorCode != 0 {
		return result, errors.New(fmt.Sprintf("Device returned error code '%d'", monthStat.ErrorCode))
	}

	for _, monthlyConsumption := range monthStat.MonthList {
		result = append(result, MonthlyConsumption {
			Month: monthlyConsumption.Month,
			Year: monthlyConsumption.Year,
			Energy_Wh: monthlyConsumption.Energy_Wh,
		})
	}

	return result, nil
}

func (hs110 *Hs110) GetScanInfo() ([]AP, error) {
	result := []AP{}

	resp, err := hs110.CommandExecutor.Execute(
		commands.BuildQuery(commands.Build_Netif(commands.Build_GetScanInfo())))

	if err !=nil {
		return result, errors.Wrap(err, "Error while executing the command.")
	}

	scanInfo, err := commands.CreateScanInfo(resp)
	if err != nil {
		return result, err
	}

	if scanInfo.ErrorCode != 0 {
		return result, errors.New(fmt.Sprintf("Device returned error code '%d'", scanInfo.ErrorCode))
	}

	for _, ap := range(scanInfo.ApList) {
		result = append(result, AP{ap.SSID, ap.KeyType})
	}

	return result, nil
}

func (hs110 *Hs110) SetStaInfo(aSSID string, aPassword string) error {
	resp, err := hs110.CommandExecutor.Execute(
		commands.BuildQuery(commands.Build_Netif(commands.Build_SetStaInfo(aSSID, aPassword, 3))))

	if err !=nil {
		return errors.Wrap(err, "Error while executing the command.")
	}

	setStaInfo, err := commands.CreateSetStaInfo(resp)
	if err != nil {
		return err
	}

	if setStaInfo.ErrorCode != 0 {
		return errors.New(fmt.Sprintf("Device returned error code '%d'", setStaInfo.ErrorCode))
	}

	return nil
}
