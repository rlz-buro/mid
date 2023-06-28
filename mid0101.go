package mid

// MID 0101 Multi-spindle result
// The multi-spindle result is sent after each sync tightening and if it is subscribed.
// The multiple results contain the common status of the multiple as well as the individual tightening result
// (torque and angle) of each spindle.
// This telegram is also used for PowerMACS systems running a Press.
// The layout of the telegram is exactly the same but some of the fields have slightly different definitions.
// The fields for Torque are used for Force values and the fields for Angle are used for Stroke values.
// A press system always uses revision 4 or higher of the telegram.
type MID0101REV001 struct {
	// Number of spindles running in the multiple.
	// The number of spindles is two bytes long and specified by 2 ASCII digits, range 01-50.
	NumberOfSpindles int `mid:"23-24"`
	// The VIN number is 25 bytes long and is specified by 25 ASCII characters.
	VINNumber string `mid:"27-51"`
	// The Job ID is two bytes long and specified by two ASCII digits. Range: 00-99
	JobID int `mid:"54-55"`
	// The parameter set ID is three bytes long and specified by three ASCII digits. Range: 000-999.
	ParameterSetID int `mid:"58-60"`
	// This parameter gives the total number of tightening in the batch.
	// The batch size is four bytes long and specified by four ASCII digits. Range: 0000-9999.
	BatchSize int `mid:"63-66"`
	// The batch counter information is four bytes long specifying and specified by four ASCII digits. Range: 0000-9999.
	BatchCounter int `mid:"69-72"`
	// The batch status is specified by one ASCII character.
	// 0=batch NOK (batch not completed), 1=batch OK, 2=batch not used.
	BatchStatus int `mid:"75"`
	// The torque min limit in Nm is multiplied by 100 and sent as an integer (2 decimals truncated).
	// It is six bytes long and is specified by six ASCII digits.
	TorqueMinLimit int `mid:"78-83"`
	// The torque max limit in Nm is multiplied by 100 and sent as an integer (2 decimals truncated).
	// It is six bytes long and is specified by six ASCII digits.
	TorqueMaxLimit int `mid:"86-91"`
	// The torque final target in Nm is multiplied by 100 and sent as an integer (2 decimals truncated).
	// It is six bytes long and is specified by six ASCII digits.
	TorqueFinalTarget int `mid:"94-99"`
	// The angle min value in degrees. Each turn represents 360 degrees.
	// It is five bytes long and specified by five ASCII digits. Range: 00000-99999.
	AngleMin int `mid:"102-106"`
	// The angle max value in degrees. Each turn represents 360 degrees.
	// It is five bytes long and specified by five ASCII digits. Range: 00000-99999.
	AngleMax int `mid:"109-113"`
	// The target angle value in degrees. Each turn represents 360 degrees.
	// It is five bytes long and specified by five ASCII digits. Range: 00000-99999.
	FinalAngleTarget int `mid:"116-120"`
	// Time stamp for the last change in the current parameter set settings.
	// It is 19 bytes long and is specified by 19 ASCII characters (YYYY-MM-DD:HH:MM:SS).
	DateTimeOfLastChangeInParameterSetSettings string `mid:"123-141"`
	// Time stamp. 19 ASCII characters (YYYY-MM-DD:HH:MM:SS).
	TimeStamp string `mid:"144-162"`
	// The sync tightening ID is a unique ID for each sync tightening result.
	// Each individual result of each spindle is stamped with this ID.
	// The tightening ID is incremented after each sync tightening. 5 ASCII digits, range 00000-65535.
	SyncTighteningID int `mid:"165-169"`
	// The status of all the spindles.
	// OK if the individual status of each spindle is OK, NOK if at least one spindle status is NOK.
	// One ASCII digit 1=OK, 0=NOK.
	SyncOverallStatus int `mid:"172"`
	// TODO: add support of the repeated fields
	// Status of each spindel is represented as repeated sequance of ASCII charters.
	SpindelStatus []SpindelStatus `midLen:"18" midCount:"23-24"`

	// TODO: add support of the orderByte fields
	// _fieldOrder01 uint8 // 21-22 1
	// _fieldOrder02 uint8 // 25-26 2
	// _fieldOrder03 uint8 // 52-53 3
	// _fieldOrder04 uint8 // 56-57 4
	// _fieldOrder05 uint8 // 61-62 5
	// _fieldOrder06 uint8 // 67-68 6
	// _fieldOrder07 uint8 // 73-74 7
	// _fieldOrder08 uint8 // 76-77 8
	// _fieldOrder09 uint8 // 84-85 9
	// _fieldOrder10 uint8 // 92-93 10
	// _fieldOrder11 uint8 // 100-101 11
	// _fieldOrder12 uint8 // 107-108 12
	// _fieldOrder13 uint8 // 114-115 13
	// _fieldOrder14 uint8 // 121-122 14
	// _fieldOrder15 uint8 // 142-143 15
	// _fieldOrder16 uint8 // 163-164 16
	// _fieldOrder17 uint8 // 170-171 17
	// _fieldOrder18 uint8 // 173-174 18
}

type SpindelStatus struct {
	// Spindle number in the same order as in the sync list. Range 01- 99.
	SpindleNumber int `mid:"1-2"`
	// Channel ID of the spindle or press. Range 01-99
	ChannelID int `mid:"3-4"`
	// Individual overall status of the tightening of each spindle
	TighteningStatus int `mid:"5"`
	// Individual torque status of each spindle. 0=Low, 1=OK, 2 = High
	TorqueStatus int `mid:"6"`
	// The torque result of each spindle. The torque in Nm is multiplied by 100
	// and sent as an integer (2 decimals truncated). It is six bytes long and specified by six ASCII digits.
	TorqueResult int `mid:"7-12"`
	// Individual angle status of each spindle. 0=NOK, 1=OK
	AngleStatus int `mid:"13"`
	// The turning angle value in degrees for each spindle. Each turn represents 360 degrees.
	// It is five bytes long and specified by five ASCII digits. Range: 00000-99999.
	AngleValue int `mid:"14-18"`
}
