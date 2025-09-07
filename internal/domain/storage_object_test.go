package domain

import (
	"testing"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func TestStorageObject_TableName(t *testing.T) {
	var s StorageObject
	if s.TableName() != "storage_object" {
		t.Errorf("expected storage_object, got %s", s.TableName())
	}
}

func TestStorageObject_Constraints(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open sqlite: %v", err)
	}
	if err := db.AutoMigrate(&StorageObject{}); err != nil {
		t.Fatalf("failed to migrate: %v", err)
	}
	// Test not null Path
	obj2 := StorageObject{Name: "n2", MimeType: "image/png"}
	if err := db.Create(&obj2).Error; err == nil {
		t.Error("expected error for missing Path (not null)")
	}
	// Test not null Name
	obj3 := StorageObject{Path: "p3", MimeType: "image/png"}
	if err := db.Create(&obj3).Error; err == nil {
		t.Error("expected error for missing Name (not null)")
	}
	// Test not null MimeType
	obj4 := StorageObject{Path: "p4", Name: "n4"}
	if err := db.Create(&obj4).Error; err == nil {
		t.Error("expected error for missing MimeType (not null)")
	}
	// Caso válido
	obj := StorageObject{Path: "p10", Name: "n10", MimeType: "image/png"}
	if err := db.Create(&obj).Error; err != nil {
		t.Fatalf("should insert valid: %v", err)
	}

	// Test unique index Path
	dup := StorageObject{Path: "p10", Name: "n10", MimeType: "image/png"}
	if err := db.Create(&dup).Error; err == nil {
		t.Error("expected error for duplicate Path (unique index)")
	}
}
