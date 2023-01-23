package service

import (
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	medapp "github.com/mnogohoddovochka/med-app"
	"github.com/mnogohoddovochka/med-app/pkg/repository"
)

const (
	salt       = "criklgk23ixry23rhzn89"
	signingKey = "ghR#gctErgtyhVzgd#Mnas"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	DoctorId int `json:"doctor_id`
}

type AuthService struct {
	repo repository.Authorisation
}

func NewAuthService(repo repository.Authorisation) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateDoctor(doctor medapp.Doctor) (int, error) {
	doctor.Password = generateHashPassword(doctor.Password)
	return s.repo.CreateDoctor(doctor)
}

func (s *AuthService) GenerateToken(login, password string) (string, error) {
	doctor, err := s.repo.GetDoctor(login, generateHashPassword(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		doctor.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func generateHashPassword(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
