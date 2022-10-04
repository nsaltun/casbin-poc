package main

import (
	"fmt"
	"log"

	"github.com/casbin/casbin/v2"
)

func main() {
	fmt.Println("hello CASBIN!")

	e, err := casbin.NewEnforcer("./config/casbin/rbac_model.conf", "./config/casbin/rbac_policy.csv")
	if err != nil {
		log.Fatalf("NewEnforcer could not be loaded! err:%v", err)
	}

	sub := "bob"   // the user that wants to access a resource.
	obj := "data2" // the resource that is going to be accessed.
	act := "write" // the operation that the user performs on the resource.

	ok, err := e.Enforce(sub, obj, act)

	if err != nil {
		log.Fatalf("NewEnforcer could not be loaded! err:%v", err)
	}

	if ok {
		// permit alice to read data1
		log.Printf("permit %s to %s %s", sub, act, obj)
	} else {
		// deny the request, show an error
		log.Println("deny the request, show an error")
	}
}
