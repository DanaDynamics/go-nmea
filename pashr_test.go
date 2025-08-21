package nmea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var pashrtests = []struct {
	name string
	raw  string
	err  string
	msg  PASHR
}{
	{
		name: "good sentence A",
		raw:  "$PASHR,085335.000,224.19,T,-01.26,+00.83,+00.10,0.101,0.113,0.267,1,0*07",
		msg: PASHR{
			Time: Time{
				Valid:       true,
				Hour:        8,
				Minute:      53,
				Second:      35,
				Millisecond: 0,
			},
			Heading:            224.19,
			TrueHeading:        "T",
			Roll:               -1.26,
			Pitch:              0.83,
			Heave:              0.1,
			RollAccuracy:       0.101,
			PitchAccuracy:      0.113,
			HeadingAccuracy:    0.267,
			GNSSQuality:        1,
			IMUAlignmentStatus: 0,
		},
	},
	{
		name: "invalid checksum",
		raw:  "$PASHR,085335.000,224.19,T,-01.26,+00.83,+00.10,0.101,0.0,0.267,1,0*07",
		err:  "nmea: sentence checksum mismatch [04 != 07]",
	},
	{
		name: "invalid true heading value",
		raw:  "$PASHR,085335.000,224.19,X,-01.26,+00.83,+00.10,0.101,0.113,0.267,1,0*0B",
		err:  "nmea: PASHR invalid true  heading: X",
	},
}

func TestPASHR(t *testing.T) {
	for _, tt := range pashrtests {
		t.Run(tt.name, func(t *testing.T) {

			m, err := Parse(tt.raw)
			if tt.err != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.err)
			} else {
				assert.NoError(t, err)
				pashr := m.(PASHR)
				pashr.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, pashr)
			}
		})
	}
}
