package main

func main() {
	app := App{}

	app.Initialize(DBUser, DBPassword, DBHost, DBName)
	app.Run("localhost:10000")
}
