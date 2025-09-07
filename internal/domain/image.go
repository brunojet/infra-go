package domain

import "gorm.io/gorm"

type ImageType int16

const (
	ImageTypeIcon       ImageType = 0
	ImageTypeScreenshot ImageType = 10
	ImageTypeBanner     ImageType = 20
)

type Image struct {
	BaseModel
	StorageObjectId *int64    `gorm:"uniqueIndex;not null" json:"storage_object_id"`
	ImageType       ImageType `gorm:"index;not null;check:image_type IN (0, 10, 20)" json:"image_type"`

	//Relationships
	StorageObject StorageObject `gorm:"foreignKey:StorageObjectId;references:ID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
}

func (Image) TableName() string { return "image" }

func (i *Image) AfterDelete(tx *gorm.DB) (err error) {
	if i.StorageObject.ID != 0 {
		tx.Delete(&i.StorageObject)
	}
	return nil
}
