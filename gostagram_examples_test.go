package gostagram

import (
	"fmt"
)

func ExampleClient() {
	client := NewClient("access_token")
	user, err := client.GetCurrentUser()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(user.Id)
		fmt.Println(user.Username)
		fmt.Println(user.FullName)
	}
}

func ExampleClient_secure() {
	client := NewClient("access_token")
	client.SetSignedRequest(true)
	client.SetClientSecret("client secret")

	user, err := client.GetCurrentUser()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(user.Id)
		fmt.Println(user.Username)
		fmt.Println(user.FullName)
	}
}
