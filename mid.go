package mid

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

const midTagName = "mid"

type MID struct {
	Header Header
	Data   []byte
}

type Header struct {
	// The length is the length of the header plus the data field excluding the NUL termination.
	// The header always includes information about the length of the message. The length is represented by four ASCII digits (‘0’...’9’) specifying a range of 0000 to 9999.
	// When using the message linking functionality the length represents the length of each message part number.
	// When having one ASCII part followed by an binary part the length is the total length of the message.
	Length int `mid:"1-4"`
	// The MID is four bytes long and is specified by four ASCII digits (‘0’...’9’). The MID describes how to interpret the message.
	MID int `mid:"5-8"`
	// 	The revision of the MID is specified by three ASCII digits (‘0’...’9’).
	// The MID Revision is unique per MID and is used in case different versions are available for the same MID. Using the revision number the integrator can subscribe or ask for different versions of the same MID. By default the MID revision number is three spaces long.
	// If the initial MID Revision (revision 1) is required there is three different ways to get it, either send three spaces or 000 or 001.
	Revision int `mid:"9-11"`
	// 	ONLY FOR SUBSCRIPTION MIDs.
	// The No Ack Flag is used when setting a subscription. If the No Ack flag is not set in a subscription it means that the subscriber will acknowledge each “push” message sent by the controller (reliable mode).
	// If set, the controller will only push out the information required without waiting for a receive acknowledgement from the subscriber (unreliable mode).
	// Note! NOT USED WHEN USING SEQUENCE NUMBER HANDLING
	NoAckFlag bool `mid:"12"`
	// The station the message is addressed to in the case of controller with multi-station configuration. The station ID is 2 byte long and is specified by two ASCII digits (‘0’...’9’). Two spaces are considered as station 1 (default value).
	StationID int `mid:"13-14"`
	// The spindle the message is addressed to in the case several spindles are connected to the same controller. The spindle ID is 2 bytes long and is specified by two ASCII digits (‘0’...’9’). Two spaces are considered as spindle 1 (default value).
	SpindleID int `mid:"15-16"`
	// 	From OP Spec. 2.0. 1-99-1. For acknowledging on “Link Level” with MIDs 0997 and 0998.
	// Not used if space or zero and not 1-99.
	// At communication restart MID 0001/MID 0002 it must be set to one and info in MID 0002 is telling if possible to use or not. It is backward compatible and If used it will substitute the No Ack flag and all special subscription data messages ACK MIDs.
	SequenceNumber int `mid:"17-18"`
	// 	From OP spec. 2.0. Linking function can be up to 9 = possible to send 9*9999 bytes messages. ~ 90 kB.
	// Used when the message length is overflowing the max length of 9999.Not used if space or zero.
	NumberOfMessageParts int `mid:"19"`
	// 	From OP spec. 2.0. Linking function, can be 1- 9 at message length > 9999.
	// Not used if space or zero
	MessagePartNumber int `mid:"20"`
}

func Marshal(v MID) ([]byte, error) {
	header, err := marshal(&v.Header)
	if err != nil {
		return nil, err
	}
	data := v.Data
	return append(header, data...), nil
}

func marshal(v any) ([]byte, error) {
	raw := []byte{}
	rv := reflect.ValueOf(v).Elem()
	rt := reflect.TypeOf(v).Elem()
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		tag := field.Tag.Get(midTagName)
		s, e, err := parseTag(tag, 1, 20)
		if err != nil {
			return nil, fmt.Errorf("invalid mid tag: %w", err)
		}
		var v string
		switch field.Type.Kind() {
		case reflect.Int:
			v = strconv.Itoa(int(rv.Field(i).Int()))
			if len(v) < e-s+1 {
				v = strings.Repeat("0", e-s+1-len(v)) + v
			}
		case reflect.Bool:
			if rv.Field(i).Bool() {
				v = "1"
			} else {
				v = "0"
			}
		case reflect.String:
			v = rv.Field(i).String()
		}
		raw = append(raw, []byte(v)...)
	}
	return raw, nil
}

func Unmarshal(data []byte, v *MID) error {
	if l := len(data); l < 20 {
		return fmt.Errorf("invalid header: header size should be 20 bytes but actual header has only %d", l)
	}
	if err := unmarshal(data, &v.Header, 1, 20); err != nil {
		return err
	}
	if len(data) > 20 {
		v.Data = data[20:]
	}
	return nil
}

func unmarshal(data []byte, v any, min, max int) error {
	rv := reflect.ValueOf(v).Elem()
	rt := reflect.TypeOf(v).Elem()
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		tag := field.Tag.Get(midTagName)
		s, e, err := parseTag(tag, 1, 20)
		if err != nil {
			return fmt.Errorf("invalid mid tag: %w", err)
		}
		token := data[s-1 : e]
		if string(token) == strings.Repeat(" ", len(token)) {
			continue
		}
		switch field.Type.Kind() {
		case reflect.Int:
			val, err := strconv.Atoi(string(token))
			if err != nil {
				return fmt.Errorf("invalid header: %w", err)
			}
			rv.Field(i).SetInt(int64(val))
		case reflect.Bool:
			val, err := strconv.Atoi(string(token))
			if err != nil {
				return fmt.Errorf("invalid header: %w", err)
			}
			rv.Field(i).SetBool(val != 0)
		case reflect.String:
			rv.Field(i).SetString(string(token))
		}
	}
	return nil
}

func parseTag(tag string, min, max int) (int, int, error) {
	var (
		start int
		end   int
		err   error
	)
	tokens := strings.Split(tag, "-")
	switch len(tokens) {
	case 1:
		start, err = strconv.Atoi(tokens[0])
		if err != nil {
			return 0, 0, err
		}
		end = start
	case 2:
		start, err = strconv.Atoi(tokens[0])
		if err != nil {
			return 0, 0, err
		}
		end, err = strconv.Atoi(tokens[1])
		if err != nil {
			return 0, 0, err
		}
	default:
		return 0, 0, fmt.Errorf("wrong mid tag format: %q", tag)
	}
	if start < min || end > max {
		return 0, 0, fmt.Errorf("indexes should be %d <= i <= %d: start - %d end - %d", min, max, start, end)
	}
	return start, end, nil
}
