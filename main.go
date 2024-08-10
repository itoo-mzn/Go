package main

import (
	"log"
	"os"

	"github.com/gocarina/gocsv"
)

type Country struct {
	Name       string `csv:"国名"`
	ISOCode    string `csv:"ISOコード"`
	Population int    `csv:"人口"`
}

func main() {
	lines := []Country{
		{Name: "アメリカ", ISOCode: "US", Population: 31},
		{Name: "日本", ISOCode: "JP", Population: 2},
		{Name: "中国", ISOCode: "CN", Population: 40},
	}

	f, err := os.Create("country.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if err := gocsv.MarshalFile(&lines, f); err != nil {
		log.Fatal(err)
	}

}
