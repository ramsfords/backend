package rapid_utils

import (
	v1 "github.com/ramsfords/types_gen/v1"
)

func GetPackageType(baseCommodity v1.Commodity) string {
	switch baseCommodity.PackageType {
	case 0:
	case 1:
		return "PLT"
	case 2:
	case 3:
	case 4:
	case 5:
	case 6:
	case 8:
		return "PLT"
	}
	return ""
}
