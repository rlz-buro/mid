package mid_test

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/suite"

	"github.com/rlz-buro/mid"
)

type MIDTestSuite struct {
	suite.Suite
}

func TestMIDTestSuite(t *testing.T) {
	suite.Run(t, new(MIDTestSuite))
}

func (suite *MIDTestSuite) TestUnmarshal() {
	data := []byte("00200001001000000000")
	m := mid.MID{}
	err := mid.UnmarshalMID(data, &m)
	suite.NoError(err)
	suite.Equal(len(data), m.Header.Length)
	suite.Equal(1, m.Header.MID)
	suite.Equal(1, m.Header.Revision)
}

func (suite *MIDTestSuite) TestMarshal() {
	data := []byte("00200001001000000000")
	m := mid.MID{
		Header: mid.Header{
			Length:   20,
			MID:      1,
			Revision: 1,
		},
	}
	raw, err := mid.MarshalMID(m)
	suite.NoError(err)
	suite.Len(raw, len(data))
	suite.Equal(data, raw)
}

func (suite *MIDTestSuite) TestUnmarshal101() {
	data := []byte("02100101   0        01020286132552$L604117         030004004050000060000070080026000900340010      11     12     1300003142020-02-02:16:43:38152020-06-09:09:56:30160805217118050111003088100034060211003007100045")

	mid101 := mid.MID0101REV001{}
	err := mid.Unmarshal(data, &mid101)
	suite.NoError(err)
	spew.Dump(mid101)
}
