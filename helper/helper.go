package helper

import (
	"bwastartup/user"
	"fmt"
	"strings"
)

type Response struct {
	Meta Meta `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Errors  interface{} `json:"errors"`
}

func ApiResponse(status bool, message string, code int, data interface{}, err interface{}) Response {

	//normal
	//var errors []string
	//for _,e := range err.(validator.ValidationErrors){
	//	errors = append(errors, e.Error())
	//}
	//errMessage := gin.H{"ERROR": errors}
	//

	fmt.Printf("type: %T\n", err)

	var dataErr  interface{}
	//logic error = " jika tidak ada error
	if len(err.(string)) > 0 {
		dataErr = strings.Split(err.(string), "\n")
	} else {
		dataErr = ""
	}

	Meta := Meta{
		Status:  status,
		Message: message,
		Code:    code,
		Errors:  dataErr,
	}

	jsonResponse := Response{
		Meta: Meta,
		Data: data,
	}

	return jsonResponse
}

func MappingResponseUser(u user.User, token string) user.DtoResponseUser {

	return user.DtoResponseUser{
		ID:         u.ID,
		Name:       u.Name,
		Occupation: u.Occupation,
		Email:      u.Email,
		Token:      token,
	}

}
