package main

import (
	common "aoc2024/aoccommon"
	"fmt"
	"slices"
	"sort"
	"strings"
)

func star1(edges [][2]int, adjacentByNode [][]int) {
	tValue := int('t' - 97)

	cliqueCount := 0
	for n1, adjacent := range adjacentByNode {
		for i, n2 := range adjacent {
			if n2 <= n1 {continue}
			for _, n3 := range adjacent[i+1:] {
				if n3 <= n1 {continue}
				// fmt.Println(n1, n2, n3, n1/26, n2/26, n3/26, tValue)
				if n1/26 == tValue || n2/26 == tValue || n3/26 == tValue {
					if slices.Contains(adjacentByNode[n2], n3) {
						cliqueCount++
					}
				}
			}
		}
	}
	fmt.Println("Number of 3-Cliques that starts with t:", cliqueCount)
}

// func star2(edges [][2]int, adjacentByNode [][]int) {
// 	cliquesOfCurrentSize := [][]int{}
// 	for _, edge := range edges {
// 		cliquesOfCurrentSize = append(cliquesOfCurrentSize, []int{min(edge[0], edge[1]), max(edge[0], edge[1])})
// 	}
// 	adjacencyMatrix := common.Array2D(26*26, 26*26, false)
// 	for n1, adjacent := range adjacentByNode {
// 		for _, n2 := range adjacent {
// 			adjacencyMatrix[n1][n2] = true
// 		}
// 	}

// 	for {
// 		fmt.Printf("Cliques of size %d: %d\n", len(cliquesOfCurrentSize[0]), len(cliquesOfCurrentSize))
// 		cliquesOfNextSize := [][]int{}

// 		for _, clique := range cliquesOfCurrentSize {
// 			newNodeLoop:
// 			for newNode := range adjacentByNode {
// 				for _, node := range clique {
// 					if !adjacencyMatrix[newNode][node] {
// 						continue newNodeLoop
// 					}
// 				}
// 				cliquesOfNextSize = append(cliquesOfNextSize, append(append([]int{}, clique...), newNode))
// 			}
// 		}
// 		if len(cliquesOfNextSize) > 0 {
// 			cliquesOfCurrentSize = cliquesOfNextSize
// 		} else {
// 			break
// 		}
// 	}
// }

func star2(edges [][2]int, adjacentByNode [][]int) {
	adjacencyMatrix := common.Array2D(26*26, 26*26, false)
	for n1, adjacent := range adjacentByNode {
		for _, n2 := range adjacent {
			adjacencyMatrix[n1][n2] = true
		}
	}

	nodes := []int{}
	for n1, adjacent := range adjacentByNode {
		if len(adjacent) > 0 {
			nodes = append(nodes, n1)
		}
	}
	cliques := [][]int{}
	BronKerbosch(&[]int{}, &nodes, []int{}, &adjacencyMatrix, &cliques)
	largestCliqueSize := 0
	for _, clique := range cliques {
		largestCliqueSize = max(largestCliqueSize, len(clique))
	}
	for _, clique := range cliques {
		if len(clique) == largestCliqueSize {
			names := []string{}
			for _, number := range clique {
				name := string([]rune{rune(number/26+97), rune(number%26+97)})
				names = append(names, name)
			}
			sort.Strings(names)
			fmt.Println("Password:", strings.Join(names, ","))
		}
	}
}

// Thanks wikipedia!
// 
// algorithm BronKerbosch2(R, P, X) is
//     if P and X are both empty then
//         report R as a maximal clique
//     choose a pivot vertex u in P ⋃ X
//     for each vertex v in P \ N(u) do
//         BronKerbosch2(R ⋃ {v}, P ⋂ N(v), X ⋂ N(v))
//         P := P \ {v}
//         X := X ⋃ {v}

func BronKerbosch(Rref *[]int, Pref *[]int, X []int, adjacencyMatrixRef *[][]bool, cliques *[][]int) {
	R, P, adjacencyMatrix := *Rref, *Pref, *adjacencyMatrixRef
	// fmt.Printf("BronKerbosch(%v, %v, %v)\n", R, P, X)
	if len(P) + len(X) == 0 {
		// report clique
		// fmt.Println("Maximal clique of size", len(R), R)
		*cliques = append(*cliques, append([]int{}, R...))
		return
	}
	var u int
	if len(P) > 0 {
		u = P[0]
	} else {
		u = X[0]
	}
	Rnew := make([]int, len(R)+1)
	copy(Rnew, R)
	Pnew := make([]int, 0, len(P))
	Xnew := make([]int, 0, len(X) + len(P))
	for i, v := range P {
		if !adjacencyMatrix[u][v] {
			Rnew[len(R)] = v
			Pnew = Pnew[:0]
			for _, p := range P {
				if p != -1 && adjacencyMatrix[v][p] {
					Pnew = append(Pnew, p)
				}
			}
			Xnew = Xnew[:0]
			for _, x := range X {
				if adjacencyMatrix[v][x] {
					Xnew = append(Xnew, x)
				}
			}
			BronKerbosch(&Rnew, &Pnew, Xnew, adjacencyMatrixRef, cliques)
			P[i] = -1
			X = append(X, v)
		}
	}
}


func main() {
	var lines = common.ReadLines("day23.txt")
	// var lines = common.ReadLines("test.txt")
	// lines = []string{"ta-rx","rx-bn","bn-ta"}
	// isNode := make([]bool, 26*26)
	adjacentByNode := make([][]int, 26*26)
	edges := [][2]int{}
	for _, line := range lines {
		// fmt.Println(int(line[0] - 97), int(line[0] - 97) * 26, int(line[1] - 97))
		n1 := int(line[0] - 97) * 26 + int(line[1] - 97)
		n2 := int(line[3] - 97) * 26 + int(line[4] - 97)
		edges = append(edges, [2]int{n1, n2})
		adjacentByNode[n1] = append(adjacentByNode[n1], n2)
		adjacentByNode[n2] = append(adjacentByNode[n2], n1)
	}
	for _, adjacent := range adjacentByNode {
		sort.Ints(adjacent)
	}
	star1(edges, adjacentByNode)
	star2(edges, adjacentByNode)
}