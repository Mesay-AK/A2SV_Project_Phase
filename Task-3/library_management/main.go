package main

import (
    "library_management/controllers"
    "library_management/services"
)

func main() {
    // Create a new Library instance
    library := services.NewLibrary()

    // Create a LibraryConsole with the Library instance
    console := controllers.NewLibraryConsole(library) // Use the constructor function

    // Start the console interaction loop
    console.Run()
}
