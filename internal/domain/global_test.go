package domain_test

import (
	"testing"

	"github.com/brunojet/infra-go/pkg/domain"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func SetupDomainTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		t.Fatalf("failed to open sqlite: %v", err)
	}
	if err := db.Exec("PRAGMA foreign_keys = ON").Error; err != nil {
		t.Fatalf("failed to enable foreign keys: %v", err)
	}
	if err := domain.AutoMigrate(db); err != nil {
		t.Fatalf("failed to migrate: %v", err)
	}
	return db
}

func TestBaseEntity_BeforeCreate_Valid(t *testing.T) {
	e := &domain.BaseEntity{Name: "ok", Active: true}
	// pass nil DB (method only validates Nome length)
	if err := e.BeforeCreate(nil); err != nil {
		t.Fatalf("expected no error for valid Nome, got: %v", err)
	}
}

func TestBaseEntity_BeforeCreate_Invalid(t *testing.T) {
	e := &domain.BaseEntity{Name: "", Active: true}
	if err := e.BeforeCreate(&gorm.DB{}); err == nil {
		t.Fatalf("expected error for empty Nome, got nil")
	}
}
