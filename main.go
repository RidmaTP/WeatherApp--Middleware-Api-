package main

import (
	"encoding/json"

	"io/ioutil"
	"log"
	"net/http"
	"github.com/labstack/echo/v4"
)

type Weather struct {
	Condition[] struct {
		Id int `json : "id"`
		Main string `json:"main"`
		Icon string `json:"icon"`
		Description string `json:"description"`
	}`json:"weather"`
	Main struct {
		Temperature float64 `json:"temp"`
		Preassure   float64 `json:"pressure"`
		Humidity    float64 `json:"humidity"`
	} `json:"main"`
	Sys struct {
		
		Message string `json:"message"`
		Country string  `json:"country"`
		
	} `json:"sys"`
}

type sendWeather struct{
	ID int `json:"id"`
	Main string `json :"condition"`
	Icon string `json:"icon"`
	Temperature float64 `json:"temp"`
	Preassure   float64 `json:"pressure"`
	Humidity    float64 `json:"humidity"`
	Message string `json:"message"`
	Country string  `json:"country"`
	Description string `json:"description"`

}

func main() {

	e := echo.New()

	e.GET("/:location",setStats )

	e.Logger.Fatal(e.Start(":8000"))
}

func setStats(c echo.Context) error { // unmarshals the json to struct stats and print it in page body
	var weather Weather
	location := c.Param("location")
	err := json.Unmarshal([]byte(getStats(location)), &weather) //parsing json to struct Weather
	if err != nil {
		log.Printf("error while decoding stats")
	}
	var sendWeather sendWeather
	sendWeather.Temperature = weather.Main.Temperature
	sendWeather.Preassure = weather.Main.Preassure
	sendWeather.Humidity = weather.Main.Humidity
	sendWeather.ID = weather.Condition[0].Id
	sendWeather.Icon = weather.Condition[0].Icon
	sendWeather.Main = weather.Condition[0].Main
	sendWeather.Country = weather.Sys.Country
	sendWeather.Message = weather.Sys.Message
	sendWeather.Description = weather.Condition[0].Description
	defer 
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
  	c.Response().WriteHeader(http.StatusOK)
  	return json.NewEncoder(c.Response()).Encode(sendWeather)

}

func getStats(location string) string { //getting the json response from open weather api and returning it as a string

	url := "https://api.openweathermap.org/data/2.5/weather?q=" + location + "&appid=c15efb199beb65917bc377eab272588d"
	response, err := http.Get(url)
	if err != nil {
		log.Printf("error while getting json")
	}

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("error while getting stats")
	}

	return string(bytes)

}
func postStats(){

}
