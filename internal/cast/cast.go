package cast

import "strconv"

func ParseInt(s string) int {
	res, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return int(res)
}

func ParseInt64(s string) int64 {
	res, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return res
}

func ParseFloat64(s string) float64 {
	res, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}
	return res
}
