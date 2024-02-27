package main

import (
	"fmt"
	"log"

	userpb "github.com/anuj070894/go_microservices_new/gen/go/users/v1"

	"github.com/bufbuild/protovalidate-go"
)

func main() {
	user := &userpb.User{
		FullName:  "Anuj Kumar",
		BirthYear: 1997,
	}
	v, err := protovalidate.New()

	if err != nil {
		log.Fatalln("Error protovalidate", err)
	}

	if err := v.Validate(user); err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("OK")
	}
}
