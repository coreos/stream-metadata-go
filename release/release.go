package release

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

// GcpImage represents a GCP cloud image
type GcpImage struct {
	Project string `json:"project,omitempty"`
	Family  string `json:"family,omitempty"`
	Name    string `json:"name,omitempty"`
}

// Release contains details from release.json
type Release struct {
	Release       string                 `json:"release"`
	Stream        string                 `json:"stream"`
	Metadata      Metadata               `json:"metadata"`
	Architectures map[string]ReleaseArch `json:"architectures"`
}

// ReleaseArch release details
type ReleaseArch struct {
	Media ReleaseMedia `json:"media"`
}

// ReleaseMedia contains release details for various platforms
type ReleaseMedia struct {
	Aliyun       *ReleasePlatformBase `json:"aliyun"`
	Aws          *ReleasePlatformAws  `json:"aws"`
	Azure        *ReleasePlatformBase `json:"azure"`
	Digitalocean *ReleasePlatformBase `json:"digitalocean"`
	Exoscale     *ReleasePlatformBase `json:"exoscale"`
	Gcp          *ReleasePlatformGcp  `json:"gcp"`
	Ibmcloud     *ReleasePlatformBase `json:"ibmcloud"`
	Metal        *ReleasePlatformBase `json:"metal"`
	Openstack    *ReleasePlatformBase `json:"openstack"`
	Qemu         *ReleasePlatformBase `json:"qemu"`
	Vmware       *ReleasePlatformBase `json:"vmware"`
	Vultr        *ReleasePlatformBase `json:"vultr"`
}

// Release platform with no cloud images
type ReleasePlatformBase struct {
	Artifacts map[string]ImageFormat `json:"artifacts"`
}

// ReleasePlatformAws contains AWS image information
type ReleasePlatformAws struct {
	ReleasePlatformBase
	Images map[string]ReleaseCloudImage `json:"images"`
}

// ReleasePlatformGcp GCP image detail
type ReleasePlatformGcp struct {
	ReleasePlatformBase
	Image *GcpImage `json:"image"`
}

// ReleaseCloudImage generic image detail
type ReleaseCloudImage struct {
	Image string `json:"image"`
}
