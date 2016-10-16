package main

import (
	"fmt"
	"math/rand"
	"time"
)


type playerMap map[string]int

func main(){
	var draftOrder [10]string
	rand.Seed(time.Now().UTC().UnixNano())
	var tallies [2]int
	for i := 0; i < 1000; i++ {
		draftees := []string{"Will", "Ron", "Greg", "Steve", "Derek", "Cameron", "Brandon", "Bryant", "JoshT", "JoshL"}

		m := map[string]int{
			"Will"                : 3,
			"Ron"                : 3,
			"Greg"                : 1,
			"Steve"        : 1,
			"Derek"        : 1,
			"Cameron"        : 1,
			"Brandon"        : 1,
			"Bryant"        : 1,
			"JoshT"        : 1,
			"JoshL"        : 1,
		}

		for i := 0; i < 10; i++ {
			// Display integer.
			player := pickPicker(draftees, m)
			fmt.Println(player)
			draftOrder[i] = player
		}

		if draftOrder[0] == "Will" || draftOrder[0] == "Ron"{
			tallies[0] = tallies[0] + 1
		}
		if draftOrder[1] == "Will" || draftOrder[1] == "Ron"{
			tallies[1] = tallies[1] + 1
		}

		fmt.Println("Showing draft order")
		for i := 0; i < 10; i++ {
			fmt.Print(draftOrder[i], " ")

		}
	}

	fmt.Println("tallies[0]: ", tallies[0], "tallies[1]: ", tallies[1])

}
func createCase(drafteesLeft []string, m map[string]int ) (int, []string){
	var drafteeList []string
	totalOddsLeft := 0
	for key, value := range m {
		drafteeList = createOddstoPick(drafteeList, key , value)
		totalOddsLeft = totalOddsLeft + value
		//Append(drafteeList, key)
	}

	return totalOddsLeft, drafteeList
}

func createOddstoPick(drafteeList []string, draftee string, odds int) []string{
	for i := 0; i < odds; i++ {
		drafteeList = append(drafteeList, draftee)
	}
	return drafteeList

}

func pickPicker( drafteesLeft []string, m map[string]int) (string){
	oddsLeft, drafteeList := createCase(drafteesLeft, m)
	fmt.Println("len(drafteeList) : ", len(drafteeList))
	fmt.Println("odds left : ", oddsLeft)
	randnumber := rand.Intn(len(drafteeList))
	pickedPerson := drafteeList[randnumber]
	delete(m, pickedPerson)
	return pickedPerson
}