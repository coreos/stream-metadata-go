package stream

import (
	"crypto/sha256"
	"fmt"
	"io"
	"net/http"
)

// Fetch an artifact, validating its checksum.  If applicable,
// the artifact will not be decompressed.  Does not
// validate GPG signature.
func (a *Artifact) Fetch(w io.Writer) error {
	resp, err := http.Get(a.Location)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%s returned status: %s", a.Location, resp.Status)
	}
	hasher := sha256.New()
	reader := io.TeeReader(resp.Body, hasher)

	_, err = io.Copy(w, reader)
	if err != nil {
		return err
	}

	err = resp.Body.Close()
	if err != nil {
		return err
	}

	// Validate sha256 checksum
	foundChecksum := fmt.Sprintf("%x", hasher.Sum(nil))
	if a.Sha256 != foundChecksum {
		return fmt.Errorf("checksum mismatch for %s; expected=%s found=%s", a.Location, a.Sha256, foundChecksum)
	}

	return nil
}
