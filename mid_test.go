package mid_test

import (
	"fmt"
	"testing"

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
	err := mid.Unmarshal(data, &m)
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
	raw, err := mid.Marshal(m)
	suite.NoError(err)
	suite.Len(raw, len(data))
	fmt.Println(string(raw))
	suite.Equal(data, raw)
}
