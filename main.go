package main

import (
	"fmt"

	"github.com/Padliwinata/iam-sdk/iam"
)

func main() {
	token := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJodHRwOi8vb3NzLklELmRvbWFpbi5jb20iLCJzdWIiOiJkb21haW58NjRjY2E3ZDM0YmI0OGI2MDhkM2JkZjM2IiwiYXVkIjpbImh0dHA6Ly9sb2NhbGhvc3Q6ODA4MSIsImh0dHA6Ly9sb2NhbGhvc3Q6ODA4MSIsImh0dHA6Ly9sb2NhbGhvc3Q6ODA4MS9sb2dpbi1zdWNjZXNzIl0sImF6cCI6IjY0Y2NhNjYzNGJiNDhiNjA4ZDNiZGYzMSIsImV4cCI6MTY5NDU4NTI3NiwiaWF0IjoxNjkxOTkzMjc2LCJzY29wZXMiOiJlbWFpbCBvcGVuaWQgcHJvZmlsZSB1c2VyOmNyZWF0ZSB1c2VyOnJlYWQifQ.fMdGyGwjAJncvJYkn3FlFWSv6-T0Y38gLibfbedIS9M"
	secretKey := "$2b$04$VFIar.GWpZXLQqLk3sVoEehKdaHuU2JJoY6j5J.2g9AsHZFR8SkAu"

	claims, err := iam.DecodeWithSecret(token, secretKey)
	if err != nil {
		fmt.Println("Error decoding:", err)
	} else {
		fmt.Println("Token berhasil di decode!")
		fmt.Println("Claims:", claims)
	}

	isAuthenticated := iam.CheckAuth(token, secretKey)
	fmt.Println("Is authenticated:", isAuthenticated)

	if scope, ok := claims["scopes"].(string); ok {
		hasPermission := iam.CheckPermission(token, secretKey, scope)
		fmt.Println("Has permission:", hasPermission)
	}

}
