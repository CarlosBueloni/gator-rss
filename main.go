package main

import (
	"fmt"
	"github.com/carlosbueloni/gator-rss/internal/config"
)

func main() {
	config_content := config.Read()
	fmt.Println(config_content.CurrentUserName)
	fmt.Println(config_content.DbURL)

	config.SetCurrentUser("johhn")

	config_content = config.Read()
	fmt.Println(config_content.CurrentUserName)
	fmt.Println(config_content.DbURL)

}
