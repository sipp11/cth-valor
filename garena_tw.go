package main

import (
	"os"
	"strconv"
	"strings"

	"github.com/anaskhan96/soup"
)

var twHeroListURL = "https://pro.moba.garena.tw/heroList"
var twHeroURL = "https://moba.garena.tw/game/hero/"

type TwHeroList struct {
	Name  string
	Role  string
	Url   string
	Image string
}

type TwSkill struct {
	Order int
	Image string
	Text  string
}

type TwRating struct {
	Name  string
	Value float64
}

type TwHeroDetail struct {
	Skills  []TwSkill
	Ratings []TwRating
}

func GetTwHeroList() []TwHeroList {
	var heroes []TwHeroList
	resp, err := soup.Get(twHeroListURL)
	if err != nil {
		os.Exit(1)
	}
	doc := soup.HTMLParse(resp)
	heroList := doc.FindAll("div", "class", "herolist-list__item")
	for i := 0; i < len(heroList); i++ {
		one := heroList[i]
		role := one.Attrs()["data-type"]
		link := one.Find("a").Attrs()["href"]
		img := one.Find("img").Attrs()["src"]
		name := one.Find("div", "class", "herolist-list__item-name").Text()
		name = strings.TrimSpace(name)
		hero := TwHeroList{Name: name, Url: link, Image: img, Role: role}
		heroes = append(heroes, hero)
	}
	return heroes
}

func GetTwHeroDetail(url string) TwHeroDetail {
	// url := fmt.Sprintf("%s%s", twHeroURL, id)
	resp, err := soup.Get(url)
	if err != nil {
		os.Exit(1)
	}
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
	return TwHeroDetail{Skills: skills, Ratings: points}
}