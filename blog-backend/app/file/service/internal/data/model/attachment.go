package model

import "gorm.io/gorm"

type Attachment struct {
	gorm.Model

	Name      *string
	Path      *string
	FileKey   *string
	ThumbPath *string
	MediaType *string
	Suffix    *string
	Width     *int32
	Height    *int32
	Size      *int64
	Type      *int32
}
