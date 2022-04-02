// Package main contains an example use of this library; it
// prints the current Fedora CoreOS EC2(AWS) x86_64 AMI in the
// us-east-2 region.
package main

import (
	"fmt"
	"os"

	"github.com/coreos/stream-metadata-go/fedoracoreos"
	"github.com/coreos/stream-metadata-go/stream"
)

const (
	targetArch = "x86_64"
	region     = "us-east-2"
)

func downloadISO(fcosstable stream.Stream) error {
	iso, err := fcosstable.QueryDisk(targetArch, "metal", "iso")
	if err != nil {
		return err
	}

	fmt.Printf("Downloading %s\n", iso.Location)
	path, err := iso.Download(".")
	if err != nil {
		return fmt.Errorf("Failed to download %s: %w", iso.Location, err)
	}
	fmt.Printf("Downloaded %s\n", path)

	return nil
}

func printAMI(fcosstable stream.Stream) error {
	arch, ok := fcosstable.Architectures[targetArch]
	if !ok {
		return fmt.Errorf("No %s architecture in stream", targetArch)
	}
	awsimages := arch.Images.Aws
	if awsimages == nil {
		return fmt.Errorf("No %s AWS images in stream", targetArch)
	}
	var regionVal stream.SingleImage
	if regionVal, ok = awsimages.Regions[region]; !ok {
		return fmt.Errorf("No %s AWS images in region %s", targetArch, region)
	}
	fmt.Printf("%s\n", regionVal.Image)

	return nil
}

func run() error {
	if len(os.Args) != 2 {
		return fmt.Errorf("usage: example aws-ami|download-iso")
	}
	arg := os.Args[1]
	fcosstable, err := fedoracoreos.FetchStream(fedoracoreos.StreamStable)
	if err != nil {
		return err
	}
	if arg == "aws-ami" {
		return printAMI(*fcosstable)
	} else if arg == "download-iso" {
		return downloadISO(*fcosstable)
	} else {
		return fmt.Errorf("invalid operation %s", arg)
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
