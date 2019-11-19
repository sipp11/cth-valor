package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

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

func GetHeroList() []HeroLink {
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
