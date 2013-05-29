package main

import (
	"fmt"
)

func install_package(pkg string) {
	fmt.Println("Looking for install package")
	err := db_init()
	if err != nil {
		panic(err)
	}
	defer db_close()

	fmt.Println("Installation complete")
}
