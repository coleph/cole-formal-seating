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

var TableNumber = 1
var Place1 = 0
var Place2 = 9

type Person struct {
	Firstname string
	Lastname  string
	Table     string
	Table2nd  string
}

var Waiters []Person

func main() {
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

	// randimization credit to youbasic.org go tutorial
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(people), func(i, j int) { people[i], people[j] = people[j], people[i] })
	// assign 8 people at a time a table number 1-31, int is converted to a string for Table
	for count := 0; count <= 30; count++ {

		for index := Place1; index < Place2; index++ {
			TableNumberString := strconv.FormatInt(int64(TableNumber), 10)
			people[index].Table = TableNumberString
		}
		// change the position of people[] modified
		TableNumber = TableNumber + 1
		Place1 = Place1 + 8
		Place2 = Place2 + 8

	}
	for index := 248; index < 279; index++ {
		people[index].Table = "Waiter"
	}
	for index := 279; index < 290; index++ {
		people[index].Table = "KC"
	}

	fmt.Println(people)
	// shuffled and everyone is assigned a table or waiter/kc

}
