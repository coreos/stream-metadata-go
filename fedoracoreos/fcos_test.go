package fedoracoreos

import "testing"

func TestGetStreamURL(t *testing.T) {
	u := GetStreamURL(StreamStable)
	if u.String() != "https://builds.coreos.fedoraproject.org/streams/stable.json" {
		t.Fatalf("Invalid stream url")
	}
}
