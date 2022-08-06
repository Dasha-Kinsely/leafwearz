package repos

import "leafwearz/models/entities"

func (database *ReadingResources) FindUserByEmail(email string) (entities.User, error) {
	var user entities.User
	err := database.Database.Where("email = ?", email).Take(&user)
	if err.Error != nil {
		return entities.User{}, err.Error
	} else {
		return user, nil
	}
}