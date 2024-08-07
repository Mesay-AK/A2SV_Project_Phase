package controllers

import (
    "bufio"
    "fmt"
    "library_management/services"
    "library_management/models"
    "os"
    "strconv"
)

type LibraryConsole struct {
    Library *services.Library
}

// NewLibraryConsole initializes a LibraryConsole with a Library instance
func NewLibraryConsole(lib *services.Library) *LibraryConsole {
    return &LibraryConsole{
        Library: lib,
    }
}

// Run starts the console interaction loop
func (console *LibraryConsole) Run() {
    scanner := bufio.NewScanner(os.Stdin)

    for {
        fmt.Println("Choose an option:")
        fmt.Println("1. Add Book")
		fmt.Println("2. Remove Book")
        fmt.Println("3. List Available Books")
        fmt.Println("4. Borrow Book")
        fmt.Println("5. Return Book")
		fmt.Println("6. List of borrowed books by a member")
        fmt.Println("7. Exit")

        scanner.Scan()
        option := scanner.Text()

        switch option {
        case "1":
            console.addBook(scanner)
        case "2":
			console.removeBook(scanner)
		case "3":
            console.listAvailableBooks()
        case "4":
            console.borrowBook(scanner)
        case "5":
            console.returnBook(scanner)
		case "6":
			console.listBorrowedBooks(scanner)
        case "7":
            fmt.Println("Exiting...")
            return
        default:
            fmt.Println("Invalid option. Please try again.")
        }
    }
}

// Prompts the user to provide the book detail and calls the Add method 
func (console *LibraryConsole) addBook(scanner *bufio.Scanner) {

    fmt.Print("Enter Book ID: ")
    scanner.Scan()

    id, _ := strconv.Atoi(scanner.Text())
    fmt.Print("Enter Book Title: ")
    scanner.Scan()

    title := scanner.Text()
    book := models.Book{ID: id, Title: title, Status: "Available"}
    console.Library.AddBook(book)

    fmt.Println("Book added successfully")
}

// Prompts the user to provide the Id of the book to be removed and calls the Remove method 
func (console *LibraryConsole) removeBook(scanner *bufio.Scanner) {
    fmt.Print("Enter Book ID to Remove: ")
    scanner.Scan()
    bookID, _ := strconv.Atoi(scanner.Text())
    
    change :=console.Library.RemoveBook(bookID)
    if change == nil{
        fmt.Println("Book removed successfully")
    }else{
        fmt.Println("Error :",change)
    }
    
    
}

// calls the ListAvailableBooks and desplays the available books
func (console *LibraryConsole) listAvailableBooks() {
    books := console.Library.ListAvailableBooks()
    fmt.Println("Available Books:")

	if len(books) > 0 {
		for _, book := range books {
        fmt.Printf("ID: %d, Title: %s\n", book.ID, book.Title)
    }
	}else{
		fmt.Println("There is no book currently Available.")
	}

}
// Prompts the user to provide the book detail and calls the BorrowBook method 
func (console *LibraryConsole) borrowBook(scanner *bufio.Scanner) {

    fmt.Print("Enter Book ID to Borrow: ")
    scanner.Scan()
    bookID, _ := strconv.Atoi(scanner.Text())

    fmt.Print("Enter Member ID: ")
    scanner.Scan()
    memberID, _ := strconv.Atoi(scanner.Text())

    err := console.Library.BorrowBook(bookID, memberID)
    if err != nil {
        fmt.Println("Error borrowing book:", err)
    } else {
        fmt.Println("Book borrowed successfully")
    }
}

// prompts the user to provide bookId and memberId to return and calls the ReturnBook method
func (console *LibraryConsole) returnBook(scanner *bufio.Scanner) {

    fmt.Print("Enter Book ID to Return: ")
    scanner.Scan()
    bookID, _ := strconv.Atoi(scanner.Text())

    fmt.Print("Enter Member ID: ")
    scanner.Scan()
    memberID, _ := strconv.Atoi(scanner.Text())

    err := console.Library.ReturnBook(bookID, memberID)
    // fmt.Print(err)
    if err != nil {
        fmt.Println("Error returning book: \n", err)
    } else {
        fmt.Println("Book returned successfully")
    }
}
  
// Prompts user to provide the member Id and calls ListBorrowedBooks and Prints the list
func (console *LibraryConsole) listBorrowedBooks(scanner *bufio.Scanner) {
    fmt.Print("Enter Member ID: ")
    scanner.Scan()
    memberID, _ := strconv.Atoi(scanner.Text())
    
    books := console.Library.ListBorrowedBooks(memberID)
    fmt.Println("Borrowed Books:")
    if len(books)> 0 {
        for _, book := range books {
        fmt.Printf("ID: %d, Title: %s\n", book.ID, book.Title)
    }
    }else{
        fmt.Println("There is no book borrowed under this ID.")

    }
}
