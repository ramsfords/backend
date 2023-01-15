package rapid_utils

import v1 "github.com/ramsfords/types_gen/v1"

// FreightClass_CLASS_NONE FreightClass = 0
// FreightClass_Class55    FreightClass = 1
// FreightClass_Class60    FreightClass = 2
// FreightClass_Class65    FreightClass = 3
// FreightClass_Class70    FreightClass = 4
// FreightClass_Class775   FreightClass = 5
// FreightClass_Class85    FreightClass = 6
// FreightClass_Class925   FreightClass = 7
// FreightClass_Class100   FreightClass = 8
// FreightClass_Class110   FreightClass = 9
// FreightClass_Class125   FreightClass = 10
// FreightClass_Class150   FreightClass = 11
// FreightClass_Class175   FreightClass = 12
// FreightClass_Class200   FreightClass = 13
// FreightClass_Class250   FreightClass = 14
// FreightClass_Class300   FreightClass = 15
// FreightClass_Class400   FreightClass = 16
// FreightClass_Class500   FreightClass = 17
// FreightClass_Class18    FreightClass = 18
// FreightClass_Class50    FreightClass = 19
func GetClass(commodity v1.Commodity) string {
	switch commodity.FreightClass {
	case 19:
		return "50"
	case 1:
		return "55"
	case 2:
		return "60"
	case 3:
		return "65"
	case 4:
		return "70"
	case 5:
		return "77.5"
	case 6:
		return "85"
	case 7:
		return "92.5"
	case 8:
		return "100"
	case 9:
		return "110"
	case 10:
		return "125"
	case 11:
		return "150"
	case 12:
		return "175"
	case 13:
		return "200"
	case 14:
		return "250"
	case 15:
		return "300"
	case 16:
		return "400"
	case 17:
		return "500"
	default:
		return "0"
	}
}
