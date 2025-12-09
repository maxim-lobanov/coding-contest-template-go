package main

import (
	"cmp"
	"slices"
	"strings"

	"github.com/maxim-lobanov/coding-contest-template-go/internal/cast"
	"github.com/maxim-lobanov/coding-contest-template-go/internal/geometry"
)

func solution(input []string) string {
	allPoints := make([]geometry.Point3D, len(input))
	for i := 0; i < len(input); i++ {
		parts := strings.Split(input[i], ",")
		allPoints[i] = geometry.Point3D{
			X: cast.ParseInt(parts[0]),
			Y: cast.ParseInt(parts[1]),
			Z: cast.ParseInt(parts[2]),
		}
	}

	sortedConnections := getSortedConnections(allPoints)
	stepsCount := findSingleComponent(len(allPoints), sortedConnections)
	result := allPoints[sortedConnections[stepsCount][0]].X * allPoints[sortedConnections[stepsCount][1]].X

	return cast.ToString(result)
}

func getSortedConnections(allPoints []geometry.Point3D) [][2]int {
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

func findSingleComponent(pointsCount int, sortedConnections [][2]int) int {
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

	for i := 0; i < len(sortedConnections); i++ {
		mergeComponents(sortedConnections[i][0], sortedConnections[i][1])

		if componentsCount <= 1 {
			return i
		}
	}

	return -1
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
