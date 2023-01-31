package quote

import v1 "github.com/ramsfords/types_gen/v1"

// FreightClass_ClASSNONE FreightClass = 0
// FreightClass_CLASS50   FreightClass = 1
// FreightClass_CLASS55   FreightClass = 2
// FreightClass_CLASS60   FreightClass = 3
// FreightClass_CLASS65   FreightClass = 4
// FreightClass_CLASS70   FreightClass = 5
// FreightClass_CLASS775  FreightClass = 6
// FreightClass_CLASS85   FreightClass = 7
// FreightClass_CLASS925  FreightClass = 8
// FreightClass_CLASS100  FreightClass = 9
// FreightClass_CLASS110  FreightClass = 10
// FreightClass_CLASS125  FreightClass = 11
// FreightClass_CLASS150  FreightClass = 12
// FreightClass_CLASS175  FreightClass = 13
// FreightClass_CLASS200  FreightClass = 14
// FreightClass_CLASS250  FreightClass = 15
// FreightClass_CLASS300  FreightClass = 16
// FreightClass_CLASS400  FreightClass = 17
// FreightClass_CLASS500  FreightClass = 18
func GetClass(commodity v1.Commodity) string {
	switch commodity.FreightClass {
	case 0:
		return "0"
	case 1:
		return "50"
	case 2:
		return "55"
	case 3:
		return "60"
	case 4:
		return "65"
	case 5:
		return "70"
	case 6:
		return "77.5"
	case 7:
		return "85"
	case 8:
		return "92.5"
	case 9:
		return "100"
	case 10:
		return "110"
	case 11:
		return "125"
	case 12:
		return "150"
	case 13:
		return "175"
	case 14:
		return "200"
	case 15:
		return "250"
	case 16:
		return "300"
	case 17:
		return "400"
	case 18:
		return "500"
	default:
		return "0"
	}
}
