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
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

// CreateTag is the resolver for the createTag field.
func (r *mutationResolver) CreateTag(ctx context.Context, tagName string) (*model.Tag, error) {
	slog.Info("Create tag", "name", tagName)
	result, err := r.dB.Exec("INSERT INTO tag (name) VALUES (?)", tagName)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &model.Tag{ID: fmt.Sprintf("%d", id), Name: tagName}, nil
}

// CreateAsset is the resolver for the createAsset field.
func (r *mutationResolver) CreateAsset(ctx context.Context, assetName string) (*model.Asset, error) {
	slog.Info("Create asset", "name", assetName)
	result, err := r.dB.Exec("INSERT INTO asset (name) VALUES (?)", assetName)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &model.Asset{ID: fmt.Sprintf("%d", id), Name: assetName}, nil
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

// Search is the resolver for the search field.
func (r *queryResolver) Search(ctx context.Context, input model.Search, skip *int, limit *int) (*model.SearchResult, error) {
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(
		attribute.String("query", input.Text),
		attribute.Bool("searchAssetName", input.SearchAssetName),
		attribute.Bool("searchTagName", input.SearchTagName),
		attribute.Bool("searchTagValue", input.SearchTagValue))

	slog.Info("Search", "query", input.Text, "searchAssetName", input.SearchAssetName, "searchTagName", input.SearchTagName, "searchTagValue", input.SearchTagValue)
	searchResult := &model.SearchResult{}
	if input.SearchAssetName {
		assets, err := r.searchAssetName(ctx, input.Text)
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			return nil, err
		}
		searchResult.Assets = assets
	}
	if input.SearchTagName {
		tags, err := r.searchTagName(ctx, input.Text)
		if err != nil {
			return nil, err
		}
		searchResult.Tags = tags
	}
	if input.SearchTagValue {
		tagValues, err := r.searchTagValue(ctx, input.Text)
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			return nil, err
		}
		searchResult.TagValues = tagValues
	}

	span.SetStatus(codes.Ok, "Search completed")
	return searchResult, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

func (r *queryResolver) Tag(ctx context.Context, id *string, skip *int, limit *int) ([]*model.Tag, error) {
	span := trace.SpanFromContext(ctx)
	if id != nil {
		span.SetAttributes(attribute.String("id", *id))
	}
	slog.Info("Tags", "id", id, "skip", skip, "limit", limit)
	var result *sql.Rows
	var err error
	span.SetAttributes(
		attribute.String("db.system", "mysql"),
		attribute.String("db.operation.name", "select"),
		attribute.Int("skip", *skip),
		attribute.Int("limit", *limit))
	if id != nil {
		query := "SELECT id,name FROM tag WHERE id = ?"
		span.SetAttributes(
			attribute.String("db.query.text", query),
			attribute.String("db.query.parameter.id", *id))
		result, err = r.dB.Query(query, *id)
	} else {
		query := "SELECT id,name FROM tag LIMIT ?,?"
		span.SetAttributes(attribute.String("db.query.text", query))
		result, err = r.dB.Query(query, *skip, *limit)
	}
	if err != nil {
		slog.WarnContext(ctx, "Error querying database", "error", err)
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}
	defer result.Close()

	tags := []*model.Tag{}
	for result.Next() {
		var tag model.Tag
		err := result.Scan(&tag.ID, &tag.Name)
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			return nil, err
		}
		tag.Assets, err = r.assetsByTagId(ctx, tag.ID)
		tags = append(tags, &tag)
	}
	span.SetAttributes(attribute.Int("tags.count", len(tags)))
	span.SetStatus(codes.Ok, "Tags completed")
	return tags, nil
}

func (r *queryResolver) Asset(ctx context.Context, id *string, skip *int, limit *int) ([]*model.Asset, error) {
	span := trace.SpanFromContext(ctx)
	if id != nil {
		span.SetAttributes(attribute.String("id", *id))
	}
	slog.Info("Assets", "id", id, "skip", skip, "limit", limit)
	var result *sql.Rows
	var err error
	if id != nil {
		query := "SELECT id,name FROM asset WHERE id = ?"
		span.SetAttributes(
			attribute.String("db.query.text", query),
			attribute.String("db.query.parameter.id", *id))
		result, err = r.dB.Query(query, *id)
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			return nil, err
		}
	} else {
		query := "SELECT id,name FROM asset LIMIT ?,?"
		span.SetAttributes(attribute.String("db.query.text", query))
		result, err = r.dB.Query(query, skip, limit)
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			return nil, err
		}
	}
	defer result.Close()

	assets := []*model.Asset{}
	for result.Next() {
		var asset model.Asset
		err := result.Scan(&asset.ID, &asset.Name)
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			return nil, err
		}
		asset.TagValues, err = r.tagValuesByAssetId(ctx, asset.ID)
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			return nil, err
		}
		assets = append(assets, &asset)
	}
	span.SetAttributes(attribute.Int("assets.count", len(assets)))
	span.SetStatus(codes.Ok, "Assets completed")
	return assets, nil
}
func (r *queryResolver) TagValue(ctx context.Context, id *string, skip *int, limit *int) ([]*model.TagValue, error) {
	span := trace.SpanFromContext(ctx)
	if id != nil {
		span.SetAttributes(attribute.String("id", *id))
	}
	slog.Info("TagValues", "id", id, "skip", skip, "limit", limit)
	var result *sql.Rows
	var err error
	if id != nil {
		query := "SELECT id,tag_id,asset_id,value FROM tagvalue WHERE id = ? LIMIT ?,?"
		span.SetAttributes(
			attribute.String("db.query.text", query),
			attribute.String("db.query.parameter.id", *id))
		result, err = r.dB.Query(query, *id, *skip, *limit)
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			return nil, err
		}
	} else {
		query := "SELECT id,tag_id,asset_id,value FROM tagvalue"
		span.SetAttributes(attribute.String("db.query.text", query))
		result, err = r.dB.Query(query)
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			return nil, err
		}
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

	span.SetAttributes(attribute.Int("tagValues.count", len(tagValues)))
	span.SetStatus(codes.Ok, "TagValues completed")
	return tagValues, nil
}
