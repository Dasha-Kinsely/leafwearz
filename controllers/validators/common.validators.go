package validators

import "leafwearz/controllers/serializers"

type AuthValidatorsResources struct {
	Serializers serializers.Serializers
}

func NewAuthValidators(s serializers.Serializers) AuthValidators {
	return &AuthValidatorsResources{
		Serializers: s,
	}
}