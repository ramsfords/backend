package location_api

import (
	"context"

	v1 "github.com/ramsfords/types_gen/v1"
)

func (loc Location) EchoGetAllLocations(context.Context, *v1.Empty) (*v1.Locations, error) {
	return nil, nil
}
