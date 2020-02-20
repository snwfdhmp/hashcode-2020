package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func parseFile(path string) (Context, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return Context{}, err
	}

	// nbLib nbBooks maxDays
	// bookValue1 bookValue2 bookValue3
	// nbBookLib1 signupCostLib1 bookShipPerDayLib1
	// 2
	// 2
	// ...
	lines := strings.Split(string(b), "\n")

	line0Arr := strings.Split(lines[0], " ")
	nbLib, err := strconv.Atoi(line0Arr[1])
	if err != nil {
		return Context{}, err
	}

	nbBooks, err := strconv.Atoi(line0Arr[0])
	if err != nil {
		return Context{}, err
	}

	maxDays, err := strconv.Atoi(line0Arr[2])
	if err != nil {
		return Context{}, err
	}

	bookScores := strings.Split(lines[1], " ")
	fmt.Printf("BookScores: %#v\n", bookScores)
	for i := range bookScores {
		if bookScores[i] == "" {
			return Context{}, fmt.Errorf("i:%d contains ''", i)
		}
	}
	libraries := make([]Library, nbLib)

	for iLib := 0; iLib < len(lines[2:])/2; iLib++ {
		curLib := Library{}
		booksArr := strings.Split(lines[2+(2*iLib)+1], " ")
		libInfosArr := strings.Split(lines[2+(2*iLib)], " ")
		nbBooks, err = strconv.Atoi(libInfosArr[0])
		if err != nil {
			return Context{}, err
		}

		curLib.SignupTime, err = strconv.Atoi(libInfosArr[0])
		if err != nil {
			return Context{}, err
		}

		curLib.BooksPerDay, err = strconv.Atoi(libInfosArr[2])
		if err != nil {
			return Context{}, err
		}

		curLib.Books = make([]Book, nbBooks)
		for iBook := 0; iBook < nbBooks; iBook++ {
			id, err := strconv.Atoi(booksArr[iBook])
			if err != nil {
				fmt.Printf("id: %d\n", id)
				return Context{}, err
			}

			score, err := strconv.Atoi(bookScores[id])
			if err != nil {
				fmt.Printf("id: %d\n", id)
				return Context{}, err
			}

			curLib.Books[iBook] = Book{
				ID:    id,
				Score: score,
			}
		}
		libraries[iLib] = curLib
	}

	return Context{
		Libraries: libraries,
		DayMax:    maxDays,
	}, nil
}
