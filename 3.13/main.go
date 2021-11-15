package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	fmt.Println("Roll the dice and try to guess the result in three attempts! Or press 'q' to quit.")
	fmt.Println("There are two dice to guess!")
	var input string
	var replay string
	var attemptFirst int = 1
	rand.Seed(time.Now().UnixNano())
	resultFirst := rand.Intn(6) + 1
	resultSecond := rand.Intn(6) + 1

	for {
		fmt.Print("Your guess: ")
		fmt.Scan(&input)

		if input == "q" {
			break
		}

		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Enter a valid number.")
			continue
		}

		if num == resultFirst {
			fmt.Println("Congratulations!! Correct!! The result is:", resultFirst)
			fmt.Println("Now guess the second.")

			var attemptSecond int = 1

			for {
				fmt.Print("Your guess: ")
				fmt.Scan(&input)

				if input == "q" {
					break
				}

				num, err := strconv.Atoi(input)
				if err != nil {
					fmt.Println("Enter a valid number.")
					continue
				}

				if num == resultSecond {
					fmt.Println("Congratulations!! Correct!! The result is:", resultSecond)
					fmt.Println(("Thanks for playing."))
					break
				}

				if attemptSecond == 2 {
					fmt.Println("Last chance!")
				}

				if attemptSecond == 3 {
					fmt.Println("Sorry! No more chances!")
					fmt.Print("Play again? (y/n)")
					fmt.Scan(&replay)
					if replay == "y" {
						main()
					}
					break
				}

				attemptSecond = attemptSecond + 1
			}

			fmt.Print("Play again? (y/n)")
			fmt.Scan(&replay)
			if replay == "y" {
				main()
			}
			break
		}

		if attemptFirst == 2 {
			fmt.Println("Last chance!")
		}

		if attemptFirst == 3 {
			fmt.Println("Sorry! No more chances!")
			fmt.Print("Play again? (y/n)")
			fmt.Scan(&replay)
			if replay == "y" {
				main()
			}
			break
		}

		attemptFirst = attemptFirst + 1

	}
}
