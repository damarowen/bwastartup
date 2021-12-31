package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

//JWTService is a contract of what jwtService can do
type IJwtService interface {
	GenerateToken(UserID int, Name string, Email string) string
	ValidateToken(token string) (*jwt.Token, error)
}

type JwtCustomClaim struct {
	UserID int `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	jwt.StandardClaims
}

type jwtService struct {
	issuer    string
	secretKey string
}

//NewJWTService method is creates a new instance of JWTService
func NewJWTService() IJwtService {
	return &jwtService{
		issuer:    "belajargolang",
		secretKey: os.Getenv("JWT_SECRET"),
	}
}



func (j *jwtService) GenerateToken(UserID int, Name string, Email string) string {
	claims := &JwtCustomClaim{
		UserID,
		Name,
		Email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 30).Unix(),
			Issuer:    j.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	//* algoritma 526
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", t_.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})
}


