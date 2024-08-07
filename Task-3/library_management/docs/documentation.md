# Library Management System

## Overview

This Library Management System is designed to help manage a library's collection of books and its members. It provides functionalities to add and remove books, borrow and return books, and list available and borrowed books.

## Installation

To set up and run the Library Management System, follow these steps:

1. ***Clone the repository:***
    ```sh
    git clone https://github.com/yourusername/library_management.git
    cd library_management
    ```

2. ***Initialize Go modules:***
    ```sh
    go mod init library_management
    go mod tidy
    ```


## Features
+ Add a new book to the library
+ Remove a book from the library
+ Borrow a book for a member
+ Return a borrowed book
+ List all available books
+ List all borrowed books by a member
    

## Usage
It is a console based program that provides differnt options to provide functionalities. These options are:

        1.  Add Book
	 2.  Remove Book
        3.  List Available Books
        4.  Borrow Book
        5.  Return Book
	6.  List of borrowed books by a member
        7.  Exit

#### *These usage are implemented using the following methods:*

#### AddBook
-  Adds a new book to the library's collection.


#### RemoveBook
-  Removes a book from the library's collection.


#### BorrowBook
- Allows a member to borrow a book if it is available.


#### ReturnBook
- Allows a member to return a borrowed book.


#### ListAvailableBooks
- Lists all books that are currently available in the library.


#### ListBorrowedBooks
- Lists all books borrowed by a specific member.

#### Run
- **Description**: Starts the console interaction loop, allowing the user to choose various options to manage the library.

