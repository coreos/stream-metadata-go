package stream

// Stream contains artifacts available in a stream
type Stream struct {
	Stream        string          `json:"stream"`
	Metadata      Metadata        `json:"metadata"`
	Architectures map[string]Arch `json:"architectures"`
}

// Architecture release details
type Arch struct {
	Artifacts map[string]PlatformArtifacts `json:"artifacts"`
	Images    Images                       `json:"images,omitempty"`
}

// PlatformArtifacts contains images for a platform
type PlatformArtifacts struct {
	Release string                 `json:"release"`
	Formats map[string]ImageFormat `json:"formats"`
}

// Images contains images available in cloud providers
type Images struct {
	Aws *AwsImage `json:"aws,omitempty"`
	Gcp *GcpImage `json:"gcp,omitempty"`
}

// AwsImage represents an image across all AWS regions
type AwsImage struct {
	Regions map[string]AwsRegionImage `json:"regions,omitempty"`
}

// AwsRegionImage represents an image in one AWS region
type AwsRegionImage struct {
	Release string `json:"release"`
	Image   string `json:"image"`
}
