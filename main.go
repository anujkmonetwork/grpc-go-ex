package main

import (
	"fmt"

	userpb "github.com/anuj070894/go_microservices_new/gen/go/users/v1"
)

func main() {
	user := userpb.User{
		Uuid:          "123",
		FullName:      "Anuj Kumar",
		MaritalStatus: userpb.MaritalStatus_MARITAL_STATUS_MARRIED,
	}

	fmt.Println(user)

	// gender := userpb.Gender{
	// 	GType: "Male",
	// 	// IsStraight: true,
	// }
}
