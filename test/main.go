package main

import (
	"fmt"
	"github.com/akamensky/argparse"
	"github.com/imageinfo/libimageinfo"
	"os"
	"path/filepath"
	"strings"
)

type programOptions struct {
	inputImageFile *os.File
	outputXmlFile  string
}

func parseArgs() programOptions {
	parser := argparse.NewParser("image2lrxmp", "Create Lightroom-friendly XMP for image files")
	inputImageFile := parser.File("i", "inputimage", os.O_RDONLY, 0400, &argparse.Options{
		Required: true,
		Help:     "Input image file",
	})
	outputXmpFile := parser.String("o", "xmpfilename", &argparse.Options{
		Required: false,
		Help:     "XMP file name if user doesn't want default",
	})
	if err := parser.Parse(os.Args); err != nil {
		fmt.Print(parser.Usage(err))
	}
	var xmpFilename string
	if outputXmpFile == nil {
		xmpFilename = strings.TrimSuffix(inputImageFile.Name(), filepath.Ext(inputImageFile.Name())) + ".xmp"
	} else {
		xmpFilename = *outputXmpFile
	}

	return programOptions{
		inputImageFile: inputImageFile,
		outputXmlFile:  xmpFilename,
	}
}

func main() {
	//args := parseArgs()

	//inputFileStat, err := os.Stat(args.inputImageFile.Name())
	//if err != nil {
	//	panic("Could not run stat on our open input file?!?!?")
	//}

	libHandle := libimageinfo.NewLibImageInfo()

	// Get the metadata
	libHandle.ReadImageMetadata("C:\\Temp\\PhotoEndOfDay\\Input\\input1\\_50A0001.CR3")
}
