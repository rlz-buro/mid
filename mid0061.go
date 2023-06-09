package mid

// MID 0061 Last tightening result data
// Upload the last tightening result.
type MID0061REV001 struct {
	// 21-22 01
	// The cell ID is four bytes long and specified by four ASCII digits. Range: 0000-9999.
	CellID int `mid:"23-26"`
	// 27-28 02
	// The channel ID is two bytes long and specified by two ASCII digits. Range: 00-99.
	ChannelID int `mid:"29-30"`
	// 31-32 03
	// The controller name is 25 bytes long and is specified by 25 ASCII characters.
	TorqueControllerName string `mid:"33-57"`
	// 58-59 04
	// The VIN number is 25 bytes long and is specified by 25 ASCII characters.
	VINNumber string `mid:"60-84"`
	// 85-86 05
	// The Job ID is two bytes long and specified by two ASCII digits. Range: 00-99
	JobID int `mid:"87-88"`
	// 89-90 06
	// The parameter set ID is three bytes long and specified by three ASCII digits. Range: 000-999.
	ParameterSetID int `mid:"91-93"`
	// 94-95 07
	// This parameter gives the total number of tightening in the batch.
	// The batch size is four bytes long and specified by four ASCII digits. Range: 0000-9999.
	BatchSize int `mid:"96-99"`
	// 100-101 08
	// The batch counter information is four bytes long specifying and specified by four ASCII digits. Range: 0000-9999.
	BatchCounter int `mid:"102-105"`
	// 106-107 09
	// The tightening status is one byte long and specified by one ASCII digit. 0=tightening NOK, 1=tightening OK.
	TighteningStatus int `mid:"108"`
	// 109-110 10
	// 0=Low, 1=OK, 2=High
	TorqueStatus int `mid:"111"`
	// 112-113 11
	// 0=Low, 1=OK, 2=High
	AngleStatus int `mid:"114"`
	// 115-116 12
	// The torque min limit is multiplied by 100 and sent as an integer (2 decimals truncated).
	// It is six bytes long and is specified by six ASCII digits.
	TorqueMinLimit int `mid:"117-122"`
	// 123-124 13
	// The torque max limit is multiplied by 100 and sent as an integer (2 decimals truncated).
	// It is six bytes long and is specified by six ASCII digits.
	TorqueMaxLimit int `mid:"125-130"`
	// 31-132 14
	// The torque final target is multiplied by 100 and sent as an integer (2 decimals truncated).
	// It is six bytes long and is specified by six ASCII digits.
	TorqueFinalTarget int `mid:"133-138"`
	// 139-140 15
	// The torque value is multiplied by 100 and sent as an integer (2 decimals truncated).
	// It is six bytes long and is specified by six ASCII digits.
	Torque int `mid:"141-146"`
	// 147-148 16
	// The angle min value in degrees. Each turn represents 360 degrees.
	// It is five bytes long and specified by five ASCII digits. Range: 00000-99999.
	AngleMin int `mid:"149-153"`
	// 154-155 17
	// The angle max value in degrees. Each turn represents 360 degrees.
	// It is five bytes long and specified by five ASCII digits. Range: 00000-99999.
	AngleMax int `mid:"156-160"`
	// 161-162 18
	// The target angle value in degrees. Each turn represents 360 degrees.
	// It is five bytes long and specified by five ASCII digits. Range: 00000-99999.
	FinalAngleTarget int `mid:"163-167"`
	// 168-169 19
	// The turning angle value in degrees. Each turn represents 360 degrees.
	// It is five bytes long and specified by five ASCII digits. Range: 00000-99999.
	Angle int `mid:"170-174"`
	// 175-176 20
	// Time stamp for each tightening.
	// It is 19 bytes long and is specified by 19 ASCII characters (YYYY-MM-DD:HH:MM:SS).
	TimeStamp string `mid:"177-195"`
	// 196-197 21
	// Time stamp for the last change in the current parameter set settings.
	// It is 19 bytes long and is specified by 19 ASCII characters (YYYY-MM- DD:HH:MM:SS).
	DateTimeOfLastChangeInParameterSetSettings string `mid:"198-216"`
	// 217-218 22
	// The batch status is specified by one ASCII character.
	// 0=batch NOK, 1=batch OK, 2=batch not used, 3=batch running
	BatchStatus int `mid:"219"`
	// 220-221 23
	//The tightening ID is a unique ID for each tightening result.
	// It is incremented after each tightening. 10 ASCII digits. Max 4294967295
	TighteningID int `mid:"222-231"`
}
