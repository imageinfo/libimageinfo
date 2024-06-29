package libimageinfo

import (
	"github.com/imageinfo/libimageinfo/filehandlers"
	"github.com/imageinfo/libimageinfo/imagemetadata"
	"os"
)

type LibImageInfo struct {
	fileHandlerMgr *filehandlers.FileHandlerMgr
}

func NewLibImageInfo() *LibImageInfo {
	return &LibImageInfo{
		fileHandlerMgr: filehandlers.FileHandlerMgrInstance(),
	}
}

func (*LibImageInfo) ReadImageMetadata(filePath string) *imagemetadata.ImageMetadata {
	// Try to read an image
	imageFile, err := os.OpenFile(filePath, os.O_RDONLY, 0400)
	if err != nil {
		panic("Could not open image file")
	}
	defer imageFile.Close()

	filehandlers.FileHandlerMgrInstance()

	return nil
}
