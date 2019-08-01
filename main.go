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
	if _, err := sql.CreateServer(ctx, "meltestserver", "merush", "test123456!!"); err != nil {
		fmt.Println("error is: ", err)
	}

	//create new DB
	if _, err := sql.CreateDB(ctx, "meltestserver", "meltestdb"); err != nil {
		fmt.Println("error is: ", err)
	}

	//delete DB
	if _, err := sql.DeleteDB(ctx, "meltestserver", "meltestdb"); err != nil {
		fmt.Println("error is: ", err)
	}
}
