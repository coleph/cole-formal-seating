package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// variables to track table numbers assigned and which indexes are being modified
var TableNumber = 1
var Place1 = 0
var Place2 = 9
var TableNumberString = ""

// data stored in each index of the slice
type Person struct {
	Firstname string
	Lastname  string
	Table     string
	Table2nd  string
	Table3rd  string
}

// re arrange the slice
func shuffle(peopleList []Person) []Person {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(peopleList), func(i, j int) { peopleList[i], peopleList[j] = peopleList[j], peopleList[i] })
	return peopleList
}

func firstSeating(peopleList []Person) {
	// randimization credit to youbasic.org go tutorial
	peopleList = shuffle(peopleList)
	// assign 8 people at a time a table number 1-31, int is converted to a string for Table
	for count := 0; count <= 30; count++ {

		for index := Place1; index < Place2; index++ {
			TableNumberString := strconv.FormatInt(int64(TableNumber), 10)

			peopleList[index].Table = TableNumberString
		}
		// change the position of people[] modified
		TableNumber = TableNumber + 1
		Place1 = Place1 + 8
		Place2 = Place2 + 8

	}
	for index := 248; index < 279; index++ {
		peopleList[index].Table = "Waiter"
	}
	for index := 279; index < 290; index++ {
		peopleList[index].Table = "KC"
	}
	//fmt.Println(peopleList)
	secondSeating(peopleList)

}

// copy of slice is sent to next function for new assignments
func secondSeating(peopleList []Person) {

	// seperates all students assigned a first table
	var index = 0
	for count := 0; count <= 31; count++ {
		for rep := 0; rep <= 8; rep++ {

			peopleList[index].Table2nd = strconv.FormatInt(int64(count+rep), 10)
			index++

		}

	}
	// put the kitchen crew and waiter students somewhere
	// if the table numbers assigned go over 31, the students will be re shuffled into a new part seating
	var sat = 1
	for index = 248; index <= 289; index++ {

		if peopleList[index].Table == "Waiter" || peopleList[index].Table2nd == "KC" {
			sat++
			peopleList[index].Table = strconv.FormatInt(int64(sat), 10)
			if sat == 31 {
				sat = 0
			}
		}
	}
	index = 0
	for index = 0; index <= 289; index++ {
		if peopleList[index].Table2nd == "32" {
			peopleList[index].Table2nd = "1"
		}
		if peopleList[index].Table2nd == "33" {
			peopleList[index].Table2nd = "2"
		}
		if peopleList[index].Table2nd == "34" {
			peopleList[index].Table2nd = "Waiter"
		}
		if peopleList[index].Table2nd == "35" {
			peopleList[index].Table2nd = "Waiter"
		}
		if peopleList[index].Table2nd == "36" {
			peopleList[index].Table2nd = "Waiter"
		}
		if peopleList[index].Table2nd == "37" {
			peopleList[index].Table2nd = "Waiter"
		}
		if peopleList[index].Table2nd == "38" {
			peopleList[index].Table2nd = "KC"
		}
		if peopleList[index].Table2nd == "39" {
			peopleList[index].Table2nd = "waiter"
		}

	}

	thirdSeating(peopleList)
}

func thirdSeating(peopleList []Person) {
	// same logic as secondSeating but it probably doesn't actually work

	var index = 0

	for count := 0; count <= 31; count++ {
		for rep := 0; rep <= 8; rep++ {

			peopleList[index].Table3rd = strconv.FormatInt(int64(count+rep+1), 10)
			index++

		}
	}
	var sat = 1
	for index = 248; index <= 289; index++ {

		if peopleList[index].Table2nd == "Waiter" || peopleList[index].Table2nd == "KC" {
			sat++
			peopleList[index].Table3rd = strconv.FormatInt(int64(sat), 10)
			if sat == 31 {
				sat = 0
			}
		}

	}

	index = 0
	for index = 0; index <= 289; index++ {
		if peopleList[index].Table3rd == "32" {
			peopleList[index].Table3rd = "1"
		}
		if peopleList[index].Table3rd == "33" {
			peopleList[index].Table3rd = "2"
		}
		if peopleList[index].Table3rd == "34" {
			peopleList[index].Table3rd = "Wa"
		}
		if peopleList[index].Table3rd == "35" {
			peopleList[index].Table3rd = "Waiter"
		}
		if peopleList[index].Table3rd == "36" {
			peopleList[index].Table3rd = "Waiter"
		}
		if peopleList[index].Table3rd == "37" {
			peopleList[index].Table3rd = "Waiter"
		}
		if peopleList[index].Table3rd == "38" {
			peopleList[index].Table3rd = "KC"
		}
		if peopleList[index].Table3rd == "39" {
			peopleList[index].Table3rd = "waiter"
		}

	}
	fmt.Println(peopleList)
}

func main() {
	// aaa

	csvFile, _ := os.Open("list.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var people []Person
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		people = append(people, Person{
			Firstname: line[0],
			Lastname:  line[1],
		})
	}

	firstSeating(people)

}
