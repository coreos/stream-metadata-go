package stream

// Stream contains artifacts available in a stream
type Stream struct {
	Stream        string           `json:"stream"`
	Meta          BuildMetadata    `json:"metadata"`
	Architectures map[string]*Arch `json:"architectures"`
}

// Arch release details for x86_64 architetcure
type Arch struct {
	Artifacts Artifacts `json:"artifacts"`
	Images    *Images   `json:"images,omitempty"`
}

// Artifacts contains shipped artifacts list
type Artifacts struct {
	Aliyun       *MediaDetails `json:"aliyun,omitempty"`
	Aws          *MediaDetails `json:"aws,omitempty"`
	Azure        *MediaDetails `json:"azure,omitempty"`
	Digitalocean *MediaDetails `json:"digitalocean,omitempty"`
	Exoscale     *MediaDetails `json:"exoscale,omitempty"`
	Gcp          *MediaDetails `json:"gcp,omitempty"`
	Ibmcloud     *MediaDetails `json:"ibmcloud,omitempty"`
	Metal        *MediaDetails `json:"metal,omitempty"`
	Openstack    *MediaDetails `json:"openstack,omitempty"`
	Packet       *MediaDetails `json:"packet,omitempty"`
	Qemu         *MediaDetails `json:"qemu,omitempty"`
	Virtualbox   *MediaDetails `json:"virtualbox,omitempty"`
	Vmware       *MediaDetails `json:"vmware,omitempty"`
	Vultr        *MediaDetails `json:"vultr,omitempty"`
}

// MediaDetails contains image artifact and release detail
type MediaDetails struct {
	Release string                  `json:"release"`
	Formats map[string]*ImageFormat `json:"formats"`
}

// Images contains images available in cloud providers
type Images struct {
	Aws          *AwsImage   `json:"aws,omitempty"`
	Azure        *CloudImage `json:"azure,omitempty"`
	Gcp          *GcpImage   `json:"gcp,omitempty"`
	Digitalocean *CloudImage `json:"digitalocean,omitempty"`
	Packet       *CloudImage `json:"packet,omitempty"`
}

// CloudImage image for Cloud provider
type CloudImage struct {
	Image string `json:"image,omitempty"`
}

// AwsImage Aws images
type AwsImage struct {
	Regions map[string]*AwsAMI `json:"regions,omitempty"`
}

// AwsAMI aws AMI detail
type AwsAMI struct {
	Release string `json:"release"`
	Image   string `json:"image"`
}

// GcpImage GCP cloud image information
type GcpImage struct {
	Project string `json:"project,omitempty"`
	Family  string `json:"family,omitempty"`
	Name    string `json:"name,omitempty"`
}
