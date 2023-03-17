package main

import (
	"fmt"
	"github.com/satori/uuid"
)

func main() {
	u1 := uuid.NewV4()
	fmt.Printf("uuidv4: %s\n", u1)
}
