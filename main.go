package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"prj-test/domain"
	"sort"
	"strconv"
	"time"
)

const (
	totalPoints       = 50
	pointsForRightAns = 50
)

var id uint64 = 1

func main() {
	fmt.Println("Вітаємо у грі Math-Monster!")

	var users []domain.User
	users = append(users, domain.User{Id: 1, Name: "Vasyl", Time: 5 * time.Second})
	users = append(users, domain.User{Id: 2, Name: "Mykola", Time: 3 * time.Second})
	users = append(users, domain.User{Id: 3, Name: "MAX", Time: 8 * time.Second})
	sortAndSave(users)
	// for {
	// 	menu()

	// 	choice := ""
	// 	fmt.Scan(&choice)

	// 	switch choice {
	// 	case "1":
	// 		user := play()
	// 		users = append(users, user)
	// 	case "2":
	// 		for _, u := range users {
	// 			fmt.Printf(
	// 				"id: %v, name: %s, time: %v\n",
	// 				u.Id, u.Name, u.Time,
	// 			)
	// 		}
	// 	case "3":
	// 		return
	// 	default:
	// 	}
	// }
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

func sortAndSave(users []domain.User) {
	sort.SliceStable(users, func(i, j int) bool {
		return users[i].Time < users[j].Time
	})

	file, err := os.OpenFile(
		"users.json",
		os.O_RDWR|os.O_CREATE|os.O_TRUNC,
		0775)
	if err != nil {
		log.Printf("os.OpenFile: %s", err)
		return
	}

	defer func() {
		err = file.Close()
		if err != nil {
			log.Printf("file.Close(): %s", err)
		}
	}()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(users)
	if err != nil {
		log.Printf("encoder.Encode(): %s", err)
	}
}
