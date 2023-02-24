package mymaps

import "fmt"

func MyMaps() {

	celebs := map[string]int{
		"Nicholas cage":     50,
		"Selena gomez":      21,
		"Jude law":          41,
		"Scarlet Johannson": 29,
	}

	for celeb, age := range celebs {
		fmt.Printf("Celebrity Name -> %v Celebrity Age -> %v\n", celeb, age)
	}

	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{40.68433, -74.399967}
	fmt.Println(m["Bell Labs"])

	companies["Splice"] = CompaniesLocations{40.68433, -74.39998}
	fmt.Println(companies["Splice"])

	delete(companies, "Splice")
	fmt.Printf("%v\n", companies)
	name, ok := companies["Splice"]
	fmt.Printf("Key 'Splice' is present?: %t - value: %v\n", ok, name)
}

type Vertex struct {
	Lat, Lon float64
}

var m map[string]Vertex

type CompaniesLocations struct {
	Lat, Lon float64
}

var companies = map[string]CompaniesLocations{
	"Google":   {40.55667, -10.4456},
	"FaceBook": {80.44567, -45660.0},
}
