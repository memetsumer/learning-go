package main

func main() {
    const (
		host     = "localhost"
		port     = 5433
		user     = "mehmet"
		password = "mehmet"
		dbname   = "mehmet"
	)

    db := connectDB(host, port, user, password, dbname)

    insertDB("sayko", "memo", db)
}
