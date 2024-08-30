package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/dkrizic/testserver/database"
	"github.com/dkrizic/testserver/graph/model"
)

// CreateTag is the resolver for the createTag field.
func (r *mutationResolver) CreateTag(ctx context.Context, input model.NewTag) (*model.Tag, error) {
	slog.Info("Create tag", "name", input.Name)
	result, err := r.dB.Exec("INSERT INTO tag (name) VALUES (?)", input.Name)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &model.Tag{ID: fmt.Sprintf("%d", id), Name: input.Name}, nil
}

// CreateAsset is the resolver for the createAsset field.
func (r *mutationResolver) CreateAsset(ctx context.Context, input model.NewAsset) (*model.Asset, error) {
	slog.Info("Create asset", "name", input.Name)
	result, err := r.dB.Exec("INSERT INTO asset (name) VALUES (?)", input.Name)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &model.Asset{ID: fmt.Sprintf("%d", id), Name: input.Name}, nil
}

// CreateTagValue is the resolver for the createTagValue field.
func (r *mutationResolver) CreateTagValue(ctx context.Context, input model.NewTagValue) (*model.TagValue, error) {
	slog.Info("Create tag value", "tag_id", input.TagID, "asset_id", input.AssetID, "value", input.Value)
	result, err := r.dB.Exec("INSERT INTO tagvalue (tag_id,asset_id,value) VALUES (?,?,?)", input.TagID, input.AssetID, input.Value)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &model.TagValue{ID: fmt.Sprintf("%d", id), Tag: &model.Tag{ID: input.TagID}, Asset: &model.Asset{ID: input.AssetID}, Value: input.Value}, nil
}

// UpdateTagValue is the resolver for the updateTagValue field.
func (r *mutationResolver) UpdateTagValue(ctx context.Context, input model.UpdateTagValue) (*model.TagValue, error) {
	slog.Info("Update tag value", "id", input.ID, "value", input.Value)
	_, err := r.dB.Exec("UPDATE tagvalue SET value = ? WHERE id = ?", input.Value, input.ID)
	if err != nil {
		return nil, err
	}
	return &model.TagValue{ID: input.ID, Value: input.Value}, nil
}

// DeleteTagValue is the resolver for the deleteTagValue field.
func (r *mutationResolver) DeleteTagValue(ctx context.Context, input model.DeleteTagValue) (*model.TagValue, error) {
	slog.Info("Delete tag value", "id", input.ID)
	result, err := r.dB.Exec("DELETE FROM tagvalue WHERE id = ?", input.ID)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &model.TagValue{ID: fmt.Sprintf("%d", id)}, nil
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

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
