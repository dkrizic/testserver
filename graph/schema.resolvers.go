package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"fmt"

	"github.com/dkrizic/testserver/graph/model"
)

// CreateTag is the resolver for the createTag field.
func (r *mutationResolver) CreateTag(ctx context.Context, input model.NewTag) (*model.Tag, error) {
	panic(fmt.Errorf("not implemented: CreateTag - createTag"))
}

// CreateAsset is the resolver for the createAsset field.
func (r *mutationResolver) CreateAsset(ctx context.Context, input model.NewAsset) (*model.Asset, error) {
	panic(fmt.Errorf("not implemented: CreateAsset - createAsset"))
}

// CreateTagValue is the resolver for the createTagValue field.
func (r *mutationResolver) CreateTagValue(ctx context.Context, input model.NewTagValue) (*model.TagValue, error) {
	panic(fmt.Errorf("not implemented: CreateTagValue - createTagValue"))
}

// UpdateTagValue is the resolver for the updateTagValue field.
func (r *mutationResolver) UpdateTagValue(ctx context.Context, input model.UpdateTagValue) (*model.TagValue, error) {
	panic(fmt.Errorf("not implemented: UpdateTagValue - updateTagValue"))
}

// DeleteTagValue is the resolver for the deleteTagValue field.
func (r *mutationResolver) DeleteTagValue(ctx context.Context, input model.DeleteTagValue) (*model.TagValue, error) {
	panic(fmt.Errorf("not implemented: DeleteTagValue - deleteTagValue"))
}

// Tags is the resolver for the tags field.
func (r *queryResolver) Tags(ctx context.Context) ([]*model.Tag, error) {
	panic(fmt.Errorf("not implemented: Tags - tags"))
}

// Assets is the resolver for the assets field.
func (r *queryResolver) Assets(ctx context.Context) ([]*model.Asset, error) {
	panic(fmt.Errorf("not implemented: Assets - assets"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
