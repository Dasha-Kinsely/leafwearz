package repos

import "gorm.io/gorm"

type CreationResources struct {
	Database *gorm.DB
}

func NewCreationRepo(db *gorm.DB) Creation {
	return &CreationResources{
		Database: db,
	}
}

type ReadingResources struct {
	Database *gorm.DB
}

func NewReadingRepo(db *gorm.DB) Reading {
	return &ReadingResources{
		Database: db,
	}
}

type UpdationResources struct {
	Database *gorm.DB
}

func NewUpdationRepo(db *gorm.DB) Updation {
	return &UpdationResources{
		Database: db,
	}
}

type DeletionResources struct {
	Database *gorm.DB
}

func NewDeletionRepo(db *gorm.DB) Deletion {
	return &DeletionResources{
		Database: db,
	}
}