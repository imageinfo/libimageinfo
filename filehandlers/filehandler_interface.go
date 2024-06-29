package filehandlers

import (
	"github.com/imageinfo/libimageinfo/imagemetadata"
	"io"
)

type ImageFileMetadataReader interface {
	Name() string
	Description() string
	ExtractImageMetadata(*io.ReadSeeker) *imagemetadata.ImageMetadata
}

type ImageFileMetadataWriter interface {
	Name() string
	Description() string
	CreateOrReplaceImageMetadata(imagemetadata.ImageMetadata, *io.ReadWriteSeeker) error
}
