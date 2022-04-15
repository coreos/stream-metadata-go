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
	assert.Equal(t, stream.Architectures["x86_64"].Images.Aliyun.Regions["eu-west-1"].Image, "m-d7o3sn3qtbax50wep8ra")

	ami, err := stream.GetAMI("x86_64", "us-east-2")
	assert.Nil(t, err)
	assert.Equal(t, ami, "ami-091b0dbc05fe2dc06")

	aliyunImage, err := stream.GetAliyunImage("x86_64", "eu-west-1")
	assert.Nil(t, err)
	assert.Equal(t, aliyunImage, "m-d7o3sn3qtbax50wep8ra")

	assert.Equal(t, "stable/x86_64", stream.FormatPrefix("x86_64"))

	// I hope I live to see the day when we might change this code to test for success and not error
	_, err = stream.GetAMI("x86_64", "mars-1")
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "stable/x86_64: No AWS images in region mars-1")

	_, err = stream.GetAMI("aarch64", "us-east-2")
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "stream:stable does not have architecture 'aarch64'")

	_, err = stream.GetAliyunImage("x86_64", "us-midwest-1")
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "stable/x86_64: No Aliyun images in region us-midwest-1")

	_, err = stream.GetAliyunImage("aarch64", "us-east-2")
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "stream:stable does not have architecture 'aarch64'")

	a, err := stream.QueryDisk("x86_64", "metal", "iso")
	assert.Nil(t, err)
	assert.Equal(t, a.Location, "https://builds.coreos.fedoraproject.org/prod/streams/stable/builds/33.20201201.3.0/x86_64/fedora-coreos-33.20201201.3.0-live.x86_64.iso")

	name, err := a.Name()
	assert.Nil(t, err)
	assert.Equal(t, name, "fedora-coreos-33.20201201.3.0-live.x86_64.iso")

	_, err = stream.QueryDisk("x86_64", "metal", "nosuchthing")
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "disk not found")
	_, err = stream.QueryDisk("x86_64", "mysterious-obelisk", "rune")
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "artifact 'mysterious-obelisk' not found")
	_, err = stream.QueryDisk("nonarch", "", "nosuchthing")
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "does not have architecture 'nonarch'")

	assert.Equal(t, stream.Architectures["x86_64"].Images.KubeVirt, &ContainerImage{
		Release:   "33.20211201.3.0",
		Image:     "quay.io/openshift-release-dev/rhcos@latest",
		DigestRef: "quay.io/openshift-release-dev/rhcos@sha256:67a81539946ec0397196c145394553b8e0241acf27b14ae9de43bc56e167f773",
	})
	assert.Equal(t, stream.Architectures["x86_64"].Artifacts["kubevirt"].Formats["qcow2.xz"].Disk.Sha256, "2be55c5aa1f53eb9a869826dacbab75706ee6bd59185b935ac9be546cc132a85")
}
