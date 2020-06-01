package vowelnum

import "testing"

func testConvertText(t *testing.T, samples []testTuple, mapping []string) {
	for _, tt := range samples {
		t.Run(tt.in, func(t *testing.T) {
			s := convertText(mapping, tt.in)
			if s != tt.out {
				t.Errorf("got %q, want %q", s, tt.out)
			}
		})
	}
}

type testTuple struct {
	in, out string
}
var encodeTests = []testTuple {
	{"hello", "h2ll4"},
	{"Hello", "h2ll4"},
	{"Golang", "g4l1ng"},
}

func TestEncode(t *testing.T) {
	testConvertText(t, encodeTests, encoderMap)
}

var decodeTests = []testTuple{
	{"h2ll4", "hello"},
	{"v1l61lk3n", "valyalkin"},
	{"g4l1ng", "golang"},
}

func TestDecode(t *testing.T) {
	testConvertText(t, decodeTests, decoderMap)
}