package gostagram

import (
	"fmt"
)

// prints current user id, username and fullname.
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

// prints instagram tags with
// golang query
func ExampleClient_secure() {
	client := NewClient("access_token")
	client.SetSignedRequest(true)
	client.SetClientSecret("client secret")


	tags, err := client.SearchTags("golang")

	if err != nil {
		fmt.Println(err)
	} else {
		for _, tag := range tags {
			fmt.Println("Tag name: ", tag.Name)
		}
	}
}
