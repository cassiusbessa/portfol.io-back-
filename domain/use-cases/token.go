package usecases

type Token interface {
	GenerateToken(payload string) (string, error)
	GetPayload(token string) (string, error)
}
