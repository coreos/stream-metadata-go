package stream

// Release contains details from release.json
type Release struct {
	Release       string                  `json:"release"`
	Stream        string                  `json:"stream"`
	Metadata      Metadata                `json:"metadata"`
	Architectures map[string]*ReleaseArch `json:"architectures"`
}

// ReleaseArch release details for x86_64 architetcure
type ReleaseArch struct {
	Media *ReleaseMedia `json:"media"`
}

// ReleaseMedia contains details about various images we ship
type ReleaseMedia struct {
	Aliyun       *ReleaseTargetPlatform `json:"aliyun"`
	Aws          *ReleaseAws            `json:"aws"`
	Azure        *ReleaseAzure          `json:"azure"`
	Digitalocean *ReleaseDigitalOcean   `json:"digitalocean"`
	Exoscale     *ReleaseTargetPlatform `json:"exoscale"`
	Gcp          *ReleaseGcp            `json:"gcp"`
	Ibmcloud     *ReleaseTargetPlatform `json:"ibmcloud"`
	Metal        *ReleaseTargetPlatform `json:"metal"`
	Openstack    *ReleaseTargetPlatform `json:"openstack"`
	Packet       *ReleaseTargetPlatform `json:"packet"`
	Qemu         *ReleaseTargetPlatform `json:"qemu"`
	Virtualbox   *ReleaseTargetPlatform `json:"virtualbox"`
	Vmware       *ReleaseTargetPlatform `json:"vmware"`
	Vultr        *ReleaseTargetPlatform `json:"vultr"`
}

// ReleaseAws contains AWS image information
type ReleaseAws struct {
	Artifacts map[string]*ImageFormat        `json:"artifacts"`
	Images    *map[string]*ReleaseCloudImage `json:"images"`
}

// ReleaseDigitalOcean DigitalOcean image detail
type ReleaseDigitalOcean struct {
	Artifacts map[string]*ImageFormat `json:"artifacts"`
	Image     string                  `json:"image"`
}

// ReleaseAzure Azure image detail
type ReleaseAzure struct {
	Artifacts map[string]*ImageFormat `json:"artifacts"`
	Images    *ReleaseAzureImages     `json:"images"`
}

// ReleaseAzureImages Azure image detail
type ReleaseAzureImages struct {
	Global *ReleaseCloudImage `json:"global"`
}

// ReleaseGcp GCP image detail
type ReleaseGcp struct {
	Artifacts map[string]*ImageFormat `json:"artifacts"`
	Image     *ReleaseCloudImageGcp   `json:"image"`
}

// ReleaseCloudImageGcp GCP cloud image information
type ReleaseCloudImageGcp struct {
	Project string `json:"project,omitempty"`
	Family  string `json:"family,omitempty"`
	Name    string `json:"name,omitempty"`
}

// ReleaseCloudImage cloud image information
type ReleaseCloudImage struct {
	Image *string `json:"image"`
}

// ReleaseTargetPlatform target platforms
type ReleaseTargetPlatform struct {
	Artifacts map[string]*ImageFormat `json:"artifacts"`
}
