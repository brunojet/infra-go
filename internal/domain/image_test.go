package domain_test

import (
	"testing"

	"github.com/brunojet/infra-go/pkg/domain"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func TestImage_TableName(t *testing.T) {
	var i domain.Image
	if i.TableName() != "image" {
		t.Errorf("expected image, got %s", i.TableName())
	}
}

func TestImage_Constraints(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open sqlite: %v", err)
	}
	if err := db.AutoMigrate(&domain.Image{}, &domain.StorageObject{}); err != nil {
		t.Fatalf("failed to migrate: %v", err)
	}
	// Test not null
	img := domain.Image{StorageObjectId: nil, ImageType: 0}
	if err := db.Create(&img).Error; err == nil {
		t.Error("expected error for missing not null fields")
	}
}

func TestImage_StorageObject_Constraints(t *testing.T) {
	db := SetupDomainTestDB(t)
	if db == nil {
		t.Fatal("failed to setup test db")
	}
	status := domain.ObjectStatusAvailable
	// Cria Image referenciando StorageObject
	img := domain.Image{
		StorageObject: domain.StorageObject{Path: "p1", Name: "n1", MimeType: "image/png", Status: &status},
	}
	if err := db.Create(&img).Error; err != nil {
		t.Fatalf("failed to create image: %v", err)
	}
	// Tenta deletar StorageObject (deve falhar por constraint)
	if err := db.Delete(&img.StorageObject).Error; err == nil {
		t.Error("expected error when deleting StorageObject with Image referencing it")
	}

	if err := db.Delete(&img).Error; err != nil {
		t.Fatalf("failed to delete image: %v", err)
	}

	var so domain.StorageObject
	if err := db.First(&so, img.StorageObject.ID).Error; err == nil {
		t.Error("expected storageobject to be deleted, but found one")
	}
}
