package release

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

func TestParseFCRIndex(t *testing.T) {
	d, err := ioutil.ReadFile("fixtures/fcos-releases.json")
	assert.Nil(t, err)
	idx := Index{}
	err = json.Unmarshal(d, &idx)
	assert.Nil(t, err)
	assert.Equal(t, idx.Stream, "stable")
	release := idx.Releases[0]
	assert.Equal(t, release.Version, "31.20200108.3.0")
	assert.Equal(t, release.Commits[0].Architecture, "x86_64")
	assert.Equal(t, release.Commits[0].Checksum, "113aa27efe1bbcf6324af7423f64ef7deb0acbf21b928faec84bf66a60a5c933")
}
