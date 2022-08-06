package repos

import (
	"leafwearz/models/entities"
	"leafwearz/pkg/encryption"
)

func (db *CreationResources) SaveUser(o entities.User) (entities.User, error) {
	o.Password = encryption.EncryptPassword(o.Password)
	if err := db.Database.Save(&o).Error; err != nil {
		return o, err
	}
	linkedProfile := entities.UserExtras{User: o}
	if err := db.Database.Table("user_extras").Create(&linkedProfile).Error; err != nil {
		return o, err
	}
	return o, nil
}