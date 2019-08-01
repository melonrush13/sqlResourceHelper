package main

import (
	"context"
	"fmt"

	"github.com/melonrush13/sqlResourceHelper/config"
	"github.com/melonrush13/sqlResourceHelper/sql"
)

func main() {
	fmt.Println("hi")

	config.LoadSettings()
	ctx := context.Background()

	//create new server
	_, err := sql.CreateServer(ctx, "melRushServer2019", "merush", "test123456!!")

	if err != nil {
		fmt.Println("error is: ", err)
	}
}
