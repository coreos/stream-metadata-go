# Release notes

## Upcoming stream-metadata-go 0.4.4 (unreleased)

Changes:

- Add support for AppleHV images

## stream-metadata-go 0.4.3 (2023-06-28)

Changes:

- Add support for Hyper-V images


## stream-metadata-go 0.4.2 (2023-06-01)

Changes:

- Drop support for Go 1.17
- Add release notes doc


## stream-metadata-go 0.4.1 (2023-01-30)

Changes:

- Add support for Secure Execution


## stream-metadata-go 0.4.0 (2022-08-11)

- Drop vendored dependencies
- Drop support for Go 1.15 and 1.16
- Fix tests


## stream-metadata-go 0.3.0 (2022-04-15)

- Switch KubeVirt image struct from `SingleImage` to new `ContainerImage`
- Add `ContainerImage.DigestRef` field so we can record both the recommended
  pullspec and one that includes the digest


## stream-metadata-go 0.2.1 (2022-04-08)

- Add support for VirtualBox images


## stream-metadata-go 0.2.0 (2022-04-05)

- Rename `RegionObject`/`RegionImage` to `SingleObject`/`SingleImage`;
  add compatibility aliases
- Drop `KubeVirtContainerDisk` types in favor of generic ones


## stream-metadata-go 0.1.8 (2022-03-15)

- Add support for KubeVirt images


## stream-metadata-go 0.1.7 (2022-01-25)

- Fix Windows build by dropping renameio dependency


## stream-metadata-go 0.1.6 (2021-11-29)

- Add support for Nutanix images


## stream-metadata-go 0.1.5 (2021-11-02)

- Add support for IBMCloud/PowerVS images
- Add `GcpImage.Release` field in stream metadata
- Drop `omitempty` from mandatory GcpImage fields


## stream-metadata-go 0.1.4 (2021-10-06)

- Make `Artifact.Signature` `omitempty`


## stream-metadata-go 0.1.3 (2021-09-16)

- Add `Metadata.Generator` field


## stream-metadata-go 0.1.2 (2021-09-15)

- Add support for Alibaba Cloud (`aliyun`) cloud images


## stream-metadata-go 0.1.1 (2021-07-20)

- Add support for Azure Stack images


## stream-metadata-go 0.1.0 (2021-04-26)

- Initial release
