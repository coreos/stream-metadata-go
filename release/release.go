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
	Release       string          `json:"release"`
	Stream        string          `json:"stream"`
	Metadata      Metadata        `json:"metadata"`
	Architectures map[string]Arch `json:"architectures"`
}

// Arch release details
type Arch struct {
	Media Media `json:"media"`
}

// Media contains release details for various platforms
type Media struct {
	Aliyun       *PlatformBase `json:"aliyun"`
	Aws          *PlatformAws  `json:"aws"`
	Azure        *PlatformBase `json:"azure"`
	Digitalocean *PlatformBase `json:"digitalocean"`
	Exoscale     *PlatformBase `json:"exoscale"`
	Gcp          *PlatformGcp  `json:"gcp"`
	Ibmcloud     *PlatformBase `json:"ibmcloud"`
	Metal        *PlatformBase `json:"metal"`
	Openstack    *PlatformBase `json:"openstack"`
	Qemu         *PlatformBase `json:"qemu"`
	Vmware       *PlatformBase `json:"vmware"`
	Vultr        *PlatformBase `json:"vultr"`
}

// PlatformBase with no cloud images
type PlatformBase struct {
	Artifacts map[string]ImageFormat `json:"artifacts"`
}

// PlatformAws contains AWS image information
type PlatformAws struct {
	PlatformBase
	Images map[string]CloudImage `json:"images"`
}

// PlatformGcp GCP image detail
type PlatformGcp struct {
	PlatformBase
	Image *GcpImage `json:"image"`
}

// CloudImage generic image detail
type CloudImage struct {
	Image string `json:"image"`
}
