package main

import (
	"fmt"
	"github.com/beevik/guid"
	"golang-test-2/hash"
	"hash/fnv"
	"time"
)

func HashString(mystring string) string {
	h := fnv.New32a()
	h.Write([]byte(mystring))
	return fmt.Sprint(h.Sum32())
}

func getSecret() interface{} {
	secretKey := "029b937dab2b4e79e24757d5c316b785397b30b6938c71f7ff6d4e7665d0a046"
	return []byte(secretKey)
}

func GetToken(email string) (string, error) {
	today := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username":   email,
		"created-at": today.Format("2006.01.02 15:04:05"),
	})

	signedToken, err := token.SignedString(getSecret())
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func main() {
	// create the token
	/*	token, err := GetToken("reydward@gmail.com")
		if err != nil {
			fmt.Println("Error getting the token")
		}*/
	token := guid.NewString()
	fmt.Println("1. Token created: ", token)

	// hashsalt the token
	tokenHashed, _ := hash.HashKeyWithSalt(token)
	fmt.Println("2. Token hashed", tokenHashed)

	// verify the hashsalt the token
	isValid, _ := hash.CheckHashKey(token, tokenHashed)
	fmt.Println("2.1. Token hashed is valid: ", isValid)

	if !isValid {
		fmt.Println("No!!")
	}

	fmt.Println("Yes!!")

}
