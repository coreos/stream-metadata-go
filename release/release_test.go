package release

import (
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/coreos/stream-metadata-go/stream"
	"github.com/stretchr/testify/assert"
)

const (
	usEast2Ami = "ami-091b0dbc05fe2dc06"
)

func TestParseFCR(t *testing.T) {
	d, err := os.ReadFile("fixtures/fcos-release.json")
	assert.Nil(t, err)
	release := Release{}
	err = json.Unmarshal(d, &release)
	assert.Nil(t, err)
	assert.Equal(t, release.Stream, "stable")
	assert.Equal(t, release.Architectures["x86_64"].Media.Aws.Images["us-east-2"].Image, usEast2Ami)
}

func TestParseFCRIndex(t *testing.T) {
	d, err := os.ReadFile("fixtures/fcos-releases.json")
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

func TestTranslate(t *testing.T) {
	d, err := os.ReadFile("fixtures/fcos-release.json")
	assert.Nil(t, err)
	rel := Release{}
	err = json.Unmarshal(d, &rel)
	assert.Nil(t, err)
	arches := rel.ToStreamArchitectures()
	st := stream.Stream{
		Stream:        rel.Stream,
		Metadata:      stream.Metadata{LastModified: time.Now().UTC().Format(time.RFC3339)},
		Architectures: arches,
	}
	assert.Equal(t, st.Architectures["x86_64"].Images.Aws.Regions["us-east-2"].Image, usEast2Ami)

	// KubeVirt
	assert.Equal(t, st.Architectures["x86_64"].Images.KubeVirt.Image, "quay.io/openshift-release-dev/rhcos@sha256:67a81539946ec0397196c145394553b8e0241acf27b14ae9de43bc56e167f773")
}
