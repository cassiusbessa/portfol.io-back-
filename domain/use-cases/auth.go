package usecases

type Crypto interface {
	HashPassword(password string) (string, error)
	CompareHashAndPassword(hash string, password string) error
}
