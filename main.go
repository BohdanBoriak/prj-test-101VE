package main

import (
	"fmt"
	"math/rand"
	"prj-test/domain"
	"strconv"
	"time"
)

const (
	totalPoints       = 50
	pointsForRightAns = 10
)

var id uint64 = 1

func main() {
	fmt.Println("Вітаємо у грі Math-Monster!")

	var users []domain.User

	for {
		menu()

		choice := ""
		fmt.Scan(&choice)

		switch choice {
		case "1":
			user := play()
			users = append(users, user)
		case "2":
			fmt.Println("Рейтинг в розробці -_-")
		case "3":
			return
		default:
		}
	}
}

func menu() {
	fmt.Println("1. Грати!")
	fmt.Println("2. Рейтинг -_-")
	fmt.Println("3. Вийти :(")
}

func play() domain.User {
	startTime := time.Now()

	myPoints := 0
	for myPoints < totalPoints {
		x, y := rand.Intn(100), rand.Intn(100)

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
	endTime := time.Now()
	duration := endTime.Sub(startTime)

	fmt.Println("Введіть ім'я:")

	name := ""
	fmt.Scan(&name)

	var u domain.User
	u.Id = id
	id++
	u.Name = name
	u.Time = duration

	return u
}
