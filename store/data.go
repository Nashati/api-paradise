package store

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/api-paradise/model"
)

func readJokes() []model.Joke {
	var jokes []model.Joke
	jsonFile, err := os.Open("./store/data/jokes.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened jokes.json")
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(byteValue, &jokes)
	if err != nil {
		fmt.Println(err)
	}

	return jokes
}

func readCategories() []model.Category {
	var categories []model.Category
	jsonFile, err := os.Open("./store/data/categories.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened categories.json")
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(byteValue, &categories)
	if err != nil {
		fmt.Println(err)
	}

	return categories
}
