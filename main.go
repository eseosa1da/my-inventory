package main

func main() {
	app := App{}

	app.Initialize(DbUser, DbPassword, DBName)
	app.Run("localhost:10000")
}
