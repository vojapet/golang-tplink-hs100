package commands

import (

)

type Envelope struct {
	System *System `mapstructure:"system"`
	Emeter *Emeter `mapstructure:"emeter"`
}


type Emeter struct {
	RealTime  *RealTime  `mapstructure:"get_realtime"`
	DayStat   *DayStat   `mapstructure:"get_daystat"`
	MonthStat *MonthStat `mapstructure:"get_monthstat"`
}

type RealTime struct {
	Voltage_mV int `mapstructure:"voltage_mv"`
	Current_mA int `mapstructure:"current_ma"`
	Power_mW   int `mapstructure:"power_mw"`
	Total_Wh   int `mapstructure:"total_wh"`
	ErrorCode  int `mapstructure:"err_code"`
}

type DailyConsumption struct {
	Day       int `mapstructure:"day"`
	Month     int `mapstructure:"month"`
	Year      int `mapstructure:"year"`
	Energy_Wh int `mapstructure:"energy_wh"`
}

type MonthlyConsumption struct {
	Month     int `mapstructure:"month"`
	Year      int `mapstructure:"year"`
	Energy_Wh int `mapstructure:"energy_wh"`
}

type DayStat struct {
	DayList   []DailyConsumption `mapstructure:"day_list"`
	ErrorCode int                `mapstructure:"error_code"`
}

type MonthStat struct {
	MonthList []MonthlyConsumption `mapstructure:"month_list"`
	ErrorCode int                  `mapstructure:"error_code"`
}

type System struct {
	SysInfo       *SysInfo       `mapstructure:"get_sysinfo"`
	SetRelayState *SetRelayState `mapstructure:"set_relay_state"`
}

type SysInfo struct {
	SW_Version string     `mapstructure:"sw_ver"`
	HW_Version string     `mapstructure:"hw_ver"`
	DeviceId   string     `mapstructure:"deviceId"`
	OEMId      string     `mapstructure:"oemId"`
	HW_ID      string     `mapstructure:"hwId"`
	RSSI       int        `mapstructure:"rssi"`
	Longitude  int        `mapstructure:"longitude_i"`
	Latitude   int        `mapstructure:"latitude_i"`
	Alias      string     `mapstructure:"alias"`
	Status     string     `mapstructure:"status"`
	MicType    string     `mapstructure:"mic_type"`
	Feature    string     `mapstructure:"feature"`
	MAC        string     `mapstructure:"mac"`
	Updating   int        `mapstructure:"updating"`
	LedOff     int        `mapstructure:"led_off"`
	RelayState int        `mapstructure:"relay_state"`
	OnTime     int        `mapstructure:"on_time"`
	ActiveMode string     `mapstructure:"active_mode"`
	IconHash   string     `mapstructure:"icon_hash"`
	DevName    string     `mapstructure:"dev_name"`
	NextAction NextAction `mapstructure:"next_action"`
	ErrorCode  int        `mapstructure:"err_code"`
}

type NextAction struct {
	Type int `mapstructure:"type"`
}

type SetRelayState struct {
	ErrorCode int `mapstructure:"err_code"`
}

