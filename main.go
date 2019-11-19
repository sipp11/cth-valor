package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Help: command option required, e.g. rov, aov, tw")
		os.Exit(0)
	}

	switch args[0] {
	case "rov":
		heroList := GetRovHeroList()
		for i := 0; i < len(heroList); i++ {
			hero := heroList[i]
			fmt.Println(hero.Name, "[", hero.RolePrimary, "]")
			GetRovHeroDetail(hero.Name)
			break
		}

	case "aov":
		heroList := GetAovHeroList()
		for i := 0; i < len(heroList); i++ {
			hero := heroList[i]
			fmt.Println(hero.Name, "[", hero.Title, "]")
			h := GetAovHeroDetail(hero.ID)
			fmt.Println(h.Title)
			fmt.Printf("Skins #%d\n", len(h.Skins))
			for j := 0; j < len(h.Skins); j++ {
				fmt.Printf("- %s\n", h.Skins[j])
			}
			fmt.Printf("Skills #%d\n", len(h.Skills))
			for j := 0; j < len(h.Skills); j++ {
				r := h.Skills[j]
				fmt.Printf("- %s %s\n  %s \n", r.SkillName, r.Cooldown, r.Desc)
			}
			break
		}
	case "tw":
		heroList := GetTwHeroList()
		for i := 0; i < len(heroList); i++ {
			hero := heroList[i]
			fmt.Println(hero.Name, "[", hero.Role, "]")
			d := GetTwHeroDetail(hero.Url)
			fmt.Printf("Skills #%d\n", len(d.Skills))
			for j := 0; j < len(d.Ratings); j++ {
				r := d.Ratings[j]
				fmt.Printf("- %s = %.1f\n", r.Name, r.Value)
			}
			break
		}

	default:
		fmt.Println("DEFAULT")
	}

}
