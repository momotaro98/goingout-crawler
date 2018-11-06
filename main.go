package main

import "fmt"

type Spot struct {
	Id         int
	Address    string
	ImageUrl   string
	Latitude   string
	Longitude  string
	SpotTypeId int
}

type SpotType struct {
	Id   int
	Name string
}

func NewSpotType(name string) *SpotType {
	return &SpotType{Id: 1, Name: name}
}

func NewSpot(address string, imageUrl string, latitude string, longitude string, spotType SpotType) *Spot {
	return &Spot{
		Id:         1,
		Address:    address,
		ImageUrl:   imageUrl,
		Latitude:   latitude,
		Longitude:  longitude,
		SpotTypeId: spotType.Id,
	}
}

func main() {
	spotType := NewSpotType("Meal")
	spot := NewSpot("address", "url", "latitude", "longitude", *spotType) // TODO: Replace these args
	fmt.Println(spot)
}
