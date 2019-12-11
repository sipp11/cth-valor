package main

import (
	"strconv"
	"strings"

	"github.com/anaskhan96/soup"
)

var twHeroListURL = "https://pro.moba.garena.tw/heroList"
var twHeroURL = "https://moba.garena.tw/game/hero/"

type TwHero struct {
	Name  string `json:"name"`
	Role  string `json:"role"`
	URL   string `json:"url"`
	Image string `json:"image"`
}

type TwSkill struct {
	Order int    `json:"order"`
	Image string `json:"image"`
	Text  string `json:"text"`
}

type TwRating struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

type TwHeroDetail struct {
	Hero    TwHero     `json:"hero"`
	Skills  []TwSkill  `json:"skills"`
	Ratings []TwRating `json:"ratings"`
}

func GetTwHeroList() []TwHero {
	var heroes []TwHero
	resp, err := soup.Get(twHeroListURL)
	CheckError(err)
	doc := soup.HTMLParse(resp)
	heroList := doc.FindAll("div", "class", "herolist-list__item")
	for i := 0; i < len(heroList); i++ {
		one := heroList[i]
		role := one.Attrs()["data-type"]
		link := one.Find("a").Attrs()["href"]
		img := one.Find("img").Attrs()["src"]
		name := one.Find("div", "class", "herolist-list__item-name").Text()
		name = strings.TrimSpace(name)
		hero := TwHero{Name: name, URL: link, Image: img, Role: role}
		heroes = append(heroes, hero)
	}
	return heroes
}

func GetTwHeroDetail(hero TwHero) TwHeroDetail {
	// url := fmt.Sprintf("%s%s", twHeroURL, id)
	resp, err := soup.Get(hero.URL)
	CheckError(err)
	doc := soup.HTMLParse(resp)
	data := doc.Find("div", "class", "hero-data")
	skillList := data.Find("div", "class", "hero-data__skills-main").FindAll("div")

	var skills []TwSkill
	for i := 0; i < len(skillList); i++ {
		skill := skillList[i]
		detail := skill.Attrs()["title"]
		img := skill.Find("img").Attrs()["src"]
		o := TwSkill{Order: i, Image: img, Text: detail}
		skills = append(skills, o)
	}

	pointList := data.FindAll("dl", "class", "hero-data__overview-item")
	var points []TwRating
	for i := 0; i < len(pointList); i++ {
		p := pointList[i]
		name := p.Find("dt").Text()
		ptt := p.Find("dd", "class", "hero-data__overview-item-val").Text()
		if pt, err := strconv.ParseFloat(ptt, 32); err == nil {
			aa := TwRating{Name: name, Value: pt}
			points = append(points, aa)
		}

	}
	return TwHeroDetail{Hero: hero, Skills: skills, Ratings: points}
}
