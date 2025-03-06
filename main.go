package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const (
	totalPoints       = 50
	pointsForRightAns = 10
)

func main() {
	fmt.Println("Вітаємо у грі Math-Monster!")

	for i := 5; i > 0; i-- {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}

	myPoints := 0
	for myPoints < totalPoints {
		x, y := rand.Intn(100), rand.Intn(100)

		// fmt.Println(x, " + ", y, " = ")
		fmt.Printf("%v + %v = ", x, y)

		var ans string
		fmt.Scan(&ans)

		sum := x + y

		ansInt, err := strconv.Atoi(ans)
		if err != nil {
			fmt.Println("Не правильно!")
		} else {
			if sum == ansInt {
				myPoints += pointsForRightAns
				fmt.Printf("Вітаю! У тебе %v очок!\n", myPoints)
			} else {
				fmt.Println("Не правильно!")
			}
		}
	}
}
