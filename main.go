package main

func main() {
	app := App{}

	app.Initialize(DbUser, DbPassword, DBHost, DBName)
	app.Run("localhost:10000")
}
