package user

type DtoResponseUser struct {
	ID         int   ` json:"id"`
	Name       string ` json:"name" `
	Occupation string    ` json:"occupation"`
	Email      string ` json:"email"`
	Token      string    ` json:"token" `
}

type DtoRegisterUserInput struct {
	Name       string `json:"name" binding:"required"`
	Occupation string `json:"occupation" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
}

type DtoLoginUserInput struct {
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
}

type DtoEmailChecker struct {
	Email      string `json:"email" binding:"required,email"`
}