package domain_test

import (
	"testing"

	"github.com/brunojet/infra-go/pkg/domain"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func TestVideo_TableName(t *testing.T) {
	var v domain.Video
	if v.TableName() != "video" {
		t.Errorf("expected video, got %s", v.TableName())
	}
}

func TestVideo_Constraints(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open sqlite: %v", err)
	}
	if err := db.AutoMigrate(&domain.Video{}, &domain.StorageObject{}); err != nil {
		t.Fatalf("failed to migrate: %v", err)
	}
	// Test not null
	vid := domain.Video{StorageObjectId: nil}
	if err := db.Create(&vid).Error; err == nil {
		t.Error("expected error for missing not null fields")
	}
}

func TestVideo_StorageObject_Constraints(t *testing.T) {
	db := SetupDomainTestDB(t)
	if db == nil {
		t.Fatal("failed to setup test db")
	}
	status := domain.ObjectStatusAvailable
	// Cria Video referenciando StorageObject
	vid := domain.Video{
		StorageObject: domain.StorageObject{Path: "p1", Name: "n1", MimeType: "video/mp4", Status: &status},
	}
	if err := db.Create(&vid).Error; err != nil {
		t.Fatalf("failed to create video: %v", err)
	}
	// Tenta deletar StorageObject (deve falhar por constraint)
	if err := db.Delete(&vid.StorageObject).Error; err == nil {
		t.Error("expected error when deleting StorageObject with Video referencing it")
	}

	if err := db.Delete(&vid).Error; err != nil {
		t.Fatalf("failed to delete video: %v", err)
	}

	var so domain.StorageObject
	if err := db.First(&so, vid.StorageObject.ID).Error; err == nil {
		t.Error("expected storageobject to be deleted, but found one")
	}
}
