package repos

import "leafwearz/models/entities"

type Creation interface {
	SaveUser(o entities.User) (entities.User, error)
}

type Reading interface {
	FindUserByEmail(email string) (entities.User, error)
}


type Updation interface {

}

type Deletion interface {

}