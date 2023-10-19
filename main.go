package main

import (
	"fmt"

	"github.com/deanrtaylor1/product-rec-go/recommendations"
)

func main() {
	// Initialize the user-item interaction matrix with sample data.
	data := []float64{
		5, 3, 0, 1,
		4, 0, 0, 1,
		1, 1, 0, 5,
		1, 0, 0, 4,
		0, 1, 5, 4,
	}

	userIndex := 1
	recommendedItems := recommendations.MapData(data, userIndex)

	fmt.Printf("Recommended items for user %d: %v\n", userIndex+1, recommendedItems)

}
