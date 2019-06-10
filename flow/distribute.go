package main

import "fmt"

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie",
		"Peter", "Gianna", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

// "A"
func distribute(name string, coins int) int {
	nlen := len(name)
	value := 0
	for i := 0; i < nlen; i++ {
		switch string(name[i]) {
		case "a", "A":
			value += 1
		case "e", "E":
			value += 1
		case "i", "I":
			value += 2
		case "o", "O":
			value += 3
		case "u", "U":
			value += 4
		}
	}
	if value > 10 {
		value = 10
	}
	return value
}

func main() {
	fmt.Println(distribution)
	fmt.Println("coins left %d", coins)

	for _, name := range users {
		v := distribute(name, coins)
		distribution[name] = v
		coins = coins - v
	}

	// distribute (distribution, coins)

	fmt.Printf("%#v", distribution)
	fmt.Printf("coins left %d\n", coins)
}
