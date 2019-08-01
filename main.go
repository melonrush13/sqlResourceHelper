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
	// _, err := sql.CreateServer(ctx, "meltestserver", "merush", "test123456!!")

	//create new DB
	_, err := sql.CreateDB(ctx, "meltestserver", "meltestdb")

	//delete a DB
	// _, err := sql.DeleteDB(ctx, "meltestserver", "meltestdb")

	if err != nil {
		fmt.Println("error is: ", err)
	}
}
