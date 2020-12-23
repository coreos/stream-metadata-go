package stream

// BuildMetadata for stream
type BuildMetadata struct {
	LastModified string `json:"last-modified"`
}

// ImageFormat contains Disk image details
type ImageFormat struct {
	Disk      *ImageType `json:"disk,omitempty"`
	Kernel    *ImageType `json:"kernel,omitempty"`
	Initramfs *ImageType `json:"initramfs,omitempty"`
	Rootfs    *ImageType `json:"rootfs,omitempty"`
}

// ImageType contains image detail
type ImageType struct {
	Location  string `json:"location"`
	Signature string `json:"signature"`
	Sha256    string `json:"sha256"`
}
