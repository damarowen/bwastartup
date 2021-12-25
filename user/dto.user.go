package user

type DtoResponseUser struct {
	ID         int   ` json:"id"`
	Name       string ` json:"name" `
	Occupation string    ` json:"occupation"`
	Email      string ` json:"email"`
	Token      string    ` json:"token" `
}

type DtoRegisterUserInput struct {
	Name       string
	Occupation string
	Email      string
	Password   string
}