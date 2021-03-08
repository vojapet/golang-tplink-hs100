package commands

// TP-Link Smart Home Protocol Command List
// ========================================
// (for TP-Link HS100 and HS110)
//
// System Commands
// ========================================

type System_subcommand pair

func Build_System(subQueries ...System_subcommand) pair {
	system := make(map[string]interface{})
	for _, subQuery := range subQueries {
		system[subQuery.key] = subQuery.value
	}
	return pair{key: "system", value: system}
}

// Get System Info (Software & Hardware Versions, MAC, deviceID, hwID etc.)
// {"system":{"get_sysinfo":null}}

func Build_GetSysInfo() System_subcommand {
	return System_subcommand{key: "get_sysinfo", value: EmptyValue{}}
}

// Reboot
// {"system":{"reboot":{"delay":1}}}
//
// Reset (To Factory Settings)
// {"system":{"reset":{"delay":1}}}
//
// Turn On/Off
// {"system":{"set_relay_state":{"state":1}}} / {"system":{"set_relay_state":{"state":0}}}

func Build_SetRelayState(state int) System_subcommand {
	setRelayState := make(map[string]int)
	setRelayState["state"] = state
	return System_subcommand{key: "set_relay_state", value: setRelayState}
}

func Build_SetRelayState_On() int {
	return 1
}

func Build_SetRelayState_Off() int {
	return 0
}

// Turn Off Device LED (Night mode)
// {"system":{"set_led_off":{"off":1}}}

func Build_SetLedOff(off int) System_subcommand {
	setLedOff := make(map[string]int)
	setLedOff["off"] = off
	return System_subcommand{key: "set_led_off", value: setLedOff}
}

// Set Device Alias
// {"system":{"set_dev_alias":{"alias":"supercool plug"}}}

func Build_SetDevAlias(anAlias string) System_subcommand {
	alias := make(map[string]string)
	alias["alias"] = anAlias
	return System_subcommand{key: "set_dev_alias", value: alias}
}

// Set MAC Address
// {"system":{"set_mac_addr":{"mac":"50-C7-BF-01-02-03"}}}
//
// Set Device ID
// {"system":{"set_device_id":{"deviceId":"0123456789ABCDEF0123456789ABCDEF01234567"}}}
//
// Set Hardware ID
// {"system":{"set_hw_id":{"hwId":"0123456789ABCDEF0123456789ABCDEF"}}}
//
// Set Location
// {"system":{"set_dev_location":{"longitude":6.9582814,"latitude":50.9412784}}}
//
// Perform uBoot Bootloader Check
// {"system":{"test_check_uboot":null}}
//
// Get Device Icon
// {"system":{"get_dev_icon":null}}
//
// Set Device Icon
// {"system":{"set_dev_icon":{"icon":"xxxx","hash":"ABCD"}}}
//
// Set Test Mode (command only accepted coming from IP 192.168.1.100)
// {"system":{"set_test_mode":{"enable":1}}}
//
// Download Firmware from URL
// {"system":{"download_firmware":{"url":"http://...."}}}
//
// Get Download State
// {"system":{"get_download_state":{}}}
//
// Flash Downloaded Firmware
// {"system":{"flash_firmware":{}}}
//
// Check Config
// {"system":{"check_new_config":null}}
//
//
// WLAN Commands
// ========================================

type Netif_subcommand pair

func Build_Netif(subQueries ...Netif_subcommand) pair {
	netif := make(map[string]interface{})
	for _, subQuery := range subQueries {
		netif[subQuery.key] = subQuery.value
	}
	return pair{key: "netif", value: netif}
}

// Scan for list of available APs
// {"netif":{"get_scaninfo":{"refresh":1}}}
func Build_GetScanInfo() Netif_subcommand {
	scanInfo := make(map[string]int)
	scanInfo["refresh"] = 1
	return Netif_subcommand{key: "get_scaninfo", value: scanInfo}
}


// Connect to AP with given SSID and Password
// {"netif":{"set_stainfo":{"ssid":"WiFi","password":"secret","key_type":3}}}

func Build_SetStaInfo(aSSID string, aPassword string, aKeyType int) Netif_subcommand {
	staInfo := struct {
		SSID     string `json:"ssid"`
		Password string `json:"password"`
		KeyType  int    `json:"key_type"`
	}{
		aSSID,
		aPassword,
		aKeyType,
	}

	return Netif_subcommand{key: "set_stainfo", value: staInfo}
}

//
// Cloud Commands
// ========================================
// Get Cloud Info (Server, Username, Connection Status)
// {"cnCloud":{"get_info":null}}
//
// Get Firmware List from Cloud Server
// {"cnCloud":{"get_intl_fw_list":{}}}
//
// Set Server URL
// {"cnCloud":{"set_server_url":{"server":"devs.tplinkcloud.com"}}}
//
// Connect with Cloud username & Password
// {"cnCloud":{"bind":{"username":"your@email.com", "password":"secret"}}}
//
// Unregister Device from Cloud Account
// {"cnCloud":{"unbind":null}}
//
//
// Time Commands
// ========================================

type Time_subcommand pair

func Build_Time(subQueries ...Time_subcommand) pair {
	time := make(map[string]interface{})
	for _, subQuery := range subQueries {
		time[subQuery.key] = subQuery.value
	}
	return pair{key: "time", value: time}
}

// Get Time
// {"time":{"get_time":null}}

func Build_GetTime() Time_subcommand {
	return Time_subcommand{key: "get_time", value: EmptyValue{}}
}

// Get Timezone
// {"time":{"get_timezone":null}}

func Build_GetTimezone() Time_subcommand {
	return Time_subcommand{key: "get_timezone", value: EmptyValue{}}
}

// Set Timezone
// {"time":{"set_timezone":{"year":2016,"month":1,"mday":1,"hour":10,"min":10,"sec":10,"index":42}}}
//
//
// EMeter Energy Usage Statistics Commands
// (for TP-Link HS110)
// ========================================

type Emeter_subcommand pair

func Build_Emeter(subQueries ...Emeter_subcommand) pair {
	emeter := make(map[string]interface{})
	for _, subQuery := range subQueries {
		emeter[subQuery.key] = subQuery.value
	}
	return pair{key: "emeter", value: emeter}
}

// Get Realtime Current and Voltage Reading
// {"emeter":{"get_realtime":{}}}

func Build_GetRealtime() Emeter_subcommand {
	return Emeter_subcommand{key: "get_realtime", value: EmptyValue{}}
}

// Get EMeter VGain and IGain Settings
// {"emeter":{"get_vgain_igain":{}}}

func Build_GetVGainIGain() Emeter_subcommand {
	return Emeter_subcommand{key: "get_vgain_igain", value: EmptyValue{}}
}

// Set EMeter VGain and Igain
// {"emeter":{"set_vgain_igain":{"vgain":13462,"igain":16835}}}
//
// Start EMeter Calibration
// {"emeter":{"start_calibration":{"vtarget":13462,"itarget":16835}}}
//
// Get Daily Statistic for given Month
// {"emeter":{"get_daystat":{"month":1,"year":2016}}}

func Build_GetDayStat(year int, month int) Emeter_subcommand {
	getDayStat := make(map[string]int)
	getDayStat["year"] = year
	getDayStat["month"] = month
	return Emeter_subcommand{key: "get_daystat", value: getDayStat}
}

// Get Montly Statistic for given Year
// {"emeter":{"get_monthstat":{"year":2016}}}

func Build_GetMonthStat(year int) Emeter_subcommand {
	getMonthStat := make(map[string]int)
	getMonthStat["year"] = year
	return Emeter_subcommand{key: "get_monthstat", value: getMonthStat}
}

// Erase All EMeter Statistics
// {"emeter":{"erase_emeter_stat":null}}

func Build_EraseEmeterStat() Emeter_subcommand {
	return Emeter_subcommand{key: "erase_emeter_stat", value: EmptyValue{}}
}

//
// Schedule Commands
// (action to perform regularly on given weekdays)
// ========================================
// Get Next Scheduled Action
// {"schedule":{"get_next_action":null}}
//
// Get Schedule Rules List
// {"schedule":{"get_rules":null}}
//
// Add New Schedule Rule
// {"schedule":{"add_rule":{"stime_opt":0,"wday":[1,0,0,1,1,0,0],"smin":1014,"enable":1,"repeat":1,"etime_opt":-1,"name":"lights on","eact":-1,"month":0,"sact":1,"year":0,"longitude":0,"day":0,"force":0,"latitude":0,"emin":0},"set_overall_enable":{"enable":1}}}
//
// Edit Schedule Rule with given ID
// {"schedule":{"edit_rule":{"stime_opt":0,"wday":[1,0,0,1,1,0,0],"smin":1014,"enable":1,"repeat":1,"etime_opt":-1,"id":"4B44932DFC09780B554A740BC1798CBC","name":"lights on","eact":-1,"month":0,"sact":1,"year":0,"longitude":0,"day":0,"force":0,"latitude":0,"emin":0}}}
//
// Delete Schedule Rule with given ID
// {"schedule":{"delete_rule":{"id":"4B44932DFC09780B554A740BC1798CBC"}}}
//
// Delete All Schedule Rules and Erase Statistics
// {"schedule":{"delete_all_rules":null,"erase_runtime_stat":null}}
//
//
// Countdown Rule Commands
// (action to perform after number of seconds)
// ========================================
// Get Rule (only one allowed)
// {"count_down":{"get_rules":null}}
//
// Add New Countdown Rule
// {"count_down":{"add_rule":{"enable":1,"delay":1800,"act":1,"name":"turn on"}}}
//
// Edit Countdown Rule with given ID
// {"count_down":{"edit_rule":{"enable":1,"id":"7C90311A1CD3227F25C6001D88F7FC13","delay":1800,"act":1,"name":"turn on"}}}
//
// Delete Countdown Rule with given ID
// {"count_down":{"delete_rule":{"id":"7C90311A1CD3227F25C6001D88F7FC13"}}}
//
// Delete All Coundown Rules
// {"count_down":{"delete_all_rules":null}}
//
//
// Anti-Theft Rule Commands (aka Away Mode)
// (period of time during which device will be randomly turned on and off to deter thieves)
// ========================================
// Get Anti-Theft Rules List
// {"anti_theft":{"get_rules":null}}
//
// Add New Anti-Theft Rule
// {"anti_theft":{"add_rule":{"stime_opt":0,"wday":[0,0,0,1,0,1,0],"smin":987,"enable":1,"frequency":5,"repeat":1,"etime_opt":0,"duration":2,"name":"test","lastfor":1,"month":0,"year":0,"longitude":0,"day":0,"latitude":0,"force":0,"emin":1047},"set_overall_enable":1}}
//
// Edit Anti-Theft Rule with given ID
// {"anti_theft":{"edit_rule":{"stime_opt":0,"wday":[0,0,0,1,0,1,0],"smin":987,"enable":1,"frequency":5,"repeat":1,"etime_opt":0,"id":"E36B1F4466B135C1FD481F0B4BFC9C30","duration":2,"name":"test","lastfor":1,"month":0,"year":0,"longitude":0,"day":0,"latitude":0,"force":0,"emin":1047},"set_overall_enable":1}}
//
// Delete Anti-Theft Rule with given ID
// {"anti_theft":{"delete_rule":{"id":"E36B1F4466B135C1FD481F0B4BFC9C30"}}}
//
// Delete All Anti-Theft Rules
// "anti_theft":{"delete_all_rules":null}}
