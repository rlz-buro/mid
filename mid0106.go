package mid

// MID 0106 Last PowerMACS tightening result Station data
// This MID contains the station part and some of the Bolt data of the last result data.
// After this message has been sent the integrator selects if it also wants to have the Bolt and step data.
// If this data is requested, then the integrator sends the message MID 0108 Last PowerMACS tightening result data acknowledge,
// with the parameter Bolt Data set to TRUE. If only the station data is wanted the parameter Bolt Data is set to FALSE.
// This telegram is also used for Power MACS systems running a Press.
// The layout of the telegram is exactly the same but some of the fields have slightly different definitions.
// The fields for Torque are used for Force values and the fields for Angle are used for Stroke values.
// Press systems also use different identifiers for the optional data on bolt and step level.
// A press system always use revision 4 or higher of the telegram
// Note: All values that are undefined in the results will be sent as all spaces (ASCII 0x20).
// This will for instance happen with the Torque Status if no measuring value for Bolt T was available for the tightening.
type MID0106REV001 struct {
	// The total number of messages needed to send all Bolt data for all Bolts.
	// The rest of the messages are of type MID 0107 Last PowerMACS tightening result Bolt data, once for each Bolt.
	// They are only sent on request from the integrator. 2 ASCII digits, range 00-99.
	TotalNoOfMessages int
	// This parameter is always 01 as this is the first message.
	MessageNumber int
	// The Data No system is a unique ID for each tightening result within the system. 10 ASCII digits, max value are 4294967295.
	DataNoSystem int
	// The station number within the PowerMACS system. 2 ASCII digits, range 01-15.
	StationNo int
	// The station name is 20 bytes long and is specified by 20 ASCII characters.
	StationName string
	// Cycle start time for each tightening sent to the control station.
	// The time is 19 byte long and is specified by 19 ASCII characters (YYYY-MM-DD:HH:MM:SS)
	Time string
	// The mode number used for the tightening. 2 ASCII digits, range 01-50. If undefined, empty spaces are sent.
	ModeNo int
	// The name of the mode used for the tightening. Specified by 20 ASCII characters. If undefined, empty spaces are sent.
	ModeName string
	// One byte long and is specified by one ASCII digit (‘0’ or ‘1’). 0=tightening NOK, 1=tightening OK.
	SimpleStatus int
	// The status of the tightening, specified by one ASCII digit. 0=OK, 1=OKR, 2=NOK, 3=TERMNOK.
	PMStatus int
	// The Wp. Id is 40 bytes long and is specified by 40 ASCII characters. If undefined, empty spaces are sent.
	WpId string
	// The total number of Bolts in the tightening, 2 ASCII digits.
	// The Bolt part in this message (indicated with double table border) is repeated Number of Bolt times.
	// The parameter numbers (13- 22) are also repeated.
	NumberOfBolts int `mid:"164-165"`
	BoltData      []struct {
		// The ordinal Bolt number, the Bolts in the station are always numbered from 01 to 50. 2 ASCII digits.
		OrdinalBoltNumber int `mid:"168-169"`
		// Specified by one ASCII digit (‘0’ or ‘1’). 0=tightening NOK, 1=tightening OK.
		SimpleBoltStatus int `mid:"172"`
		// Torque status of each Bolt, specified by one ASCII digit 0=Bolt T Low
		// 1=Bolt T OK
		// 2=Bolt T High
		// If undefined, empty spaces are sent.
		TorqueStatus int `mid:"175"`
		// Angle status of each Bolt, specified by one ASCII digit 0=Bolt A Low
		// 1=Bolt A OK
		// 2=Bolt A High
		// If undefined, empty spaces are sent.
		AngleStatus int `mid:"178"`
		// Sent as 7 ASCII digits formatted as a float.
		// The value is sent with 4 decimal places, for example 99.9999 or -9.9999.
		// If the value is larger than 99 the needed number of decimals are removed to fit the integer part,
		// i.e. 12345.123 is sent as “12345.1”.
		// The unit is Nm. If undefined, empty spaces are sent.
		BoltT float64 `mid:"181-187"`
		// Sent as 7 ASCII digits, formatted as a float, see description for Bolt T.
		// The unit is degrees. If undefined, empty spaces are sent.
		BoltA float64 `mid:"190-196"`
		// Sent as 7 ASCII digits, formatted as a float, see description for Bolt T.
		// The unit is Nm. If undefined, empty spaces are sent.
		BoltTHighLimit float64 `mid:"199-205"`
		// Sent as 7 ASCII digits, formatted as a float, see description for Bolt T.
		// The unit is Nm. If undefined, empty spaces are sent.
		BoltTLowLimit float64 `mid:"208-214"`
		// Sent as 7 ASCII digits, formatted as a float, see description for Bolt T.
		// The unit is degrees. If undefined, empty spaces are sent.
		BoltAHighLimit float64 `mid:"217-223"`
		// Sent as 7 ASCII digits, formatted as a float, see description for Bolt T.
		// The unit is degrees. If undefined, empty spaces are sent.
		BoltALowLimit float64 `mid:"226-232"`
	}

	// _fieldOrder01 uint8 // 21-22 01
	// _fieldOrder02 uint8 // 25-26 02
	// _fieldOrder03 uint8 // 29-30 03
	// _fieldOrder04 uint8 // 41-42 04
	// _fieldOrder05 uint8 // 45-46 05
	// _fieldOrder06 uint8 // 67-68 06
	// _fieldOrder07 uint8 // 88-89 07
	// _fieldOrder08 uint8 // 92-93 08
	// _fieldOrder09 uint8 // 114-115 09
	// _fieldOrder10 uint8 // 117-118 10
	// _fieldOrder11 uint8 // 120-121 11
	// _fieldOrder12 uint8 // 162-163 12
}
