package weather

import (
	"context"
	"encoding/json"
	"error"
	"fmt"
	"github.com/kr/pretty"
	"googlemaps.github.io/maps"
	"log"
)

type WeatherData struct {
	ID       string
	Name     string
	Typ      string
	Country  string
	Timezone Timezone
	Location Location
}

//NewLocation(town string, destination string)
func (wd *WeatherData) NewLocation(town string, destin string) {
	c, err := maps.NewClient( /*insert API key here, string*/ )
	if err != nil {
		log.Fatalf("Fatal error: %s", err)
	}
	r := &maps.DirectionsRequest{
		Origin:      town,
		Destination: destin,
	}
	route, _, err := c.Directions(context.Background(), r)
	if err != nl {
		log.Fatalf("Fatal error: %s", err)
	}
	pretty.Println(route)
}

func (wd *WeatherData) Elevation(APIkey string) {
	// Do elevation here
	//docs
	// https://developers.google.com/maps/documentation/elevation/start
}

func (wd *WeatherData) GenerateJSON() (json, error) {
	jsonOutput, err := json.Marshal(wd)
	if err != nil {
		fmt.Println("Error: %s", err)
	}
	return jsonOutput, err
}
