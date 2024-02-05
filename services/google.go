package services

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"
)

type GoogleUserInfo struct {
	Iss           string `json:"iss"`
	Azp           string `json:"azp"`
	Aud           string `json:"aud"`
	Sub           string `json:"sub"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Exp           int64  `json:"exp"`
	FamilyName    string `json:"family_name"`
	GivenName     string `json:"given_name"`
	Iat           int64  `json:"iat"`
	Jti           string `json:"jti"`
	Locale        string `json:"locale"`
	Name          string `json:"name"`
	Nbf           int64  `json:"nbf"`
	Picture       string `json:"picture"`
}

type GoogleService struct {
}

func NewGoogleService() *GoogleService {
	return &GoogleService{}
}

func printStruct(s interface{}) {
	value := reflect.ValueOf(s)
	typ := reflect.TypeOf(s)

	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		fieldName := typ.Field(i).Name

		fmt.Printf("%s: %v\n", fieldName, field.Interface())
	}
}

func (g *GoogleService) validUser(u GoogleUserInfo) (bool, error) {
	if !u.EmailVerified {
		return false, fmt.Errorf("Email não verificado")
	}
	if u.Email == "" {
		return false, fmt.Errorf("Email não encontrado")
	}
	if u.Name == "" {
		return false, fmt.Errorf("Nome não encontrado")
	}
	if u.Exp < time.Now().Unix() {
		return false, fmt.Errorf("Token expirado")
	}
	return true, nil
}

func (g *GoogleService) GetUserPayload(token string) (*GoogleUserInfo, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return nil, fmt.Errorf("token JWT inválido")
	}

	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, fmt.Errorf("falha ao decodificar payload do token JWT: %v", err)
	}

	var claims map[string]interface{}
	if err := json.Unmarshal(payload, &claims); err != nil {
		return nil, fmt.Errorf("falha ao deserializar payload do token JWT: %v", err)
	}

	var userInfo GoogleUserInfo

	if sub, ok := claims["sub"].(string); ok {
		userInfo.Sub = sub
	}

	if email, ok := claims["email"].(string); ok {
		userInfo.Email = email
	}

	if emailVerified, ok := claims["email_verified"].(bool); ok {
		userInfo.EmailVerified = emailVerified
	}

	if name, ok := claims["name"].(string); ok {
		fmt.Println("name", name)
		userInfo.Name = name
	}

	if picture, ok := claims["picture"].(string); ok {
		userInfo.Picture = picture
	}

	if givenName, ok := claims["given_name"].(string); ok {
		userInfo.GivenName = givenName
	}

	if exp, ok := claims["exp"].(float64); ok {
		userInfo.Exp = int64(exp)
	}

	if _, err := g.validUser(userInfo); err != nil {
		return nil, err
	}

	return &userInfo, nil
}
