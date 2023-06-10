package mid

import "fmt"

type ErrorCode int

const (
	NoError                                                                              ErrorCode = 0  // No Error
	InvalidData                                                                          ErrorCode = 1  // Invalid data
	ParameterSetIDNotPresent                                                             ErrorCode = 2  // Parameter set ID not present
	ParameterSetCanNotBeSet                                                              ErrorCode = 3  // Parameter set can not be set.
	ParameterSetNotRunning                                                               ErrorCode = 4  // Parameter set not running
	VINUploadSubscriptionAlreadyExists                                                   ErrorCode = 6  // VIN upload subscription already exists
	VINUploadSubscriptionDoesNotExists                                                   ErrorCode = 7  // VIN upload subscription does not exists
	VINInputSourceNotGranted                                                             ErrorCode = 8  // VIN input source not granted
	LastTighteningResultSubscriptionAlreadyExists                                        ErrorCode = 9  // Last tightening result subscription already exists
	LastTighteningResultSubscriptionDoesNotExist                                         ErrorCode = 10 // Last tightening result subscription does not exist
	AlarmSubscriptionAlreadyExists                                                       ErrorCode = 11 // Alarm subscription already exists
	AlarmSubscriptionDoesNotExist                                                        ErrorCode = 12 // Alarm subscription does not exist
	ParameterSetSelectionSubscriptionAlreadyExists                                       ErrorCode = 13 // Parameter set selection subscription already exists
	ParameterSetSelectionSubscriptionDoesNotExist                                        ErrorCode = 14 // Parameter set selection subscription does not exist
	TighteningIDRequestedNotFound                                                        ErrorCode = 15 // Tightening ID requested not found
	ConnectionRejectedProtocolBusy                                                       ErrorCode = 16 // Connection rejected protocol busy
	JobIDNotPresent                                                                      ErrorCode = 17 // Job ID not present
	JobInfoSubscriptionAlreadyExists                                                     ErrorCode = 18 // Job info subscription already exists
	JobInfoSubscriptionDoesNotExist                                                      ErrorCode = 19 // Job info subscription does not exist
	JobCanNotBeSet                                                                       ErrorCode = 20 // Job can not be set
	JobNotRunning                                                                        ErrorCode = 21 // Job not running
	NotPossibleToExecuteDynamicJobRequest                                                ErrorCode = 22 // Not possible to execute dynamic Job request
	JobBatchDecrementFailed                                                              ErrorCode = 23 // Job batch decrement failed
	NotPossibleToCreatePset                                                              ErrorCode = 24 // Not possible to create Pset
	ProgrammingControlNotGranted                                                         ErrorCode = 25 // Programming control not granted
	WrongToolTypeToPsetDownloadConnected                                                 ErrorCode = 26 // Wrong tool type to Pset download connected
	ToolIsInaccessible                                                                   ErrorCode = 27 // Tool is inaccessible
	JobAbortionIsInProgress                                                              ErrorCode = 28 // Job abortion is in progress
	ToolDoesNotExist                                                                     ErrorCode = 29 // Tool does not exist
	ControllerIsNotASyncMasterStationController                                          ErrorCode = 30 // Controller is not a sync Master/station controller
	MultiSpindleStatusSubscriptionAlreadyExists                                          ErrorCode = 31 // Multi-spindle status subscription already exists
	MultiSpindleStatusSubscriptionDoesNotExist                                           ErrorCode = 32 // Multi-spindle status subscription does not exist
	MultiSpindleResultSubscriptionAlreadyExists                                          ErrorCode = 33 // Multi-spindle result subscription already exists
	MultiSpindleResultSubscriptionDoesNotExist                                           ErrorCode = 34 // Multi-spindle result subscription does not exist
	OtherMasterClientAlreadyConnected                                                    ErrorCode = 35 // Other master client already connected
	LockTypeNotSupported                                                                 ErrorCode = 36 // Lock type not supported
	JobLineControlInfoSubscriptionAlreadyExists                                          ErrorCode = 40 // Job line control info subscription already exists
	JobLineControlInfoSubscriptionDoesNotExist                                           ErrorCode = 41 // Job line control info subscription does not exist
	IdentifierInputSourceNotGranted                                                      ErrorCode = 42 // Identifier input source not granted
	MultipleIdentifiersWorkOrderSubscriptionAlreadyExists                                ErrorCode = 43 // Multiple identifiers work order subscription already exists
	MultipleIdentifiersWorkOrderSubscriptionDoesNotExist                                 ErrorCode = 44 // Multiple identifiers work order subscription does not exist
	StatusExternalMonitoredInputsSubscriptionAlreadyExists                               ErrorCode = 50 // Status external monitored inputs subscription already exists
	StatusExternalMonitoredInputsSubscriptionDoesNotExist                                ErrorCode = 51 // Status external monitored inputs subscription does not exist
	IODeviceNotConnected                                                                 ErrorCode = 52 // IO device not connected
	FaultyIODeviceID                                                                     ErrorCode = 53 // Faulty IO device ID
	ToolTagIDUnknown                                                                     ErrorCode = 54 // Tool Tag ID unknown
	ToolTagIDSubscriptionAlreadyExists                                                   ErrorCode = 55 // Tool Tag ID subscription already exists
	ToolTagISubscriptionDoesNotExist                                                     ErrorCode = 56 // Tool Tag ID subscription does not exist
	ToolMotorTuningFailed                                                                ErrorCode = 57 // Tool Motor tuning failed
	NoAlarmPresent                                                                       ErrorCode = 58 // No alarm present
	ToolCurrentlyInUse                                                                   ErrorCode = 59 // Tool currently in use
	NoHistogramAvailable                                                                 ErrorCode = 60 // No histogram available
	PairingFailed                                                                        ErrorCode = 61 // Pairing failed
	PairingDenied                                                                        ErrorCode = 62 // Pairing denied
	PairingOrPairingAbortionAttemptOnWrongTooltype                                       ErrorCode = 63 // Pairing or Pairing abortion attempt on wrong tooltype
	PairingAbortionDenied                                                                ErrorCode = 64 // Pairing abortion denied
	PairingDisconnectionFailed                                                           ErrorCode = 66 // Pairing disconnection failed
	PairingAbortionFailed                                                                ErrorCode = 65 // Pairing abortion failed
	PairingInProgressOrAlreadyDone                                                       ErrorCode = 67 // Pairing in progress or already done
	PairingDeniedNoProgramControl                                                        ErrorCode = 68 // Pairing denied. No Program Control
	UnsupportedExtraDataRevision                                                         ErrorCode = 69 // Unsupported extra data revision
	CalibrationFailed                                                                    ErrorCode = 70 // Calibration failed
	SubscriptionAlreadyExists                                                            ErrorCode = 71 // Subscription already exists
	SubscriptionDoesNotExists                                                            ErrorCode = 72 // Subscription does not exists
	SubscribedMIDUnsupportedAnswerIfTryingToSubscribeOnANonExistingMID                   ErrorCode = 73 // Subscribed MID unsupported, -answer if trying to subscribe on a non-existing MID
	SubscribedMIDRevisionUnsupportedAnswerIfTryingToSubscribeOnUnsupportedMIDRevision    ErrorCode = 74 // Subscribed MID Revision unsupported,-answer if trying to subscribe on unsupported MID Revision.
	RequestedMIDRevisionUnsupportedResponseWhenTryingToRequestUnsupportedMIDRevision     ErrorCode = 76 // Requested MID Revision unsupported-response when trying to request unsupported MID Revision
	RequestedMIDUnsupportedAnswerIfTryingToRequestOnANonExistingMID                      ErrorCode = 75 // Requested MID unsupported-answer if trying to request on a non-existing MID
	RequestedOnSpecificDataNotSupportedResponseWhenTryingToRequestDataThatIsNotSupported ErrorCode = 77 // Requested on specific data not supported-response when trying to request data that is not supported
	SubscriptionOnSpecificDataNotSupportedAnswerIfTryingToSubscribeForUnsupportedData    ErrorCode = 78 // Subscription on specific data not supported-answer if trying to subscribe for unsupported data
	CommandFailed                                                                        ErrorCode = 79 // Command failed
	AudiEmergencyStatusSubscriptionExists                                                ErrorCode = 80 // Audi emergency status subscription exists
	AudiEmergencyStatusSubscriptionDoesNotExist                                          ErrorCode = 81 // Audi emergency status subscription does not exist
	AutomaticManualModeSubscribeAlreadyExist                                             ErrorCode = 82 // Automatic/Manual mode subscribe already exist
	AutomaticManualModeSubscribeDoesNotExist                                             ErrorCode = 83 // Automatic/Manual mode subscribe does not exist
	TheRelayFunctionSubscriptionAlreadyExists                                            ErrorCode = 84 // The relay function subscription already exists
	TheRelayFunctionSubscriptionDoesNotExist                                             ErrorCode = 85 // The relay function subscription does not exist
	TheSelectorSocketInfoSubscriptionAlreadyExist                                        ErrorCode = 86 // The selector socket info subscription already exist
	TheSelectorSocketInfoSubscriptionDoesNotExist                                        ErrorCode = 87 // The selector socket info subscription does not exist
	TheDiginInfoSubscriptionAlreadyExist                                                 ErrorCode = 88 // The digin info subscription already exist
	TheDiginInfoSubscriptionDoesNotExist                                                 ErrorCode = 89 // The digin info subscription does not exist
	LockAtBatchDoneSubscriptionAlreadyExist                                              ErrorCode = 90 // Lock at batch done subscription already exist
	LockAtBatchDoneSubscriptionDoesNotExist                                              ErrorCode = 91 // Lock at batch done subscription does not exist
	OpenProtocolCommandsDisabled                                                         ErrorCode = 92 // Open protocol commands disabled
	OpenProtocolCommandsDisabledSubscriptionAlreadyExists                                ErrorCode = 93 // Open protocol commands disabled subscription already exists
	OpenProtocolCommandsDisabledSubscriptionDoesNotExist                                 ErrorCode = 94 // Open protocol commands disabled subscription does not exist
	RejectRequestPowerMACSIsInManualMode                                                 ErrorCode = 95 // Reject request, Power MACS is in manual mode
	RejectConnectionClientAlreadyConnected                                               ErrorCode = 96 // Reject connection, Client already connected
	MIDRevisionUnsupported                                                               ErrorCode = 97 // MID revision unsupported
	ControllerInternalRequestTimeout                                                     ErrorCode = 98 // Controller internal request timeout
	UnknownMID                                                                           ErrorCode = 99 // Unknown MID
)

// MID 0004 Application Communication negative acknowledge
// This message is used by the controller when a request, command or subscription for any reason has
// not been performed. The data field contains the message ID of the message request that failed as well
// as an error code.
type MID0004REV001 struct {
	// MID number
	MIDNumber int `mid:"21-24"`
	// Error code for the sent message
	ErrorCode ErrorCode `mid:"25-26"`
}

func (m *MID0004REV001) Error() string {
	switch m.ErrorCode {
	case 0:
		return fmt.Sprintf("get error response on mid %d: No Error", m.MIDNumber)
	case 1:
		return fmt.Sprintf("get error response on mid %d: Invalid data", m.MIDNumber)
	case 2:
		return fmt.Sprintf("get error response on mid %d: Parameter set ID not present", m.MIDNumber)
	case 3:
		return fmt.Sprintf("get error response on mid %d: Parameter set can not be set.", m.MIDNumber)
	case 4:
		return fmt.Sprintf("get error response on mid %d: Parameter set not running", m.MIDNumber)
	case 6:
		return fmt.Sprintf("get error response on mid %d: VIN upload subscription already exists", m.MIDNumber)
	case 7:
		return fmt.Sprintf("get error response on mid %d: VIN upload subscription does not exists", m.MIDNumber)
	case 8:
		return fmt.Sprintf("get error response on mid %d: VIN input source not granted", m.MIDNumber)
	case 9:
		return fmt.Sprintf("get error response on mid %d: Last tightening result subscription already exists", m.MIDNumber)
	case 10:
		return fmt.Sprintf("get error response on mid %d: Last tightening result subscription does not exist", m.MIDNumber)
	case 11:
		return fmt.Sprintf("get error response on mid %d: Alarm subscription already exists", m.MIDNumber)
	case 12:
		return fmt.Sprintf("get error response on mid %d: Alarm subscription does not exist", m.MIDNumber)
	case 13:
		return fmt.Sprintf("get error response on mid %d: Parameter set selection subscription already exists", m.MIDNumber)
	case 14:
		return fmt.Sprintf("get error response on mid %d: Parameter set selection subscription does not exist", m.MIDNumber)
	case 15:
		return fmt.Sprintf("get error response on mid %d: Tightening ID requested not found", m.MIDNumber)
	case 16:
		return fmt.Sprintf("get error response on mid %d: Connection rejected protocol busy", m.MIDNumber)
	case 17:
		return fmt.Sprintf("get error response on mid %d: Job ID not present", m.MIDNumber)
	case 18:
		return fmt.Sprintf("get error response on mid %d: Job info subscription already exists", m.MIDNumber)
	case 19:
		return fmt.Sprintf("get error response on mid %d: Job info subscription does not exist", m.MIDNumber)
	case 20:
		return fmt.Sprintf("get error response on mid %d: Job can not be set", m.MIDNumber)
	case 21:
		return fmt.Sprintf("get error response on mid %d: Job not running", m.MIDNumber)
	case 22:
		return fmt.Sprintf("get error response on mid %d: Not possible to execute dynamic Job request", m.MIDNumber)
	case 23:
		return fmt.Sprintf("get error response on mid %d: Job batch decrement failed", m.MIDNumber)
	case 24:
		return fmt.Sprintf("get error response on mid %d: Not possible to create Pset", m.MIDNumber)
	case 25:
		return fmt.Sprintf("get error response on mid %d: Programming control not granted", m.MIDNumber)
	case 26:
		return fmt.Sprintf("get error response on mid %d: Wrong tool type to Pset download connected", m.MIDNumber)
	case 27:
		return fmt.Sprintf("get error response on mid %d: Tool is inaccessible", m.MIDNumber)
	case 28:
		return fmt.Sprintf("get error response on mid %d: Job abortion is in progress", m.MIDNumber)
	case 29:
		return fmt.Sprintf("get error response on mid %d: Tool does not exist", m.MIDNumber)
	case 30:
		return fmt.Sprintf("get error response on mid %d: Controller is not a sync Master/station controller", m.MIDNumber)
	case 31:
		return fmt.Sprintf("get error response on mid %d: Multi-spindle status subscription already exists", m.MIDNumber)
	case 32:
		return fmt.Sprintf("get error response on mid %d: Multi-spindle status subscription does not exist", m.MIDNumber)
	case 33:
		return fmt.Sprintf("get error response on mid %d: Multi-spindle result subscription already exists", m.MIDNumber)
	case 34:
		return fmt.Sprintf("get error response on mid %d: Multi-spindle result subscription does not exist", m.MIDNumber)
	case 35:
		return fmt.Sprintf("get error response on mid %d: Other master client already connected", m.MIDNumber)
	case 36:
		return fmt.Sprintf("get error response on mid %d: Lock type not supported", m.MIDNumber)
	case 40:
		return fmt.Sprintf("get error response on mid %d: Job line control info subscription already exists", m.MIDNumber)
	case 41:
		return fmt.Sprintf("get error response on mid %d: Job line control info subscription does not exist", m.MIDNumber)
	case 42:
		return fmt.Sprintf("get error response on mid %d: Identifier input source not granted", m.MIDNumber)
	case 43:
		return fmt.Sprintf("get error response on mid %d: Multiple identifiers work order subscription already exists", m.MIDNumber)
	case 44:
		return fmt.Sprintf("get error response on mid %d: Multiple identifiers work order subscription does not exist", m.MIDNumber)
	case 50:
		return fmt.Sprintf("get error response on mid %d: Status external monitored inputs subscription already exists", m.MIDNumber)
	case 51:
		return fmt.Sprintf("get error response on mid %d: Status external monitored inputs subscription does not exist", m.MIDNumber)
	case 52:
		return fmt.Sprintf("get error response on mid %d: IO device not connected", m.MIDNumber)
	case 53:
		return fmt.Sprintf("get error response on mid %d: Faulty IO device ID", m.MIDNumber)
	case 54:
		return fmt.Sprintf("get error response on mid %d: Tool Tag ID unknown", m.MIDNumber)
	case 55:
		return fmt.Sprintf("get error response on mid %d: Tool Tag ID subscription already exists", m.MIDNumber)
	case 56:
		return fmt.Sprintf("get error response on mid %d: Tool Tag ID subscription does not exist", m.MIDNumber)
	case 57:
		return fmt.Sprintf("get error response on mid %d: Tool Motor tuning failed", m.MIDNumber)
	case 58:
		return fmt.Sprintf("get error response on mid %d: No alarm present", m.MIDNumber)
	case 59:
		return fmt.Sprintf("get error response on mid %d: Tool currently in use", m.MIDNumber)
	case 60:
		return fmt.Sprintf("get error response on mid %d: No histogram available", m.MIDNumber)
	case 61:
		return fmt.Sprintf("get error response on mid %d: Pairing failed", m.MIDNumber)
	case 62:
		return fmt.Sprintf("get error response on mid %d: Pairing denied", m.MIDNumber)
	case 63:
		return fmt.Sprintf("get error response on mid %d: Pairing or Pairing abortion attempt on wrong tooltype", m.MIDNumber)
	case 64:
		return fmt.Sprintf("get error response on mid %d: Pairing abortion denied", m.MIDNumber)
	case 66:
		return fmt.Sprintf("get error response on mid %d: Pairing disconnection failed", m.MIDNumber)
	case 65:
		return fmt.Sprintf("get error response on mid %d: Pairing abortion failed", m.MIDNumber)
	case 67:
		return fmt.Sprintf("get error response on mid %d: Pairing in progress or already done", m.MIDNumber)
	case 68:
		return fmt.Sprintf("get error response on mid %d: Pairing denied. No Program Control", m.MIDNumber)
	case 69:
		return fmt.Sprintf("get error response on mid %d: Unsupported extra data revision", m.MIDNumber)
	case 70:
		return fmt.Sprintf("get error response on mid %d: Calibration failed", m.MIDNumber)
	case 71:
		return fmt.Sprintf("get error response on mid %d: Subscription already exists", m.MIDNumber)
	case 72:
		return fmt.Sprintf("get error response on mid %d: Subscription does not exists", m.MIDNumber)
	case 73:
		return fmt.Sprintf("get error response on mid %d: Subscribed MID unsupported, -answer if trying to subscribe on a non-existing MID", m.MIDNumber)
	case 74:
		return fmt.Sprintf("get error response on mid %d: Subscribed MID Revision unsupported,-answer if trying to subscribe on unsupported MID Revision.", m.MIDNumber)
	case 76:
		return fmt.Sprintf("get error response on mid %d: Requested MID Revision unsupported-response when trying to request unsupported MID Revision", m.MIDNumber)
	case 75:
		return fmt.Sprintf("get error response on mid %d: Requested MID unsupported-answer if trying to request on a non-existing MID", m.MIDNumber)
	case 77:
		return fmt.Sprintf("get error response on mid %d: Requested on specific data not supported-response when trying to request data that is not supported", m.MIDNumber)
	case 78:
		return fmt.Sprintf("get error response on mid %d: Subscription on specific data not supported-answer if trying to subscribe for unsupported data", m.MIDNumber)
	case 79:
		return fmt.Sprintf("get error response on mid %d: Command failed", m.MIDNumber)
	case 80:
		return fmt.Sprintf("get error response on mid %d: Audi emergency status subscription exists", m.MIDNumber)
	case 81:
		return fmt.Sprintf("get error response on mid %d: Audi emergency status subscription does not exist", m.MIDNumber)
	case 82:
		return fmt.Sprintf("get error response on mid %d: Automatic/Manual mode subscribe already exist", m.MIDNumber)
	case 83:
		return fmt.Sprintf("get error response on mid %d: Automatic/Manual mode subscribe does not exist", m.MIDNumber)
	case 84:
		return fmt.Sprintf("get error response on mid %d: The relay function subscription already exists", m.MIDNumber)
	case 85:
		return fmt.Sprintf("get error response on mid %d: The relay function subscription does not exist", m.MIDNumber)
	case 86:
		return fmt.Sprintf("get error response on mid %d: The selector socket info subscription already exist", m.MIDNumber)
	case 87:
		return fmt.Sprintf("get error response on mid %d: The selector socket info subscription does not exist", m.MIDNumber)
	case 88:
		return fmt.Sprintf("get error response on mid %d: The digin info subscription already exist", m.MIDNumber)
	case 89:
		return fmt.Sprintf("get error response on mid %d: The digin info subscription does not exist", m.MIDNumber)
	case 90:
		return fmt.Sprintf("get error response on mid %d: Lock at batch done subscription already exist", m.MIDNumber)
	case 91:
		return fmt.Sprintf("get error response on mid %d: Lock at batch done subscription does not exist", m.MIDNumber)
	case 92:
		return fmt.Sprintf("get error response on mid %d: Open protocol commands disabled", m.MIDNumber)
	case 93:
		return fmt.Sprintf("get error response on mid %d: Open protocol commands disabled subscription already exists", m.MIDNumber)
	case 94:
		return fmt.Sprintf("get error response on mid %d: Open protocol commands disabled subscription does not exist", m.MIDNumber)
	case 95:
		return fmt.Sprintf("get error response on mid %d: Reject request, Power MACS is in manual mode", m.MIDNumber)
	case 96:
		return fmt.Sprintf("get error response on mid %d: Reject connection, Client already connected", m.MIDNumber)
	case 97:
		return fmt.Sprintf("get error response on mid %d: MID revision unsupported", m.MIDNumber)
	case 98:
		return fmt.Sprintf("get error response on mid %d: Controller internal request timeout", m.MIDNumber)
	case 99:
		return fmt.Sprintf("get error response on mid %d: Unknown MID", m.MIDNumber)
	}
	return ""
}
