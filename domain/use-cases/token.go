package usecases

type Token interface {
	GenerateToken(payload string) (string, error)
	CheckToken(token string) (bool, error)
}
