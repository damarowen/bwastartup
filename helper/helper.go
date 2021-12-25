package helper

import (
	"bwastartup/user"
)

type Response struct {
	Meta Meta
	Data interface{} `json:"data"`
}

type Meta struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Errors  interface{} `json:"errors"`
}

func ApiResponse(status bool, message string, code int, data interface{}, err interface{}) Response {

	Meta := Meta{
		Status:  status,
		Message: message,
		Code:    code,
		Errors:  err,
	}

	jsonResponse := Response{
		Meta: Meta,
		Data: data,
	}

	return jsonResponse
}

//func ApiErrorResponse(message string, err string, data interface{}) Response {
//	fmt.Println(err, "<<<<<<<<<")
//	splittedError := strings.Split(err, "\n")
//	res := Response{
//		Status:  false,
//		Message: message,
//		Errors:  splittedError,
//		Data:    data,
//	}
//	return res
//}

func MappingResponseUser(u user.User, token string) user.DtoResponseUser {

	return user.DtoResponseUser{
		ID:         u.ID,
		Name:       u.Name,
		Occupation: u.Occupation,
		Email:      u.Email,
		Token:      token,
	}

}
