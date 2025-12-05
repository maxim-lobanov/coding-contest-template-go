package main

import (
	"fmt"

	"github.com/maxim-lobanov/coding-contest-template-go/internal/cast"
)

func solution(input []string) string {
	allRanges := parseRanges(input)

	for {
		mergedRanges, changed := TryMergeRanges(allRanges)
		allRanges = mergedRanges
		if !changed {
			break
		}
	}

	result := 0
	for _, r := range allRanges {
		result += r.Sum()
	}

	return cast.ToString(result)
}

func parseRanges(input []string) []Range {
	allRanges := make([]Range, 0)

	lineIndex := 0
	for lineIndex < len(input) && input[lineIndex] != "" {
		r := Range{}
		fmt.Sscanf(input[lineIndex], "%d-%d", &r.Start, &r.End)
		allRanges = append(allRanges, r)
		lineIndex++
	}

	return allRanges
}

func TryMergeRanges(allRanges []Range) ([]Range, bool) {
	mergedRanges := make([]Range, 0)
	for i := 0; i < len(allRanges); i++ {
		newRange := allRanges[i]
		found := false
		for j := 0; j < len(mergedRanges); j++ {
			if mergedRanges[j].CanMerge(newRange) {
				mergedRanges[j] = mergedRanges[j].Merge(newRange)
				found = true
				break
			}
		}

		if !found {
			mergedRanges = append(mergedRanges, newRange)
		}
	}

	return mergedRanges, len(mergedRanges) < len(allRanges)
}

type Range struct {
	Start, End int
}

func (r Range) Sum() int {
	return r.End - r.Start + 1
}

func (r Range) CanMerge(other Range) bool {
	return (other.Start >= r.Start && other.Start <= r.End) || (other.End >= r.Start && other.End <= r.End) || (other.Start <= r.Start && other.End >= r.End)
}

func (r Range) Merge(other Range) Range {
	return Range{
		Start: min(r.Start, other.Start),
		End:   max(r.End, other.End),
	}
}
