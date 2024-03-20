package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
    fmt.Println("Вітаємо у грі MATHREVENGE!")

    countdown := 5

    fmt.Println("Гра розпочнеться через кілька секунд...")

    for i := countdown; i > 0; i-- {
        fmt.Print(i,"...\n")
        time.Sleep(time.Second)
    }
	start := time.Now();

	fmt.Println("Гру розпочато!")
	myPoints := 0
	const (totalPoints = 50 
		   pointsPerQuestion = 5)
	for myPoints < totalPoints {
		x,y := rand.Intn(100),rand.Intn(100)
		fmt.Printf("%v + %v = ", x, y)
		ans := ""
		fmt.Scan(&ans)
		ansInt, err := strconv.Atoi(ans)
		if err != nil {
			fmt.Println("Спробуй ще!")
		} else {
			if (ansInt == x + y) {
				myPoints += pointsPerQuestion
				points := totalPoints - myPoints
				fmt.Println("Правильно, ти набрав", myPoints,"очок")
				fmt.Printf("Залишилось набрати %v!\n", points)
			} else {
				fmt.Println("Спробуй ще")
			}
		}
	}		
	end := time.Now()
	timeSpend := end.Sub(start)

	fmt.Printf("Молодчинка, впорався всього за %v", timeSpend)
	time.Sleep(5 * time.Second)

}