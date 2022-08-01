package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	charity_finder("animals")
}

func charity_finder(c string) {
	var default_sports Return_Charity
	var default_nature Return_Charity
	var default_food Return_Charity
	var default_animals Return_Charity

	default_sports.Mission = "Good Sports drives equitable access in youth sports and physical activity, by supporting children in high-need communities to achieve their greatest potential, on the field and in life."
	default_sports.Street1 = "1515 Washington Street, Suite 300"
	default_sports.Street2 = ""
	default_sports.City = "Braintree"
	default_sports.Region = "MA"
	default_sports.Postal_code = "02184"
	default_sports.Website_url = "https://www.goodsports.org/"

	default_nature.Name = "The Nature Conservancy"
	default_nature.Mission = "Guided by science, TNC creates innovative, on-the-ground solutions to our world's toughest challenges so that people and nature can thrive together. Join us as we impact conservation in all 50 states and in more than 70 countries and territories around the globe!"
	default_nature.Street1 = "4245 North Fairfax Drive, Suite 100"
	default_nature.Street2 = ""
	default_nature.City = "Arlington"
	default_nature.Region = "VA"
	default_nature.Postal_code = "22203-1606"
	default_nature.Website_url = "https://www.nature.org/en-us/"

	default_food.Name = "Feeding America"
	default_food.Mission = "Feeding America is a nonprofit network of 200 food banks leading the fight against hunger in the United States."
	default_food.Street1 = "161 North Clark Street, Suite 700"
	default_food.Street2 = ""
	default_food.City = "Chicago"
	default_food.Region = "IL"
	default_food.Postal_code = "60601"
	default_food.Website_url = "https://www.feedingamerica.org/"

	default_animals.Name = "World Wildlife Fund"
	default_animals.Mission = "As the world's leading conservation organization, World Wildlife Fund works in nearly 100 countries to tackle the most pressing issues at the intersection of nature, people, and climate. We collaborate with local communities to conserve the natural resouselected_categoryes we all depend on and build a future in which people and nature thrive. Together with partners at all levels, we transform markets and policies toward sustainability, tackle the threats driving the climate crisis, and protect and restore wildlife and their habitats."
	default_animals.Street1 = "1250 24th Street, N.W."
	default_animals.Street2 = ""
	default_animals.City = "Washington"
	default_animals.Region = "DC"
	default_animals.Postal_code = "20037-1193"
	default_animals.Website_url = "https://www.worldwildlife.org/"

	category := make(map[string]string)
	category["sports"] = "athlete"
	category["nature"] = "environment"
	category["animals"] = "dog"
	category["food"] = "food"

	selected_category := category[c]

	req, err := http.NewRequest("GET", "https://api.pledge.to/v1/organizations", nil)
	if err != nil {
		fmt.Print("error")
		os.Exit(1)
	}
	req.Header.Set("Authorization", "Bearer e34a77e141e3821b2911707475653d04")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var response Response
	var selected_charity Return_Charity

	json.Unmarshal(body, &response)
	change := false
	for i, c := range response.Charities {
		if strings.Contains(c.Name, selected_category) || strings.Contains(c.Mission, selected_category) {
			selected_charity.Name = c.Name
			selected_charity.Mission = c.Mission
			selected_charity.Street1 = c.Street1
			selected_charity.Street2 = c.Street2
			selected_charity.City = c.City
			selected_charity.Region = c.Region
			selected_charity.Postal_code = c.Postal_code
			selected_charity.Website_url = c.Website_url
			change = true
			break
		}
		i++
	}
	if !change {
		switch c {
		case "sports":
			selected_charity = default_sports
		case "nature":
			selected_charity = default_nature
		case "animals":
			selected_charity = default_animals
		case "food":
			selected_charity = default_food
		}
	}
	fmt.Println(selected_charity)
}

type Response struct {
	Page        int
	Per         int
	Uri         string
	Next        string
	Previous    string
	Total_count int
	Charities   []Charity `json:"results"`
}

type Charity struct {
	Id                string
	Name              string
	Alias             string
	Ngo_id            string
	Mission           string
	Street1           string
	Street2           string
	City              string
	Region            string
	Postal_code       string
	Country           string
	Lat               string
	Lon               string
	Website_url       string
	Profile_url       string
	Logo_url          string
	Disbursement_type string
}

type Return_Charity struct {
	Name        string
	Mission     string
	Street1     string
	Street2     string
	City        string
	Region      string
	Postal_code string
	Website_url string
}
