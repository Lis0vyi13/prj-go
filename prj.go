package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"prj-go/domain"
	"sort"
	"strconv"
	"time"
)

const (
	totalPoints       = 10
	pointsPerQuestion = 10
)

var id uint64 = 1

func main() {
	fmt.Println("Вітаємо у грі MATHREVENGE!")
	time.Sleep(1 * time.Second)

	users := getUsers()
	for _, user := range users {
		if (user.Id >= id) {
			id = user.Id + 1
		}
	}

	for {
		menu()

		choice := ""
		fmt.Scan(&choice)

		switch choice {
		case "1":
			user := play()
			users = getUsers()
			users = append(users, user)
			sortAndSave(users)
		case "2":
			users = getUsers()
			for i, user := range users {
				fmt.Printf(
					"i: %v, id: %v, name: %s, time: %v\n",
					i, user.Id, user.Name, user.Time,
				)
			}
		case "3":
			return
		default:
			fmt.Println("Зробіть коректний вибір!")
		}
	}

}

func menu() {
	fmt.Println("1. Почати гру")
	fmt.Println("2. Переглянути рейтинг")
	fmt.Println("3. Вийти")
}

func play() domain.User {
	for i := 3; i > 0; i-- {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}

	myPoints := 0
	start := time.Now()
	for myPoints < totalPoints {
		x, y := rand.Intn(100), rand.Intn(100)
		fmt.Printf("%v + %v = ", x, y)

		ans := ""
		fmt.Scan(&ans)

		ansInt, err := strconv.Atoi(ans)
		if err != nil {
			fmt.Println("Спробуй ще!")
		} else {
			if ansInt == x+y {
				myPoints += pointsPerQuestion
				points := totalPoints - myPoints
				fmt.Println("Правильно, ти набрав", myPoints, "очок!")
				fmt.Printf("Залишилось набрати %v!\n", points)
			} else {
				fmt.Println("Спробуй ще!")
			}
		}
	}
	end := time.Now()
	timeSpent := end.Sub(start)

	fmt.Printf("Молодчинка, впорався всього за %v!", timeSpent)
	fmt.Println("Введіть своє ім'я:")

	name := ""
	fmt.Scan(&name)

	// var user domain.User
	// user.Id = id
	// user.Name = name
	// user.Time = timeSpent

	user := domain.User{
		Id:   id,
		Name: name,
		Time: timeSpent,
	}
	id++

	return user
}

func sortAndSave(users []domain.User) {
	sort.SliceStable(users, func(i, j int) bool {
		return users[i].Time < users[j].Time
	})

	file, err := os.OpenFile(
		"users.json",
		os.O_RDWR|os.O_CREATE|os.O_TRUNC,
		0755)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	defer func(f *os.File) {
		err = f.Close()
		if err != nil {
			fmt.Printf("Error: %s", err)
		}
	}(file)

	encoder := json.NewEncoder(file)
	err = encoder.Encode(users)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
}

func getUsers() []domain.User {
	var users []domain.User

	info, err := os.Stat("users.json")
	if err != nil {
		if os.IsNotExist(err) {
			_, err := os.Create("users.json")
			if err != nil {
				fmt.Printf("Error: %s", err)
				return nil
			}
			return nil
		}
		fmt.Printf("Error: %s", err)
		return nil
	}

	if info.Size() != 0 {
		file, err := os.Open("users.json")
		if err != nil {
			fmt.Printf("Error: %s", err)
			return nil
		}

		defer func(f *os.File) {
			err = f.Close()
			if err != nil {
				fmt.Printf("Error: %s", err)
			}
		}(file)

		decoder := json.NewDecoder(file)
		err = decoder.Decode(&users)
		if err != nil {
			fmt.Printf("Error: %s", err)
			return nil
		}
	}

	return users
}