package mid

// This message contains the cycle data for one Bolt, both Bolt data and step data. It is only sent if the acknowledgement of
// the message MID 0106 Last PowerMACS tightening result station data had the parameter Bolt Data set to TRUE.
// The next Bolt data is sent if the acknowledgement has the parameter Bolt Data set to TRUE.
// This telegram is also used for Power MACS systems running a Press.
// The layout of the telegram is exactly the same but some of the fields have slightly different definitions.
// The fields for Torque are used for Force values and the fields for Angle are used for Stroke values.
// Press systems also use different identifiers for the optional data on bolt and step level.
// Press systems always use revision 4 or higher of the telegram.
// Values in the fixed part that are undefined in the results will be sent as all spaces (ASCII 0x20).
// This can happen with the Customer Error Code if this function is not activated.
// Note 2: The Bolt results and step results are only sent when the value exists in the result.
// This means, for example, that if no high limit is programmed for Peak T,
// then the value Peak T + will not be sent even if limits for Peak T are defined in the reporter.
type MID0107REV001 struct {
	// The total number of messages needed to send all Bolt data for all Bolts,
	// including the message MID 0106 Last Power MACS tightening result Station data, sent with the station data.
	// One message MID 0107 Last Power MACS tightening result Bolt data is sent for each Bolt.
	TotalNoOfMessages int `mid:"23-24" midPos:"01"`
	// This number counts from 02 to Total no of messages and is incremented by 1 for each sent message.
	// The first Bolt message is message number 02, since MID 0106 Last Power MACS tightening result Station data is number 01.
	// 2 ASCII digits, range 02-99.
	MessageNumber int `mid:"27-28" midPos:"02"`
	// The Data No system is a unique ID for each tightening result within the system. 10 ASCII digits, max value are 4294967295.
	DataNoSystem int `mid:"31-40" midPos:"03"`
	// The station number within the Power MACS system. 2 ASCII digits. Range 01-15.
	StationNo int `mid:"43-44" midPos:"04"`
	// Cycle start time for each tightening sent to the control station.
	// The time is 19 byte long and is specified by 19 ASCII characters (YYYY-MM-DD:HH:MM:SS)
	Time string `mid:"47-55" midPos:"05"`
	// The user defined Bolt number. 4 ASCII digits, range 0001-9999.
	BoltNumber int `mid:"58-61" midPos:"06"`
	// The name of the Bolt. 20 ASCII characters.
	BoltName string `mid:"64-83" midPos:"07"`
	// The name of the program that made the tightening, 20 ASCII characters.
	ProgramName string `mid:"86-105" midPos:"08"`
	// The status of the tightening specified by one ASCII digit.
	// 0=OK, 1=OKR, 2=NOK, 3=TERMNOK.
	PMStatus int `mid:"108" midPos:"09"`
	// Error codes from the tightening. Formatted in the same way as the E1 special variable
	Errors string `mid:"111-160" midPos:"10"`
	// Customer specific error code. 4 ASCII characters. If undefined, empty spaces are sent.
	CustomerErrorCode string `mid:"163-166" midPos:"11"`

	// TODO: add other fields
}
