package mid_test

import (
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

func (suite *MIDTestSuite) TestUnmarshal61() {
	data := []byte("02310061001 0000    010000020003Station P1               0486135507$6101010         050106004070004080002091101111120026001300370014003200150032081600000170000018000001900000202023-08-21:11:49:02212020-11-24:20:27:58220230003951962")
	mid61 := mid.MID0061REV001{}
	err := mid.Unmarshal(data, &mid61)
	suite.NoError(err)
}

func (suite *MIDTestSuite) TestUnmarshal106() {
	data := []byte("14090106            01190201031073817763040105CylinderHead-S      062023-06-16:19:28:15070308DEP ISB 4.5         0911001186132968                                1218130114115116117181.44518180.26419250.00020110.00021185.00022175.000130214115116117173.49718180.26419250.00020110.00021185.00022175.000130314115116117161.96818180.26419250.00020110.00021185.00022175.000130414115116117161.25318179.86719250.00020110.00021185.00022175.000130514115116117162.34618180.26419250.00020110.00021185.00022175.000130614115116117175.33318180.26419250.00020110.00021185.00022175.000130714115116117149.15518180.26419250.00020110.00021185.00022175.000130814115116117163.51818180.26419250.00020110.00021185.00022175.000130914115116117173.16418180.26419250.00020110.00021185.00022175.000131014115116117168.48618180.26419250.00020110.00021185.00022175.000131114115116117163.64918180.26419250.00020110.00021185.00022175.000131214115116117169.83718180.26419250.00020110.00021185.00022175.000131514115116117161.98318180.26419250.00020110.00021185.00022175.000131614115116117196.68418180.26419250.00020110.00021185.00022175.000131914115116117164.45318180.26419250.00020110.00021185.00022175.000132014115116117165.12818180.26419250.00020110.00021185.00022175.000132314115116117172.18618180.26419250.00020110.00021185.00022175.000132414115116117164.35618180.26419250.00020110.00021185.00022175.0002301Data No Station     I 100000010006")
	mid106 := mid.MID0106REV001{}
	err := mid.Unmarshal(data, &mid106)
	suite.NoError(err)
}

func (suite *MIDTestSuite) TestUnmarshal107() {
	data := []byte("02660107            010302020300005192190401052020-02-17:12:03:4806000107Bolt 01             0830Nm                092100000000000000000000000000327680000000064000000000011    1201Spindle No          I 0000001130001411501Program Time        T 192020-02-02:16:43:3600")
	mid107 := mid.MID0107REV001{}
	err := mid.Unmarshal(data, &mid107)
	suite.NoError(err)
}
