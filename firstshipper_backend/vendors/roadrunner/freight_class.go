package roadrunner

import v1 "github.com/ramsfords/types_gen/v1"

func GetActualFrightClass(class v1.FreightClass) float32 {
	switch class {
	case v1.FreightClass_CLASS50:
		return 50.0
	case v1.FreightClass_CLASS55:
		return 55.0
	case v1.FreightClass_CLASS60:
		return 60.0
	case v1.FreightClass_CLASS65:
		return 65.0
	case v1.FreightClass_CLASS70:
		return 70.0
	case v1.FreightClass_CLASS775:
		return 77.5
	case v1.FreightClass_CLASS85:
		return 85.0
	case v1.FreightClass_CLASS925:
		return 92.5
	case v1.FreightClass_CLASS100:
		return 100.0
	case v1.FreightClass_CLASS110:
		return 110.0
	case v1.FreightClass_CLASS125:
		return 125.0
	case v1.FreightClass_CLASS150:
		return 150.0
	case v1.FreightClass_CLASS175:
		return 175.0
	case v1.FreightClass_CLASS200:
		return 200.0
	case v1.FreightClass_CLASS250:
		return 250.0
	case v1.FreightClass_CLASS300:
		return 300.0
	case v1.FreightClass_CLASS400:
		return 400.0
	case v1.FreightClass_CLASS500:
		return 500.0
	default:
		return 0
	}
}
