package stream

// Metadata for a release or stream
type Metadata struct {
	LastModified string `json:"last-modified"`
}

// ImageFormat contains all artifacts for a single OS image
type ImageFormat struct {
	Disk      *Artifact `json:"disk,omitempty"`
	Kernel    *Artifact `json:"kernel,omitempty"`
	Initramfs *Artifact `json:"initramfs,omitempty"`
	Rootfs    *Artifact `json:"rootfs,omitempty"`
}

// Artifact represents one image file, plus its metadata
type Artifact struct {
	Location  string `json:"location"`
	Signature string `json:"signature"`
	Sha256    string `json:"sha256"`
}
