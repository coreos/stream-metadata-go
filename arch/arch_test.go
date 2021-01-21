// package arch contains mappings between the Golang architecture and
// the RPM architecture used by Fedora CoreOS and derivatives.
package arch

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapping(t *testing.T) {
	// Validate bidirectional mapping for current architecture
	assert.Equal(t, GoArch(CurrentRpmArch()), runtime.GOARCH)

	assert.Equal(t, GoArch("x86_64"), "amd64")
	assert.Equal(t, GoArch("aarch64"), "arm64")
	assert.Equal(t, GoArch("ppc64le"), "ppc64le")
}
