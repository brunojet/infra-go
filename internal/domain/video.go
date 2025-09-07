package domain

import "gorm.io/gorm"

type VideoType int16

const (
	VideoTypeShorts VideoType = 0
	VideoTypeReels  VideoType = 10
)

type Video struct {
	BaseModel
	StorageObjectId *int64    `gorm:"uniqueIndex;not null" json:"storage_object_id"`
	VideoType       VideoType `gorm:"not null;check:video_type IN (0, 10)" json:"video_type"`

	//Relationships
	StorageObject StorageObject `gorm:"foreignKey:StorageObjectId;references:ID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
}

func (Video) TableName() string { return "video" }

func (v *Video) AfterDelete(tx *gorm.DB) (err error) {
	if v.StorageObject.ID != 0 {
		tx.Delete(&v.StorageObject)
	}
	return nil
}
