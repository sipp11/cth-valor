package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type AovHeroLink struct {
	Name  string `json:"name"`
	Title string `json:"title"`
	ID    string `json:"heroid"`
}

type AovHeroesResponse struct {
	ErrCode    string        `json:"err_code"`
	ErrMessage string        `json:"err_msg"`
	Heroes     []AovHeroLink `json:"data"`
}

type AovSkill struct {
	SkillID        string `json:"skillid"`
	SkillName      string `json:"skillname"`
	Desc           string `json:"desc"`
	SkillIcon      string `json:"skillicon"`
	Cooldown       string `json:"cooldown"`
	CooldownGrow   string `json:"cooldowngrow"`
	EnergyCostType string `json:"energycosttype"`
	EnergyCost     string `json:"energycost"`
	EnergyCostGrow string `json:"energycostgrow"`
}

type AovHero struct {
	Name          string     `json:"name"`
	Title         string     `json:"title"`
	Icon          string     `json:"icon"`
	Tips          string     `json:"tips"`
	Story         string     `json:"story"`
	Viability     string     `json:"viability"`
	Damage        string     `json:"damage"`
	SpellDamage   string     `json:"spelldamage"`
	Difficulty    string     `json:"difficulty"`
	Basehp        string     `json:"basehp"`
	GrowthHp      string     `json:"growthhp"`
	Baseatk       string     `json:"baseatk"`
	GrowthAtk     string     `json:"growthatk"`
	Basedef       string     `json:"basedef"`
	GrowthDef     string     `json:"growthdef"`
	BaseRes       string     `json:"baseres"`
	GrowthRes     string     `json:"growthres"`
	Skins         []string   `json:"skin"`
	Skills        []AovSkill `json:"skill"`
	RecommdSkill1 string     `json:"recommdskill1"`
	RecommdSkill2 string     `json:"recommdskill2"`
}

// Skin URL
// https://www.arenaofvalor.com/images/heroes/skin/17700_big.jpg
// https://www.arenaofvalor.com/images/heroes/skin/17701_icon.jpg

// skill URL
// https://www.arenaofvalor.com/images/heroes/skill/17710.png

type AovHeroResponse struct {
	ErrCode    string  `json:"err_code"`
	ErrMessage string  `json:"err_msg"`
	Hero       AovHero `json:"data"`
}

// Page
var baseURL = "https://www.arenaofvalor.com"
var heroListURL = fmt.Sprintf("%s/web2017/herolist.html", baseURL)
var aovHeroDetail = "https://www.arenaofvalor.com/web2017/heroDetails.html?id=%s"

// API
var aovHeroURL = "https://mws.eutc.ngame.proximabeta.com/fcgi-bin/gift.fcgi?heroid=0&ticket=miniweb&"
var aovHeroDetailApi = "https://mws.eutc.ngame.proximabeta.com/fcgi-bin/gift.fcgi?heroid=%s&ticket=miniweb&_=%s"

func GetAovHeroList() []AovHeroLink {

	var heroes []AovHeroLink
	var resp AovHeroesResponse
	var byteValue []byte

	if _, err := os.Stat("aov_hero_list.json"); err == nil {
		jsonFile, err := os.Open("aov_hero_list.json")
		if err != nil {
			fmt.Println(err)
		}
		byteValue, _ = ioutil.ReadAll(jsonFile)
		fmt.Println("From File")
		defer jsonFile.Close()
	} else {
		httpResp, err := http.Get(aovHeroURL)
		if err != nil {
			fmt.Println(err)
		}
		byteValue, _ = ioutil.ReadAll(httpResp.Body)
		// write to file too
		f, _ := os.Create("aov_hero_list.json")
		f.Write(byteValue)
		defer f.Close()
		fmt.Println("From REAL Server")
	}
	json.Unmarshal(byteValue, &resp)
	heroes = resp.Heroes
	return heroes
}

func GetAovHeroDetail(id string) AovHero {
	url := fmt.Sprintf(aovHeroDetailApi, id, time.Now().Unix())
	var hero AovHero
	var resp AovHeroResponse
	var byteValue []byte

	fp := fmt.Sprintf("aov_hero_%s.json", id)

	if _, err := os.Stat(fp); err == nil {
		jsonFile, err := os.Open(fp)
		if err != nil {
			fmt.Println(err)
		}
		byteValue, _ = ioutil.ReadAll(jsonFile)
		fmt.Println("From File")
		defer jsonFile.Close()
	} else {
		httpResp, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
		}
		byteValue, _ = ioutil.ReadAll(httpResp.Body)
		// write to file too
		f, _ := os.Create(fp)
		f.Write(byteValue)
		defer f.Close()
		fmt.Println("From REAL Server")
	}
	json.Unmarshal(byteValue, &resp)
	hero = resp.Hero
	return hero
}
