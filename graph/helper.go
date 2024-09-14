package graph

import (
	"fmt"
	"github.com/dkrizic/testserver/graph/model"
)

const (
	STATIC  = "static"
	DYNAMIC = "dynamic"
)

type InternalTagCategory struct {
	ID            string
	Name          string
	Discriminator string
	Parent        *string
	Format        *string
	Open          bool
}

type InternalTag struct {
	ID            string
	Name          *string
	Discriminator string
	Parent        *string
	Value         *string
}

func (itc InternalTagCategory) AsTagCategory() (model.TagCategory, error) {
	switch itc.Discriminator {
	case STATIC:
		tagCategory := model.StaticTagCategory{
			ID:     itc.ID,
			Name:   itc.Name,
			IsOpen: itc.Open,
		}
		return tagCategory, nil
	case DYNAMIC:
		tagCategory := model.DynamicTagCategory{
			ID:     itc.ID,
			Name:   itc.Name,
			Format: itc.Format,
		}
		return tagCategory, nil
	default:
		return nil, fmt.Errorf("unknown tag category discriminator: %s", itc.Discriminator)
	}
}

func (it InternalTag) AsTag() (model.Tag, error) {
	switch it.Discriminator {
	case STATIC:
		tag := model.StaticTag{
			ID: it.ID,
		}
		if it.Name != nil {
			tag.Name = *it.Name
		}
		return tag, nil
	case DYNAMIC:
		tag := model.DynamicTag{
			ID: it.ID,
		}
		if it.Value != nil {
			tag.Value = *it.Value
		}
		return tag, nil
	default:
		return nil, fmt.Errorf("unknown tag discriminator: %s", it.Discriminator)
	}
}
