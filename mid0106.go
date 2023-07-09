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
	TotalNoOfMessages int `mid:"23-24" midPos:"01"`
	// This parameter is always 01 as this is the first message.
	MessageNumber int `mid:"27-28" midPos:"02"`
	// The Data No system is a unique ID for each tightening result within the system. 10 ASCII digits, max value are 4294967295.
	DataNoSystem int `mid:"31-40" midPos:"03"`
	// The station number within the PowerMACS system. 2 ASCII digits, range 01-15.
	StationNo int `mid:"43-44" midPos:"04"`
	// The station name is 20 bytes long and is specified by 20 ASCII characters.
	StationName string `mid:"47-66" midPos:"05"`
	// Cycle start time for each tightening sent to the control station.
	// The time is 19 byte long and is specified by 19 ASCII characters (YYYY-MM-DD:HH:MM:SS)
	Time string `mid:"69-87" midPos:"06"`
	// The mode number used for the tightening. 2 ASCII digits, range 01-50. If undefined, empty spaces are sent.
	ModeNo int `mid:"90-91" midPos:"07"`
	// The name of the mode used for the tightening. Specified by 20 ASCII characters. If undefined, empty spaces are sent.
	ModeName string `mid:"94-113" midPos:"08"`
	// One byte long and is specified by one ASCII digit (‘0’ or ‘1’). 0=tightening NOK, 1=tightening OK.
	SimpleStatus int `mid:"116" midPos:"09"`
	// The status of the tightening, specified by one ASCII digit. 0=OK, 1=OKR, 2=NOK, 3=TERMNOK.
	PMStatus int `mid:"119" midPos:"10"`
	// The Wp. Id is 40 bytes long and is specified by 40 ASCII characters. If undefined, empty spaces are sent.
	WpId string `mid:"122-161" midPos:"11"`
	// The total number of Bolts in the tightening, 2 ASCII digits.
	// The Bolt part in this message (indicated with double table border) is repeated Number of Bolt times.
	// The parameter numbers (13- 22) are also repeated.
	NumberOfBolts int        `mid:"164-165" midPos:"12"`
	BoltData      []BoltData `midCount:"164-165"`
	// TODO: add special values
}

type BoltData struct {
	// The ordinal Bolt number, the Bolts in the station are always numbered from 01 to 50. 2 ASCII digits.
	OrdinalBoltNumber int `mid:"+2" midPos:"13"`
	// Specified by one ASCII digit (‘0’ or ‘1’). 0=tightening NOK, 1=tightening OK.
	SimpleBoltStatus int `mid:"+1" midPos:"14"`
	// Torque status of each Bolt, specified by one ASCII digit 0=Bolt T Low
	// 1=Bolt T OK
	// 2=Bolt T High
	// If undefined, empty spaces are sent.
	TorqueStatus int `mid:"+1" midPos:"15"`
	// Angle status of each Bolt, specified by one ASCII digit 0=Bolt A Low
	// 1=Bolt A OK
	// 2=Bolt A High
	// If undefined, empty spaces are sent.
	AngleStatus int `mid:"+1" midPos:"16"`
	// Sent as 7 ASCII digits formatted as a float.
	// The value is sent with 4 decimal places, for example 99.9999 or -9.9999.
	// If the value is larger than 99 the needed number of decimals are removed to fit the integer part,
	// i.e. 12345.123 is sent as “12345.1”.
	// The unit is Nm. If undefined, empty spaces are sent.
	BoltT float64 `mid:"+7" midPos:"17"`
	// Sent as 7 ASCII digits, formatted as a float, see description for Bolt T.
	// The unit is degrees. If undefined, empty spaces are sent.
	BoltA float64 `mid:"+7" midPos:"18"`
	// Sent as 7 ASCII digits, formatted as a float, see description for Bolt T.
	// The unit is Nm. If undefined, empty spaces are sent.
	BoltTHighLimit float64 `mid:"+7" midPos:"19"`
	// Sent as 7 ASCII digits, formatted as a float, see description for Bolt T.
	// The unit is Nm. If undefined, empty spaces are sent.
	BoltTLowLimit float64 `mid:"+7" midPos:"20"`
	// Sent as 7 ASCII digits, formatted as a float, see description for Bolt T.
	// The unit is degrees. If undefined, empty spaces are sent.
	BoltAHighLimit float64 `mid:"+7" midPos:"21"`
	// Sent as 7 ASCII digits, formatted as a float, see description for Bolt T.
	// The unit is degrees. If undefined, empty spaces are sent.
	BoltALowLimit float64 `mid:"+7" midPos:"22"`
}
