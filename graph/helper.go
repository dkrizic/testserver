package graph

import (
	"fmt"
	"github.com/dkrizic/testserver/graph/model"
)

type InternalTagCategory struct {
	ID            string
	Name          string
	Discriminator string
	Parent        *string
	Format        *string
	Open          bool
}

const (
	STATIC  = "static"
	DYNAMIC = "dynamic"
)

func (itc InternalTagCategory) AsTagCategory() (model.TagCategory, error) {
	switch itc.Discriminator {
	case STATIC:
		tagCategory := model.StaticTagCategory{
			ID:     itc.ID,
			Name:   itc.Name,
			IsOpen: itc.Open,
		}
		return tagCategory, nil
	case "dynamic":
		tagCategory := model.DynamicTagCategory{
			ID:   itc.ID,
			Name: itc.Name,
		}
		return tagCategory, nil
	default:
		return nil, fmt.Errorf("unknown tag category discriminator: %s", itc.Discriminator)
	}
}
