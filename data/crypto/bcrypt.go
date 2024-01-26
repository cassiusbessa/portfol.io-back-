package bcrypt

import (
	usecases "github.com/Grupo-38-Orange-Juice/orange-portfolio-back/domain/use-cases"
	"golang.org/x/crypto/bcrypt"
)

type BcryptCrypto struct{}

func NewBcrypt() usecases.Crypto {
	return &BcryptCrypto{}
}

func (b *BcryptCrypto) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (b *BcryptCrypto) CompareHashAndPassword(hash string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
