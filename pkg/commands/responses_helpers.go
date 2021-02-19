package commands

import (
	"errors"
	"github.com/mitchellh/mapstructure"
)

func DecodeResponse(response Response) (Envelope, error) {
	envelope := Envelope{}
	err := mapstructure.Decode(response, &envelope)
	return envelope, err
}

func PickSystem(envelope Envelope) (System, error) {
	if envelope.System != nil {
		return *envelope.System, nil
	}
	return System{}, errors.New("System is nil.")
}

func PickSetRelayState(system System) (SetRelayState, error) {
	if system.SetRelayState != nil {
		return *system.SetRelayState, nil
	}
	return SetRelayState{}, errors.New("SetRelayState is nil")
}

func PickSysInfo(system System) (SysInfo, error) {
	if system.SysInfo != nil {
		return *system.SysInfo, nil
	}
	return SysInfo{}, errors.New("SysInfo is nil")
}

func PickEmeter(envelope Envelope) (Emeter, error) {
	if envelope.Emeter != nil {
		return *envelope.Emeter, nil
	}
	return Emeter{}, errors.New("Emeter is nil")
}

func PickDayStat(emeter Emeter) (DayStat, error) {
	if emeter.DayStat != nil {
		return *emeter.DayStat, nil
	}
	return DayStat{}, nil
}

func PickMonthStat(emeter Emeter) (MonthStat, error) {
	if emeter.MonthStat != nil {
		return *emeter.MonthStat, nil
	}
	return MonthStat{}, nil
}

func PickRealTime(emeter Emeter) (RealTime, error) {
	if emeter.RealTime != nil {
		return *emeter.RealTime, nil
	}
	return RealTime{}, errors.New("Emeter is nil")
}

func PickNetif(envelope Envelope) (Netif, error) {
	if envelope.Netif != nil {
		return *envelope.Netif, nil
	}
	return Netif{}, errors.New("Netif is nil")
}

func PickScanInfo(netif Netif) (ScanInfo, error) {
	if netif.ScanInfo != nil {
		return *netif.ScanInfo, nil
	}
	return ScanInfo{}, errors.New("ScanInfo is nil")
}

func PickSetStaInfo(netif Netif) (SetStaInfo, error) {
	if netif.SetStaInfo != nil {
		return *netif.SetStaInfo, nil
	}
	return SetStaInfo{}, errors.New("SetStaInfo is nil")
}

func CreateSetRelayState(response Response) (SetRelayState, error) {
	envelope, err := DecodeResponse(response)
	if err != nil {
		return SetRelayState{}, err
	}

	system, err := PickSystem(envelope)
	if err != nil {
		return SetRelayState{}, err
	}

	setRelayState, err := PickSetRelayState(system)
	if err != nil {
		return SetRelayState{}, err
	}

	return setRelayState, nil
}

func CreateSysInfo(response Response) (SysInfo, error) {
	envelope, err := DecodeResponse(response)
	if err != nil {
		return SysInfo{}, err
	}

	system, err := PickSystem(envelope)
	if err != nil {
		return SysInfo{}, err
	}

	sysInfo, err := PickSysInfo(system)
	if err != nil {
		return SysInfo{}, err
	}

	return sysInfo, nil
}

func CreateRealTime(response Response) (RealTime, error) {
	envelope, err := DecodeResponse(response)
	if err != nil {
		return RealTime{}, err
	}

	emeter, err := PickEmeter(envelope)
	if err != nil {
		return RealTime{}, err
	}

	realTime, err := PickRealTime(emeter)
	if err != nil {
		return RealTime{}, err
	}

	return realTime, nil
}

func CreateDayStat(response Response) (DayStat, error) {
	envelope, err := DecodeResponse(response)
	if err != nil {
		return DayStat{}, err
	}

	emeter, err := PickEmeter(envelope)
	if err != nil {
		return DayStat{}, err
	}

	dayStat, err := PickDayStat(emeter)
	if err != nil {
		return DayStat{}, err
	}

	return dayStat, nil
}

func CreateMonthStat(response Response) (MonthStat, error) {
	envelope, err := DecodeResponse(response)
	if err != nil {
		return MonthStat{}, err
	}

	emeter, err := PickEmeter(envelope)
	if err != nil {
		return MonthStat{}, err
	}

	monthStat, err := PickMonthStat(emeter)
	if err != nil {
		return MonthStat{}, err
	}

	return monthStat, nil
}

func CreateScanInfo(response Response) (ScanInfo, error) {
	envelope, err := DecodeResponse(response)
	if err != nil {
		return ScanInfo{}, err
	}

	netif, err := PickNetif(envelope)
	if err != nil {
		return ScanInfo{}, err
	}

	scanInfo, err := PickScanInfo(netif)
	if err != nil {
		return ScanInfo{}, err
	}

	return scanInfo, nil
}

func CreateSetStaInfo(response Response) (SetStaInfo, error) {
	envelope, err := DecodeResponse(response)
	if err != nil {
		return SetStaInfo{}, err
	}

	netif, err := PickNetif(envelope)
	if err != nil {
		return SetStaInfo{}, err
	}

	setStaInfo, err := PickSetStaInfo(netif)
	if err != nil {
		return SetStaInfo{}, err
	}

	return setStaInfo, nil
}
