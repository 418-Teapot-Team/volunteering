package main

import (
	"fmt"
	"github.com/BoryslavGlov/logrusx"
	"log"
)

func main() {

	logg, err := logrusx.New("budget-tracker")
	if err != nil {
		log.Fatal("error while trying to create logg instance")
	}

	fmt.Println(logg)
}
