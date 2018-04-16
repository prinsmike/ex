package main

import (
	"fmt"
	"log"
	"os/user"
)

func main() {
	u, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("User: %s\n", u.Username)
	fmt.Printf("UID: %s\n", u.Uid)
	fmt.Printf("GID: %s\n", u.Gid)
	fmt.Printf("Home: %s\n", u.HomeDir)
}
