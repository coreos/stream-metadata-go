package stream

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseFCS(t *testing.T) {
	d, err := ioutil.ReadFile("fixtures/fcos-stable.json")
	assert.Nil(t, err)
	stream := Metadata{}
	err = json.Unmarshal(d, &stream)
	assert.Nil(t, err)
	assert.Equal(t, stream.Stream, "stable")
	assert.Equal(t, stream.Architectures["x86_64"].Images.Aws.Regions["us-east-2"].Image, "ami-091b0dbc05fe2dc06")
}
