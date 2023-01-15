package utils

import v1 "github.com/ramsfords/types_gen/v1"

func GetFreightClass(req *v1.Commodity) v1.FreightClass {
	return req.FreightClass

}
