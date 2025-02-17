# Library Management System in Go

This is a simple library management system implemented in Go. The system allows members to borrow books and magazines, and it keeps track of the availability of these items. The code is structured using interfaces and structs to manage books, magazines, and members.

## Features

- **Borrowable Items**: Books and magazines can be borrowed by members.
- **Member Management**: Members can borrow items, and their borrowed items are tracked.
- **Availability Tracking**: The system tracks the availability of books and magazines.
- **Flexible Repository System**: The system uses interfaces to manage different types of borrowable items and members, making it easy to extend.

## Code Structure

The code is organized into several components:

1. **Interfaces**:
   - `Borrowable`: Defines methods for borrowable items (e.g., books, magazines).
   - `MemberRepository`: Defines methods for checking member availability.
   - `BorrowableRepository`: Defines methods for checking borrowable item availability.

2. **Structs**:
   - `BorrowManager`: Manages the borrowing process by coordinating between member and borrowable repositories.
   - `BooksManager`: Manages a collection of books and implements the `BorrowableRepository` interface.
   - `MembersManager`: Manages a collection of members and implements the `MemberRepository` interface.
   - `MagazineManager`: Manages a collection of magazines and implements the `BorrowableRepository` interface.
   - `Book`: Represents a book with properties like title, author, ISBN, and availability.
   - `Magazine`: Represents a magazine with properties like title, issue number, and availability.
   - `Member`: Represents a library member with properties like name, ID, and a list of borrowed items.

3. **Main Function**:
   - Initializes the system with some sample data (books, magazines, and members).
   - Demonstrates borrowing items and prints the current state of the library.

## Usage

### Running the Code

To run the code, simply execute the following command in your terminal:

```bash
go run main.go
```

### Sample Output

The program will output the current state of the library, including which members have borrowed which items and the availability of books and magazines.

```plaintext
Member: Lucas Ferreira
 - Borrowed: Clean Code
Member: Alice Johnson
 - Borrowed: Tech Today

📚 Available Books:
 - The Go Programming Language by Alan A. A. Donovan [978-0134190440] - Available ✅
 - Clean Code by Robert C. Martin [978-0132350884] - Unavailable ❌
 - Design Patterns by Erich Gamma [978-0201633610] - Unavailable ❌
 - The Pragmatic Programmer by Andrew Hunt [978-0201616223] - Available ✅
 - Refactoring by Martin Fowler [978-0201485677] - Unavailable ❌

📰 Available Magazines:
 - Tech Today (Issue 42) - Unavailable ❌
 - Science Weekly (Issue 10) - Unavailable ❌
```

## Extending the System

The system is designed to be easily extendable. You can add new types of borrowable items (e.g., DVDs, CDs) by implementing the `Borrowable` interface and creating a corresponding manager struct that implements the `BorrowableRepository` interface.

### Example: Adding a DVD Manager

```go
type DVD struct {
    Title     string
    Director  string
    Available bool
}

func (d *DVD) GetTitle() string               { return d.Title }
func (d *DVD) IsAvailable() bool              { return d.Available }
func (d *DVD) UpdateAvailability(status bool) { d.Available = status }

type DVDManager struct {
    dvds []*DVD
}

func (dm *DVDManager) CheckAvailability(titleName string) (bool, Borrowable) {
    for _, dvd := range dm.dvds {
        if titleName == dvd.Title {
            return true, dvd
        }
    }
    return false, nil
}
```

Then, you can add DVDs to the system and update the `BorrowManager` to use the `DVDManager` as needed.

## Conclusion

This library management system is a simple yet powerful example of how to use Go's interfaces and structs to create a flexible and extendable system. It demonstrates key concepts like encapsulation, polymorphism, and separation of concerns, making it a great starting point for more complex applications.