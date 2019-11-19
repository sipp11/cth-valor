package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/anaskhan96/soup"
)

type Skill struct {
	Name  string `json:"name"`
	Desc  string `json:"description"`
	Image string `json:"image"`
	IsEx  string `json:"isEx"`
}

type Skin struct {
	Name   string `json:"name"`
	Image  string `json:"image"`
	Banner string `json:"image_banner"`
}

type Spell struct {
	Name   string `json:"name"`
	Detail string `json:"detail"`
	Image  string `json:"image"`
}

type Arcana struct {
	Name   string `json:"name"`
	Detail string `json:"detail"`
	Image  string `json:"image"`
}

type Enchantment struct {
	Type   string `json:"type"`
	Name   string `json:"name"`
	Detail string `json:"detail"`
	Image  string `json:"image"`
}

type Build struct {
	Type   string `json:"type"`
	Name   string `json:"name"`
	Detail string `json:"detail"`
	Image  string `json:"image"`
}

type BuildSet struct {
	Name  string  `json:"name"`
	Items []Build `json:"items"`
}

type Hero struct {
	Name            string                     `json:"name"`
	Story           string                     `json:"story"`
	Skills          []Skill                    `json:"skills"`
	Skins           []Skin                     `json:"skins"`
	Spell           Spell                      `json:"spell"`
	Arcanas         []([]Arcana)               `json:"runes"`
	EnchantmentSets map[string]([]Enchantment) `json:"enchantments"`
	BuildSets       []BuildSet                 `json:"itemsets"`
}

type Content struct {
	Hero Hero `json:"data"`
}

type PageProp struct {
	Content Content `json:"content"`
}

type InitialProp struct {
	PageProps PageProp `json:"pageProps"`
}

type Prop struct {
	InitialProps InitialProp `json:"initialProps"`
}

type Data struct {
	Props Prop `json:"props"`
}

type HeroLink struct {
	Name          string `json:"name"`
	Slug          string `json:"slug"`
	Image         string `json:"image"`
	RolePrimary   string `json:"role_text"`
	RoleSecondary string `json:"role_second_text"`
}

type Heroes struct {
	HeroList []HeroLink `json:"heroes"`
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    Heroes `json:"data"`
}

func GetRovHeroList() []HeroLink {
	var heroes []HeroLink
	var resp Response
	var byteValue []byte

	if _, err := os.Stat("hero_list.json"); err == nil {
		jsonFile, err := os.Open("hero_list.json")
		if err != nil {
			fmt.Println(err)
		}
		byteValue, _ = ioutil.ReadAll(jsonFile)
		fmt.Println("From File")
		defer jsonFile.Close()
	} else {
		httpResp, err := http.Get("https://rov.in.th/api/v2/getHeroList")
		if err != nil {
			fmt.Println(err)
		}
		byteValue, _ = ioutil.ReadAll(httpResp.Body)
		// write to file too
		f, _ := os.Create("hero_list.json")
		f.Write(byteValue)
		defer f.Close()
		fmt.Println("From REAL Server")
	}
	json.Unmarshal(byteValue, &resp)
	heroes = resp.Data.HeroList
	return heroes
}

func GetRovHeroDetail(slug string) {
	fp := fmt.Sprintf("/tmp/aov/%s.html", slug)
	var doc soup.Root
	if _, err := os.Stat(fp); err == nil {
		htmlFile, _ := ioutil.ReadFile(fp)
		doc = soup.HTMLParse(string(htmlFile))
	} else {
		url := fmt.Sprintf("https://rov.in.th/hero/%s", slug)
		resp, err := soup.Get(url)
		if err != nil {
			os.Exit(1)
		}
		doc = soup.HTMLParse(resp)
	}
	data := doc.Find("script", "id", "__NEXT_DATA__")
	// fmt.Println("data", data.Text())

	// var result map[string]interface{}
	// json.Unmarshal([]byte(data.Text()), &result)
	// props := result["props"]

	// ["initialProps"]["pageProps"]
	// fmt.Println(props)

	var jsonData Data
	// fmt.Println(data.Text())
	json.Unmarshal([]byte(data.Text()), &jsonData)
	hero := jsonData.Props.InitialProps.PageProps.Content.Hero

	fmt.Println("----------------------------------")
	fmt.Println("## Skill")
	// fmt.Println(hero.Skills)
	for i := 0; i < len(hero.Skills); i++ {
		a := hero.Skills[i]
		fmt.Println("-", a.Name)
	}
	fmt.Println("----------------------------------")
	fmt.Println("## Skin")
	for i := 0; i < len(hero.Skins); i++ {
		a := hero.Skins[i]
		fmt.Println("-", a.Name)
	}
	fmt.Println("----------------------------------")
	fmt.Println("## Story")
	fmt.Printf("   %s\n", hero.Story)
	fmt.Println("----------------------------------")
	fmt.Println("## Spell")
	fmt.Printf("   %s\n", hero.Spell.Name)
	fmt.Println("----------------------------------")
	fmt.Printf("## Arcana\n   ")
	for _, v := range hero.Arcanas {
		for i := 0; i < len(v); i++ {
			a := v[i]
			if a.Name != "" {
				fmt.Printf("%s, ", a.Name)
			}
		}
	}
	fmt.Println()
	fmt.Println("----------------------------------")
	fmt.Print("## Enchancements")
	for k, v := range hero.EnchantmentSets {
		fmt.Printf("\n * %s: ", k)
		for i := 0; i < len(v); i++ {
			a := v[i]
			fmt.Printf("%s, ", a.Name)
		}
	}
	fmt.Println()
	fmt.Println("----------------------------------")
	fmt.Print("## Build")
	for _, v := range hero.BuildSets {
		fmt.Print("\n- ")
		for i := 0; i < len(v.Items); i++ {
			a := v.Items[i]
			fmt.Printf("%s, ", a.Name)
		}
	}
	fmt.Println()
}
