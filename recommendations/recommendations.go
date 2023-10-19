package recommendations

import (
	"fmt"
	"sort"

	"gonum.org/v1/gonum/mat"
)

/*
WALKTHROUGH:

Given the sample user-item interaction matrix:

	⎡ 5  3  0  1 ⎤
	⎢ 4  0  0  1 ⎥
	⎢ 1  1  0  5 ⎥
	⎢ 1  0  0  4 ⎥
	⎣ 0  1  5  4 ⎦

Each row represents a user, and each column represents an item.
The values in the matrix represent interactions (e.g., purchase counts) between users and items.
For instance, the entry (1,2) with a value of 3 means that the first user has interacted with the second item 3 times.

1. The program starts by converting this array into a dense matrix format using GoNum's mat.NewDense() method.
2. Next, Singular Value Decomposition (SVD) is performed on this matrix. SVD breaks the matrix down into three other matrices: U, Sigma, and V^T.
   - U and V^T are orthogonal matrices that represent the "features" of users and items, respectively.
   - Sigma is a diagonal matrix that contains all singular values (representing the "strength" of these features).
3. The program then approximates the original matrix by multiplying U, Sigma, and V^T together. This approximated matrix gives us estimated interaction values even for those user-item pairs that had originally a value of 0 (i.e., no interaction).
4. Using this approximated matrix, the program can make recommendations. For instance, for the second user (index 1), the program checks which items this user hasn't interacted with (value of 0 in the original matrix).
5. From this subset of items, the program then ranks them based on their values in the approximated matrix, suggesting that items with higher approximated values would be more relevant for the user.
6. The result is printed, indicating which items are recommended for the second user.

By using this methodology, one can generate recommendations for any user in the provided dataset based on their past interactions and the underlying patterns detected by the SVD.
*/

func MapData(data []float64, userIndex int) []int {

	// Convert the sample data into a dense matrix.
	R := mat.NewDense(5, 4, data)

	// Initialize an SVD object to decompose the matrix.
	var svd mat.SVD
	// Decompose the matrix using SVD. If factorization fails, an error will be raised.
	ok := svd.Factorize(R, mat.SVDThin)
	if !ok {
		panic("SVD factorization failed")
	}

	// Extract the U matrix from the SVD result.
	U := &mat.Dense{}
	svd.UTo(U)

	// Extract the singular values (diagonal of Sigma matrix).
	Sigma := svd.Values(nil)

	// Extract the VT (transpose of V) matrix from the SVD result.
	VT := &mat.Dense{}
	svd.VTo(VT)

	// Compute the approximated matrix by multiplying U, Sigma, and VT together.
	var approxR mat.Dense
	approxR.Product(U, mat.NewDiagDense(len(Sigma), Sigma), VT)

	// Print the approximated matrix to the console.
	fmt.Println("Approximated Matrix:")
	fmt.Println(mat.Formatted(&approxR))

	// Get the items recommended for the user at index 1 (second user) and print them.
	recommendedItems := getRecommendations(userIndex, R, &approxR)

	return recommendedItems
}

func getRecommendations(user int, originalMatrix, approxMatrix *mat.Dense) []int {
	// Get the total number of items (columns) from the matrix.
	numItems := originalMatrix.RawMatrix().Cols
	// Initialize a slice to store recommended item indices.
	recommendations := []int{}

	// Loop through each item for the specified user.
	for i := 0; i < numItems; i++ {
		// If the user hasn't interacted with the item (value is 0 in the original matrix),
		// add it to the list of items to be potentially recommended.
		if originalMatrix.At(user, i) == 0 {
			recommendations = append(recommendations, i)
		}
	}

	// Sort the list of potential recommendations based on their values in the approximated matrix.
	// Items with higher values in the approximated matrix are ranked first.
	sort.Slice(recommendations, func(i, j int) bool {
		return approxMatrix.At(user, recommendations[i]) > approxMatrix.At(user, recommendations[j])
	})

	// Return the sorted list of recommended items.
	return recommendations
}
