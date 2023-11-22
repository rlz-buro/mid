package mid

// Old tightening upload.
type MID0065REV001 struct {
	// The tightening ID is a unique ID for each tighteningresult.
	// It is incremented after each tightening. 10 ASCII digits. Max 4294967295
	TighteningID int `mid:"23-32" midPos:"01"`
	// The VIN number is 25 bytes long and is specified by 25 ASCII characters.
	VINNumber string `mid:"35-59" midPos:"02"`
	// The parameter set ID is three bytes long and specified by three ASCII digits. Range: 000-999.
	ParameterSetID int `mid:"62-64" midPos:"03"`
	// The batch counter information is four bytes long specifying and specified by four ASCII digits. Range: 0000-9999.
	BatchCounter int `mid:"67-70" midPos:"04"`
	// The tightening status is one byte long and specified by one ASCII digit. 0=tightening NOK, 1=tightening OK.
	TighteningStatus int `mid:"73" midPos:"05"`
	// 0=Low, 1=OK, 2=High
	TorqueStatus int `mid:"76" midPos:"06"`
	// 0=Low, 1=OK, 2=High
	AngleStatus int `mid:"79" midPos:"07"`
	// The torque value is multiplied by 100 and sent as an integer (2 decimals truncated).
	// It is six bytes long and is specified by six ASCII digits.
	Torque int `mid:"82-87" midPos:"8"`
	// The turning angle value in degrees. Each turn represents 360 degrees.
	// It is five bytes long and specified by five ASCII digits. Range: 00000-99999.
	Angle int `mid:"90-94" midPos:"9"`
	// Time stamp for each tightening.
	// It is 19 bytes long and is specified by 19 ASCII characters (YYYY-MM-DD:HH:MM:SS).
	TimeStamp string `mid:"97-115" midPos:"10"`
	// The batch status is specified by one ASCII character.
	// 0=batch NOK, 1=batch OK, 2=batch not used, 3=batch running
	BatchStatus int `mid:"118" midPos:"11"`
}
