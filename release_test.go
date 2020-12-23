package stream

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseFCR(t *testing.T) {
	d, err := ioutil.ReadFile("fixtures/fcos-release.json")
	assert.Nil(t, err)
	release := Release{}
	err = json.Unmarshal(d, &release)
	assert.Nil(t, err)
	assert.Equal(t, release.Stream, "stable")
	assert.Equal(t, release.Architectures["x86_64"].Media.Aws.Images["us-east-2"].Image, "ami-091b0dbc05fe2dc06")
}
