package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
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

	case "media":
		// example command : go run *.go media tw 2019-12-11
		today := time.Now()
		day := today.Format("2006-01-02")
		server := args[1]
		if len(args) == 3 {
			day = args[2]
		}
		root := filepath.Join("./data", server, day)
		files := GetJsonIn(root)

		switch server {
		case "tw":
			var hero TwHeroDetail
			for _, file := range files {
				jsonData := ReadFile(file)
				err := json.Unmarshal(jsonData, &hero)
				CheckError(err)
				MediaDownload(server, day, "hero", hero.Hero.Image)
				for i := 0; i < len(hero.Skills); i++ {
					fetched := MediaDownload(server, day, "skill", hero.Skills[i].Image)
					if fetched {
						time.Sleep(1 * time.Second)
					}
				}
			}

		case "aov":
			var hero AovHero
			baseURL := "https://www.arenaofvalor.com/images/heroes/"
			var url string
			var fetched bool
			for _, file := range files {
				jsonData := ReadFile(file)
				err := json.Unmarshal(jsonData, &hero)
				CheckError(err)

				// each skin has 2 images: icon & large
				for _, skin := range hero.Skins {
					url = fmt.Sprintf("%s%s/%s_%s.jpg", baseURL, "skin", skin, "big")
					fetched = MediaDownload(server, day, "skin", url)
					if fetched {
						time.Sleep(1 * time.Second)
					}
					url = fmt.Sprintf("%s%s/%s_%s.jpg", baseURL, "skin", skin, "icon")
					fetched = MediaDownload(server, day, "skin", url)
					if fetched {
						time.Sleep(1 * time.Second)
					}
				}
				// spell 1
				// https://www.arenaofvalor.com/images/heroes/skill/80104.png
				url = fmt.Sprintf("%s%s/%s.png", baseURL, "skill", hero.RecommdSkill1)
				fetched = MediaDownload(server, day, "spell", url)
				if fetched {
					time.Sleep(1 * time.Second)
				}
				// spell 2
				// https://www.arenaofvalor.com/images/heroes/skill/80104.png
				url = fmt.Sprintf("%s%s/%s.png", baseURL, "skill", hero.RecommdSkill2)
				fetched = MediaDownload(server, day, "spell", url)
				if fetched {
					time.Sleep(1 * time.Second)
				}

				// skill
				// https://www.arenaofvalor.com/images/heroes/skill/80104.png
				for ind, skill := range hero.Skills {
					if ind > 3 {
						continue
					}
					url = fmt.Sprintf("%s%s/%s.png", baseURL, "skill", skill.SkillIcon)
					fetched = MediaDownload(server, day, "skill", url)
					if fetched {
						time.Sleep(1 * time.Second)
					}
				}
			}

		case "rov":
			var hero Hero
			var fetched bool
			for _, file := range files {
				jsonData := ReadFile(file)
				err := json.Unmarshal(jsonData, &hero)
				CheckError(err)

				// hero
				fetched = MediaDownload(server, day, "hero", hero.Link.Image)
				if fetched {
					time.Sleep(1 * time.Second)
				}
				// skill
				for _, skill := range hero.Skills {
					fetched = MediaDownload(server, day, "skill", skill.Image)
					if fetched {
						time.Sleep(1 * time.Second)
					}
				}
				// skin
				for _, skin := range hero.Skins {
					fetched = MediaDownload(server, day, "skin", skin.Image)
					if fetched {
						time.Sleep(1 * time.Second)
					}
					fetched = MediaDownload(server, day, "skin", skin.Banner)
					if fetched {
						time.Sleep(1 * time.Second)
					}
				}
				// spell
				fetched = MediaDownload(server, day, "spell", hero.Spell.Image)
				if fetched {
					time.Sleep(1 * time.Second)
				}
				// arcana
				for _, arcana := range hero.Arcanas {
					for _, subarcana := range arcana {
						if len(subarcana.Image) > 0 {
							fetched = MediaDownload(server, day, "arcana", subarcana.Image)
							if fetched {
								time.Sleep(1 * time.Second)
							}
						}
					}
				}
				// enchantment

				for _, eSet := range hero.EnchantmentSets {
					for _, enchant := range eSet {
						fetched = MediaDownload(server, day, "enchantment", enchant.Image)
						if fetched {
							time.Sleep(1 * time.Second)
						}
					}
				}

				// item
				for _, build := range hero.BuildSets {
					for _, item := range build.Items {
						fetched = MediaDownload(server, day, "item", item.Image)
						if fetched {
							time.Sleep(1 * time.Second)
						}
					}
				}
			}

		default:
			fmt.Println("DEFAULT")
		}
	}
}
