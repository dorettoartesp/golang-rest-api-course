package main

import "fmt"

// App - the struct which contains things like pointers
// to database connections
type App struct {
}

// Run - sets up our application
func (app *App) Run() error {
	fmt.Println("Setting up our APP")
	return nil
}

func main() {

	fmt.Println("Go REST API course")
	app := App{}

	if err := app.Run(); err != nil {
		fmt.Println("Error starting our res API")
		fmt.Println(err)
	}

}
