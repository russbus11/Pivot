package marvel

import (
	"fmt"
	"log"
)

func main() {
	client := marvel.NewClient(marvel.BaseURL)
	chars, err := client.GetCast(5)
	if err != nil {
		log.Fatal(err)
	}

	for _, char := range chars {
		fmt.Printf("Name: %v | Description: %v\n", char.Name, char.Description)
	}
}
