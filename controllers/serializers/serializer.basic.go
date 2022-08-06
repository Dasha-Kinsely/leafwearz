package serializers

import (
	"leafwearz/models/serialization"
	"strings"
)
type Serializers interface {
	ErrorResponse(issues string, message string, data interface{}) serialization.CommonResponse
	NormalResponse(message string, data interface{}) serialization.CommonResponse
}

// If something wrong happens before this response is generated, write errors as a struct obj
func (serializer *SerializersResources) ErrorResponse(issues string, message string, data interface{}) serialization.CommonResponse {
	splittedError := strings.Split(issues, "\n")
	res := serialization.CommonResponse{
		Status: false,
		Message: message,
		Errors: splittedError,
		Data: data,
	}
	return res
}

// This is how a normal, data-carrying response should behave
func (serializer *SerializersResources) NormalResponse(message string, data interface{}) serialization.CommonResponse {
	res := serialization.CommonResponse{
		Status: true,
		Message: message,
		Errors: nil,
		Data: data,
	}
	return res
}