package main

import "fmt"

type Borrowable interface {
	GetTitle() string
	IsAvailable() bool
	UpdateAvailability(status bool)
}

type MemberRepository interface {
	CheckAvailability(memberName string) (bool, *Member)
}

type BorrowableRepository interface {
	CheckAvailability(memberName string) (bool, Borrowable)
}

type BorrowManager struct {
	memberRepo     MemberRepository
	borrowableRepo BorrowableRepository
}

func (bm *BorrowManager) Borrow(memberName string, borrowableName string) {
	hasMember, member := bm.memberRepo.CheckAvailability(memberName)
	hasBorrowable, item := bm.borrowableRepo.CheckAvailability(borrowableName)
	if hasMember && hasBorrowable {
		member.AddItemsBorrowed(item)
		item.UpdateAvailability(false)
	}
}

type BooksManager struct {
	books []*Book
}

func (bm *BooksManager) CheckAvailability(bookName string) (bool, Borrowable) {
	for _, book := range bm.books {
		if bookName == book.Title {
			return true, book
		}
	}
	return false, nil
}

type MembersManager struct {
	members []*Member
}

func (mm *MembersManager) CheckAvailability(memberName string) (bool, *Member) {
	for _, member := range mm.members {
		if memberName == member.Name {
			return true, member
		}
	}
	return false, nil
}

type MagazineManager struct {
	magazines []*Magazine
}

func (mm *MagazineManager) CheckAvailability(titleName string) (bool, Borrowable) {
	for _, magazine := range mm.magazines {
		if titleName == magazine.Title {
			return true, magazine
		}
	}
	return false, nil
}

type Book struct {
	Title     string
	Author    string
	ISBN      string
	Available bool
}

func (b *Book) GetTitle() string               { return b.Title }
func (b *Book) IsAvailable() bool              { return b.Available }
func (b *Book) UpdateAvailability(status bool) { b.Available = status }

type Magazine struct {
	Title     string
	Issue     int
	Available bool
}

func (m *Magazine) GetTitle() string               { return m.Title }
func (m *Magazine) IsAvailable() bool              { return m.Available }
func (m *Magazine) UpdateAvailability(status bool) { m.Available = status }

type Member struct {
	Name          string
	Id            string
	ItemsBorrowed []Borrowable
}

func (m *Member) AddItemsBorrowed(item Borrowable) {
	m.ItemsBorrowed = append(m.ItemsBorrowed, item)
}

func main() {

	books := []*Book{
		{"The Go Programming Language", "Alan A. A. Donovan", "978-0134190440", true},
		{"Clean Code", "Robert C. Martin", "978-0132350884", true},
		{"Design Patterns", "Erich Gamma", "978-0201633610", false},
		{"The Pragmatic Programmer", "Andrew Hunt", "978-0201616223", true},
		{"Refactoring", "Martin Fowler", "978-0201485677", false},
	}
	magazines := []*Magazine{
		{"Tech Today", 42, true},
		{"Science Weekly", 10, false},
	}
	members := []*Member{
		{"Lucas Ferreira", "M001", []Borrowable{}},
		{"Alice Johnson", "M002", []Borrowable{}},
	}

	membersManager := &MembersManager{members: members}
	booksManager := &BooksManager{books: books}
	magazinesManager := &MagazineManager{magazines: magazines}
	borrowManager := &BorrowManager{memberRepo: membersManager, borrowableRepo: booksManager}

	borrowManager.Borrow("Lucas Ferreira", "Clean Code")
	borrowManager.borrowableRepo = magazinesManager
	borrowManager.Borrow("Alice Johnson", "Tech Today")

	for _, member := range membersManager.members {
		fmt.Printf("Member: %s\n", member.Name)
		for _, item := range member.ItemsBorrowed {
			fmt.Printf(" - Borrowed: %s\n", item.GetTitle())
		}
	}
	fmt.Println("\n📚 Available Books:")
	for _, book := range books {
		status := "Unavailable ❌"
		if book.Available {
			status = "Available ✅"
		}
		fmt.Printf(" - %s by %s [%s] - %s\n", book.Title, book.Author, book.ISBN, status)
	}
	fmt.Println("\n📰 Available Magazines:")
	for _, magazine := range magazines {
		status := "Unavailable ❌"
		if magazine.Available {
			status = "Available ✅"
		}
		fmt.Printf(" - %s (Issue %d) - %s\n", magazine.Title, magazine.Issue, status)
	}

}
