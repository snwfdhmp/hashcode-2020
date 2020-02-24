package main

// import "fmt"

type Book struct {
	Score int
	ID    int
}

type Library struct {
	ID          int
	Books       []Book
	SignupTime  int
	BooksPerDay int
}

func (l *Library) BookValueSum(numberOfDays int) int {
	numberOfDays -= l.SignupTime
	sum := 0
	for i := 0; i < l.BooksPerDay*numberOfDays && i < len(l.Books); i++ {
		sum += l.Books[i].Score
	}
	// fmt.Printf("%#v: Sum %d\n", *l, sum)
	return sum
}

func (l *Library) Sort() {
	operationHappened := true
	for operationHappened == true {
		operationHappened = false
		for i := 0; i < len(l.Books)-1; i++ {
			if l.Books[i].Score < l.Books[i+1].Score {
				l.Books[i], l.Books[i+1] = l.Books[i+1], l.Books[i]
				operationHappened = true
			}
		}
	}
}

func (l *Library) SetBookAsUsed(id int) {
	for i := 0; i < len(l.Books); i++ {
		if l.Books[i].ID == id {
			l.Books[i] = l.Books[len(l.Books)-1]
			l.Books = l.Books[:len(l.Books)-1]
		}
	}
}
