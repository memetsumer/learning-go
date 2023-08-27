package main

import (
	"fmt"

	"github.com/joho/godotenv"
)

func main() {

    err := godotenv.Load()
    if  err != nil {
        fmt.Println("Error loading .env file: ", err)
    }

	// const (
	// 	host     = "localhost"
	// 	port     = 5433
	// 	user     = "mehmet"
	// 	password = "mehmet"
	// 	dbname   = "mehmet"
	// )

	// db := connectDB(host, port, user, password, dbname)

	// insertDB("sayko", "memo", db)

	// run()

    RunWeather();
}
