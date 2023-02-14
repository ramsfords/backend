package quote

import (
	v1 "github.com/ramsfords/types_gen/v1"
)

// PackageType_PACKAGENONE PackageType = 0
// PackageType_PALLET      PackageType = 1
// PackageType_BOX         PackageType = 2
// PackageType_BUNDLE      PackageType = 3
// PackageType_LOOSE       PackageType = 4
// PackageType_ROLL        PackageType = 5
// PackageType_PIECES      PackageType = 6
// PackageType_CASE        PackageType = 7
// PackageType_BUCKET      PackageType = 8
// PackageType_BAG         PackageType = 9
// PackageType_BALE        PackageType = 10
// PackageType_CARTON      PackageType = 11
// PackageType_CRATE       PackageType = 12
// PackageType_CYLINDER    PackageType = 13
// PackageType_DRUMS       PackageType = 14
// PackageType_ROLE        PackageType = 15
// PackageType_SKID        PackageType = 16
// PackageType_TOTE        PackageType = 17
// PackageType_TUBE        PackageType = 18
// PackageType_REEL        PackageType = 19
// PackageType_PAIL        PackageType = 20

func GetPackageType(baseCommodity v1.Commodity) string {
	switch baseCommodity.PackageType {
	case 0:
		return "PLT"
	case 1:
		return "PLT"
	case 2:
		return "BOX"
	case 3:
		return "BUNDLE"
	case 4:
		return "LOOSE"
	case 5:
		return "ROLL"
	case 6:
		return "PIECES"
	case 8:
		return "CASE"
	case 9:
		return "BAG"
	case 10:
		return "BALE"
	case 11:
		return "CARTON"
	case 12:
		return "CRATE"
	case 13:
		return "CYLINDER"
	case 14:
		return "DRUMS"
	case 15:
		return "ROLE"
	case 16:
		return "SKID"
	case 17:
		return "TOTE"
	case 18:
		return "TUBE"
	case 19:
		return "REEL"
	case 20:
		return "PAIL"
	}
	return "PLT"
}
