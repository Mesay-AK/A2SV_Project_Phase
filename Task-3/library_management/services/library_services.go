package services

import(
	"library_management/models"
	"fmt"
)


type Library struct{
	BookList map[int]models.Book
	MemberList map[int]models.Member
}

func NewLibrary() *Library {
    return &Library{
        BookList:   make(map[int]models.Book),
        MemberList: make(map[int]models.Member),
    }
}

type LibraryManager interface {
	AddBook(book models.Book)
	RemoveBook(bookID int) error
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book
	checkExistance(bookID int, memberID int)bool
}

func (Lib *Library)AddBook(book models.Book){
	Lib.BookList[book.ID] = book
}

func (Lib *Library)RemoveBook(bookID int)error{
	if len(Lib.BookList) > 0{
		_, book := Lib.BookList[bookID]
		if book {
		delete(Lib.BookList, bookID)
		return nil
		}

	}
	return fmt.Errorf("the book is not found")
}

func (Lib *Library)BorrowBook(bookID int, memberID int)error{

	if Lib.checkExistance(bookID, memberID){
		book:= Lib.BookList[bookID]
		member:= Lib.MemberList[memberID]

	// Checking the availability of the book (Available or borrowed)
	if book.Status != "Available" {
	 	return fmt.Errorf("book with ID %d is not available", bookID)
	 
}

	//Changing the status of the book 
	book.Status = "Borrowed"
	Lib.BookList[bookID] = book

	// Adding book to member's borrowed list
	member.BorrowedBooks = append(member.BorrowedBooks, book)
	Lib.MemberList[memberID] = member


	}else{
		return fmt.Errorf("either member or the book doesn't exist")
	}
	
	return nil
}

func (Lib *Library)ReturnBook(bookID int, memberID int)error{

	if Lib.checkExistance(bookID, memberID){
		book:= Lib.BookList[bookID]
		member:= Lib.MemberList[memberID]

		index := -1
		for i, Current_Book := range member.BorrowedBooks{
			// cheking if the book is borrowed by the member
			if Current_Book.ID == bookID{
				index = i
				break
			}
		}

		if index != -1{
			return fmt.Errorf("book with ID %d is not borrowed by member %d", bookID, memberID)
				
		}
		book.Status = "Available"
		Lib.BookList[bookID] = book

		// Removing book from member's borrowed list
		member.BorrowedBooks = append(member.BorrowedBooks[:index], member.BorrowedBooks[index+1:]...)
		Lib.MemberList[memberID] = member

	
	}else{
		return fmt.Errorf("either member or the book doesn't exist")
	}
	return nil
	
}

func (Lib *Library)ListAvailableBooks() []models.Book{
	var Check_Status = Lib.BookList

	var Found []models.Book

	for _, Current_Book := range Check_Status{

		// checking if the book is avaliable
		if Current_Book.Status == "Available"{
			Found= append(Found, Current_Book)
		}

	}
	return Found
}

func (Lib *Library)ListBorrowedBooks(memberID int) []models.Book{

	// Returning the list of borrowed books
	member := Lib.MemberList[memberID]

	return member.BorrowedBooks

}

func (Lib *Library)checkExistance(bookID int, memberID int)bool{

	_, ok := Lib.BookList[bookID]
	_, exists := Lib.MemberList[memberID]

	// Checking if the person is a memeber

	if !exists {
		fmt.Print(fmt.Errorf("member with ID %d does not exist", memberID))
		return false
	}
	// Checking whether the book exists in the library list

	if !ok {
		fmt.Print(fmt.Errorf("book with ID %d does not exist", bookID))
		return false
	}
	return true

}
