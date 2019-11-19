package main

import "fmt"

func main() {
	heroList := GetHeroList()

	for i := 0; i < len(heroList); i++ {
		hero := heroList[i]
		fmt.Println(hero.Name, "[", hero.RolePrimary, "]")
	}
}
