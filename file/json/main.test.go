package main

import (
	//"log"
	"fmt"

	project "./pkg"
)

func main() {
	cfg := project.LoadConfiguration("./goerli.json")

	fmt.Println(cfg.Database.Host)
	fmt.Println(cfg.Host)
	fmt.Println(cfg.Port)
}
