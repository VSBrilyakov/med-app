package service

import (
	"crypto/sha1"
	"fmt"

	medapp "github.com/mnogohoddovochka/med-app"
	"github.com/mnogohoddovochka/med-app/pkg/repository"
)

const salt = "criklgk23ixry23rhzn89"

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

func generateHashPassword(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
