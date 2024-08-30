package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/dkrizic/testserver/database"
	"log/slog"

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
func (r *queryResolver) Tags(ctx context.Context, id *string) ([]*model.Tag, error) {
	slog.Info("Tags")
	var result *sql.Rows
	var err error
	if id != nil {
		result, err = r.dB.Query("SELECT id,name FROM tag WHERE id = ?", *id)
	} else {
		result, err = r.dB.Query("SELECT id,name FROM tag")
	}
	if err != nil {
		return nil, err
	}
	defer result.Close()

	tags := []*model.Tag{}
	for result.Next() {
		var tag model.Tag
		err := result.Scan(&tag.ID, &tag.Name)
		if err != nil {
			return nil, err
		}
		tag.Assets, err = r.assetsByTagId(ctx, tag.ID)
		tags = append(tags, &tag)
	}
	return tags, nil
}

// Assets is the resolver for the assets field.
func (r *queryResolver) Assets(ctx context.Context, id *string) ([]*model.Asset, error) {
	slog.Info("Assets")
	var result *sql.Rows
	var err error
	if id != nil {
		result, err = r.dB.Query("SELECT id,name FROM asset WHERE id = ?", *id)
	} else {
		result, err = r.dB.Query("SELECT id,name FROM asset")
	}
	if err != nil {
		return nil, err
	}
	defer result.Close()

	assets := []*model.Asset{}
	for result.Next() {
		var asset model.Asset
		err := result.Scan(&asset.ID, &asset.Name)
		if err != nil {
			return nil, err
		}
		asset.TagValues, err = r.tagValuesByAssetId(ctx, asset.ID)
		if err != nil {
			return nil, err
		}
		assets = append(assets, &asset)
	}
	return assets, nil
}

// TagValues is the resolver for the tagValues field.
func (r *queryResolver) TagValues(ctx context.Context, id *string) ([]*model.TagValue, error) {
	slog.Info("TagValues")
	var result *sql.Rows
	var err error
	if id != nil {
		result, err = r.dB.Query("SELECT id,tag_id,asset_id,value FROM tagvalue WHERE id = ?", *id)
	} else {
		result, err = r.dB.Query("SELECT id,tag_id,asset_id,value FROM tagvalue")
	}
	if err != nil {
		return nil, err
	}
	defer result.Close()

	tagValues := []*model.TagValue{}
	for result.Next() {
		tagValueTemp := &database.TagValue{}
		err := result.Scan(&tagValueTemp.ID, &tagValueTemp.TagID, &tagValueTemp.AssetID, &tagValueTemp.Value)
		var tagValue model.TagValue
		tagValue.ID = tagValueTemp.ID
		tagValue.Tag, err = r.tagById(ctx, tagValueTemp.TagID)
		tagValue.Asset, err = r.assetById(ctx, tagValueTemp.AssetID)
		tagValue.Value = tagValueTemp.Value

		if err != nil {
			return nil, err
		}
		tagValues = append(tagValues, &tagValue)
	}
	return tagValues, nil
}

func (r *Resolver) assetById(ctx context.Context, id string) (*model.Asset, error) {
	slog.Info("AssetById")
	result, err := r.dB.Query("SELECT id,name FROM asset WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	var asset model.Asset
	for result.Next() {
		err := result.Scan(&asset.ID, &asset.Name)
		if err != nil {
			return nil, err
		}
	}
	return &asset, nil
}

func (r *Resolver) assetsByTagId(ctx context.Context, id string) ([]*model.Asset, error) {
	slog.Info("AssetsByTagId")
	result, err := r.dB.Query("SELECT asset_id FROM tagvalue WHERE tag_id = ?", id)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	assets := []*model.Asset{}
	for result.Next() {
		var asset model.Asset
		err := result.Scan(&asset.ID)
		if err != nil {
			return nil, err
		}
		asset.TagValues, err = r.tagValuesByAssetId(ctx, asset.ID)
		if err != nil {
			return nil, err
		}
		assets = append(assets, &asset)
	}
	return assets, nil
}

func (r *Resolver) tagValuesByAssetId(ctx context.Context, id string) ([]*model.TagValue, error) {
	slog.Info("TagValuesByAssetId")
	result, err := r.dB.Query("SELECT id,tag_id,asset_id,value FROM tagvalue WHERE asset_id = ?", id)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	tagValues := []*model.TagValue{}
	for result.Next() {
		tagValueTemp := &database.TagValue{}
		err := result.Scan(&tagValueTemp.ID, &tagValueTemp.TagID, &tagValueTemp.AssetID, &tagValueTemp.Value)
		var tagValue model.TagValue
		tagValue.ID = tagValueTemp.ID
		tagValue.Tag, err = r.tagById(ctx, tagValueTemp.TagID)
		tagValue.Asset, err = r.assetById(ctx, tagValueTemp.AssetID)
		tagValue.Value = tagValueTemp.Value

		if err != nil {
			return nil, err
		}
		tagValues = append(tagValues, &tagValue)
	}
	slog.InfoContext(ctx, "TagValuesByAssetId", "tagValues", tagValues)
	return tagValues, nil
}

func (r *Resolver) tagValuesByTagId(ctx context.Context, id string) ([]*model.TagValue, error) {
	slog.Info("TagValuesByTagId")
	result, err := r.dB.Query("SELECT id,tag_id,asset_id,value FROM tagvalue WHERE tag_id = ?", id)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	tagValues := []*model.TagValue{}
	for result.Next() {
		tagValueTemp := &database.TagValue{}
		err := result.Scan(&tagValueTemp.ID, &tagValueTemp.TagID, &tagValueTemp.AssetID, &tagValueTemp.Value)
		var tagValue model.TagValue
		tagValue.ID = tagValueTemp.ID
		tagValue.Tag, err = r.tagById(ctx, tagValueTemp.TagID)
		tagValue.Asset, err = r.assetById(ctx, tagValueTemp.AssetID)
		tagValue.Value = tagValueTemp.Value

		if err != nil {
			return nil, err
		}
		tagValues = append(tagValues, &tagValue)
	}
	return tagValues, nil
}

func (r *Resolver) tagById(ctx context.Context, id string) (*model.Tag, error) {
	slog.Info("TagById")
	result, err := r.dB.Query("SELECT id,name FROM tag WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	var tag model.Tag
	for result.Next() {
		err := result.Scan(&tag.ID, &tag.Name)
		if err != nil {
			return nil, err
		}
	}
	return &tag, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
