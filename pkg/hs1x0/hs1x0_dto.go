package hs1x0


type PowerConsumption struct {
	Current     int
	Voltage     int
	Power       int
	TotalEnergy int
}

type DailyConsumption struct {
	Day       int
	Month     int
	Year      int
	Energy_Wh int
}

type MonthlyConsumption struct {
	Month     int
	Year      int
	Energy_Wh int
}
