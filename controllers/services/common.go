package services

import "leafwearz/controllers/repos"

type AuthServicesResources struct {
	InsertionRepos repos.Creation
	FetchingRepos repos.Reading
	EditionRepos repos.Updation
}

func NewAuthServices(
	insertionRepos repos.Creation, 
	fetchingRepos repos.Reading,
	editionRepos repos.Updation) AuthServices {
		return &AuthServicesResources{
			InsertionRepos: insertionRepos,
			FetchingRepos: fetchingRepos,
			EditionRepos: editionRepos,
		}
}