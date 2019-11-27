package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// prepare .tmp directory for caching
	if _, err := os.Stat(".tmp"); os.IsNotExist(err) {
		os.MkdirAll(".tmp", os.FileMode.Perm(0755))
	}

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
			h := GetRovHeroDetail(hero)

			heroByte, _ := json.Marshal(h)
			WriteToJson("rov", strings.ToLower(hero.Name), heroByte)
			time.Sleep(3 * time.Second)
		}

	case "aov":
		heroList := GetAovHeroList()
		for i := 0; i < len(heroList); i++ {
			hero := heroList[i]
			fmt.Println(hero.Name, "[", hero.Title, "]")
			h := GetAovHeroDetail(hero.ID)
			// fmt.Println(h.Title)
			// fmt.Printf("Skins #%d\n", len(h.Skins))
			// for j := 0; j < len(h.Skins); j++ {
			// 	fmt.Printf("- %s\n", h.Skins[j])
			// }
			// fmt.Printf("Skills #%d\n", len(h.Skills))
			// for j := 0; j < len(h.Skills); j++ {
			// 	r := h.Skills[j]
			// 	fmt.Printf("- %s %s\n  %s \n", r.SkillName, r.Cooldown, r.Desc)
			// }
			heroByte, _ := json.Marshal(h)
			WriteToJson("aov", strings.ToLower(hero.Name), heroByte)
			time.Sleep(3 * time.Second)
		}
	case "tw":
		heroList := GetTwHeroList()
		for i := 0; i < len(heroList); i++ {
			hero := heroList[i]
			fmt.Println(hero.Name, "[", hero.Role, "]")
			d := GetTwHeroDetail(hero)
			// fmt.Printf("Skills #%d\n", len(d.Skills))
			// for j := 0; j < len(d.Ratings); j++ {
			// 	r := d.Ratings[j]
			// 	fmt.Printf("- %s = %.1f\n", r.Name, r.Value)
			// }
			heroByte, _ := json.Marshal(d)
			WriteToJson("tw", strings.ToLower(hero.Name), heroByte)
			time.Sleep(3 * time.Second)
		}

	default:
		fmt.Println("DEFAULT")
	}

}
