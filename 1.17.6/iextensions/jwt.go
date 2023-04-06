package iextensions

import (
	"log"

	jwt "github.com/golang-jwt/jwt/v4"
)

func TestIExtensionsJwtFuncs(tokenStr string) {
	// ? `token.Valid`
	// i did it but failed because token of casdoor in ci-auto-test expire at once.

	parser := jwt.NewParser(jwt.WithoutClaimsValidation())

	token, _, _ := parser.ParseUnverified(tokenStr, jwt.MapClaims{})

	log.Printf("%+v\n", token)
	log.Printf("%+v\n", token.Valid)

	// ? map 打印
	log.Println(token.Claims.(jwt.MapClaims))
}
