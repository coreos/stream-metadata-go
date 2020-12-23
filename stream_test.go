package stream

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseFCS(t *testing.T) {
	d, err := ioutil.ReadFile("fixtures/fcos-stream.json")
	assert.Nil(t, err)
	stream := Stream{}
	err = json.Unmarshal(d, &stream)
	assert.Nil(t, err)
	assert.Equal(t, stream.Stream, "stable")
	assert.Equal(t, stream.Architectures["x86_64"].Artifacts["metal"].Formats["raw.xz"].Disk.Sha256, "2848b111a6917455686f38a3ce64d2321c33809b9cf796c5f6804b1c02d79d9d")
	assert.Equal(t, stream.Architectures["x86_64"].Images.Aws.Regions["us-east-2"].Image, "ami-091b0dbc05fe2dc06")
}
