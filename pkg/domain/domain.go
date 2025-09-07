package domain

import (
	internal "github.com/brunojet/infra-go/internal/domain"
	"gorm.io/gorm"
)

const (
	ObjectStatusPending    = internal.ObjectStatusPending
	ObjectStatusProcessing = internal.ObjectStatusProcessing
	ObjectStatusAvailable  = internal.ObjectStatusAvailable
	ObjectStatusError      = internal.ObjectStatusError
	ImageTypeIcon          = internal.ImageTypeIcon
	ImageTypeScreenshot    = internal.ImageTypeScreenshot
	ImageTypeBanner        = internal.ImageTypeBanner
	VideoTypeShorts        = internal.VideoTypeShorts
	VideoTypeReels         = internal.VideoTypeReels
)

type ObjectStatus = internal.ObjectStatus
type VideoType = internal.VideoType
type ImageType = internal.ImageType

type BaseModel = internal.BaseModel
type BaseEntity = internal.BaseEntity
type StorageObject = internal.StorageObject
type Video = internal.Video
type Image = internal.Image

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&StorageObject{}, &Image{}, &Video{})
}
