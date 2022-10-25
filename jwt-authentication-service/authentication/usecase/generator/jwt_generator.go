package generator

import "github.com/golang-jwt/jwt"

type JwtGenerator struct {
	secretKeySample []byte
}

func NewGenerator(secretKey string) *JwtGenerator {
	return &JwtGenerator{
		secretKeySample: []byte(secretKey),
	}
}

func (g *JwtGenerator) GenerateJWT(claims map[string]interface{}) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	g.addClaims(token, claims)

	tokenString, err := token.SignedString(g.secretKeySample)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (g *JwtGenerator) addClaims(token *jwt.Token, c map[string]interface{}) {
	claims := token.Claims.(jwt.MapClaims)
	for k, v := range c {
		claims[k] = v
	}
}
