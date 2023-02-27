package utils

import v1 "github.com/ramsfords/types_gen/v1"

func GetComodityByIndex(index int, commodities []*v1.Commodity) *v1.Commodity {
	for i, j := range commodities {
		if int(j.Index) == index {
			return commodities[i]
		}
	}
	return nil
}
