package main

import (
	"cmp"
	"slices"
	"strings"

	"github.com/maxim-lobanov/coding-contest-template-go/internal/algo"
	"github.com/maxim-lobanov/coding-contest-template-go/internal/cast"
)

func solution(input []string) string {
	stepsCount := cast.ParseInt(input[0])
	allPoints := make([]algo.Point3D, len(input))
	for i := 1; i < len(input); i++ {
		parts := strings.Split(input[i], ",")
		allPoints[i] = algo.Point3D{
			X: cast.ParseInt(parts[0]),
			Y: cast.ParseInt(parts[1]),
			Z: cast.ParseInt(parts[2]),
		}
	}

	sortedConnections := getSortedConnections(allPoints)
	foundComponents := findAllComponents(len(allPoints), sortedConnections, stepsCount)
	result := len(foundComponents[0]) * len(foundComponents[1]) * len(foundComponents[2])

	return cast.ToString(result)
}

func getSortedConnections(allPoints []algo.Point3D) [][2]int {
	allPairs := make([][2]int, 0)
	for i := 0; i < len(allPoints); i++ {
		for j := i + 1; j < len(allPoints); j++ {
			allPairs = append(allPairs, [2]int{i, j})
		}
	}

	slices.SortFunc(allPairs, func(indexPairA [2]int, indexPairB [2]int) int {
		p1PairA := allPoints[indexPairA[0]]
		p2PairA := allPoints[indexPairA[1]]
		p1PairB := allPoints[indexPairB[0]]
		p2PairB := allPoints[indexPairB[1]]

		distA := p1PairA.DistanceToPoint(p2PairA)
		distB := p1PairB.DistanceToPoint(p2PairB)
		return cmp.Compare(distA, distB)
	})

	return allPairs
}

func findAllComponents(pointsCount int, sortedConnections [][2]int, steps int) [][]int {
	components := make([][]int, pointsCount)
	componentsCount := pointsCount
	pointToComponentInd := make([]int, pointsCount)

	for i := 0; i < pointsCount; i++ {
		components[i] = []int{i}
		pointToComponentInd[i] = i
	}

	mergeComponents := func(firstPointInd, secondPointInd int) {
		firstComponentInd := pointToComponentInd[firstPointInd]
		secondComponentInd := pointToComponentInd[secondPointInd]
		if firstComponentInd == secondComponentInd {
			return
		}

		components[firstComponentInd] = append(components[firstComponentInd], components[secondComponentInd]...)
		for _, secondComponentPoint := range components[secondComponentInd] {
			pointToComponentInd[secondComponentPoint] = firstComponentInd
		}
		components[secondComponentInd] = nil

		componentsCount--
	}

	for i := 0; i < min(steps, len(sortedConnections)); i++ {
		mergeComponents(sortedConnections[i][0], sortedConnections[i][1])
	}

	remainComponents := [][]int{}
	for _, comp := range components {
		if comp != nil {
			remainComponents = append(remainComponents, comp)
		}
	}

	slices.SortFunc(remainComponents, func(first, second []int) int {
		return -cmp.Compare(len(first), len(second))
	})

	return remainComponents
}

/*
func debugConnections(allPoints []geometry.Point3D, connections [][2]int) {
	for _, pair := range connections {
		p1 := allPoints[pair[0]]
		p2 := allPoints[pair[1]]
		dist := p1.DistanceToPoint(p2)
		fmt.Printf("%d,%d,%d + %d,%d,%d = %f\n", p1.X, p1.Y, p1.Z, p2.X, p2.Y, p2.Z, dist)
	}
}
*/
