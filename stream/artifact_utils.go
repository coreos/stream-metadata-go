package stream

import (
	"crypto/sha256"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"

	"github.com/google/renameio"
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

	// Validate sha256 checksum
	foundChecksum := fmt.Sprintf("%x", hasher.Sum(nil))
	if a.Sha256 != foundChecksum {
		return fmt.Errorf("checksum mismatch for %s; expected=%s found=%s", a.Location, a.Sha256, foundChecksum)
	}

	return nil
}

/// Download fetches the specified artifact and saves it to the target
/// directory.  The full file path will be returned as a string.
/// If the target file path exists, it will be overwritten.
/// If the download fails, the temporary file will be deleted.
func (a *Artifact) Download(destdir string) (string, error) {
	loc, err := url.Parse(a.Location)
	if err != nil {
		return "", err
	}
	name := path.Base(loc.Path)
	destfile := path.Join(destdir, name)
	w, err := renameio.TempFile("", destfile)
	if err != nil {
		return "", err
	}

	defer func() {
		// Ignore an error to unlink
		_ = w.Cleanup()
	}()
	err = a.Fetch(w)
	if err != nil {
		return "", err
	}
	if err := w.File.Chmod(0644); err != nil {
		return "", err
	}
	err = w.CloseAtomicallyReplace()
	if err != nil {
		return "", err
	}

	return destfile, nil
}
