package graph

import (
	"context"
	"fmt"
	"github.com/dkrizic/testserver/database"
	"github.com/dkrizic/testserver/graph/model"
	"github.com/dkrizic/testserver/telemetry"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"log/slog"
)

func (r *queryResolver) searchAssetName(ctx context.Context, text string) ([]*model.Asset, error) {
	ctx, span := telemetry.Tracer().Start(ctx, "searchAssetName")
	defer span.End()
	span.SetAttributes(attribute.String("text", text))

	slog.InfoContext(ctx, "searchAssetName", "text", text)
	query := "SELECT id,name FROM asset WHERE name LIKE ?"
	wildcard := fmt.Sprintf("%%%s%%", text)
	span.SetAttributes(
		attribute.String("db.query.text", query),
		attribute.String("db.parameter.text", wildcard))
	result, err := r.dB.Query(query, wildcard)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return nil, err
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
		assets = append(assets, &asset)
	}
	span.SetAttributes(attribute.Int("assets.count", len(assets)))
	span.SetStatus(codes.Ok, "SearchAssetName completed")
	return assets, nil
}

func (r *queryResolver) searchTagName(ctx context.Context, text string) ([]*model.Tag, error) {
	ctx, span := telemetry.Tracer().Start(ctx, "searchTagName")
	defer span.End()
	slog.DebugContext(ctx, "searchTagName", "text", text)

	query := "SELECT id,name FROM tag WHERE name LIKE ?"
	wildcard := fmt.Sprintf("%%%s%%", text)
	span.SetAttributes(
		attribute.String("db.query.text", query),
		attribute.String("db.parameter.text", wildcard))
	result, err := r.dB.Query(query, wildcard)
	if err != nil {
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
			return nil, err
		}
		tags = append(tags, &tag)
	}
	span.SetAttributes(attribute.Int("tags.count", len(tags)))
	span.SetStatus(codes.Ok, "SearchTagName completed")
	return tags, nil
}

func (r *queryResolver) searchTagValue(ctx context.Context, text string) ([]*model.TagValue, error) {
	ctx, span := telemetry.Tracer().Start(ctx, "searchTagValue")
	defer span.End()
	span.SetAttributes(attribute.String("text", text))

	slog.DebugContext(ctx, "searchTagValue", "text", text)
	query := "SELECT id,tag_id,asset_id,value FROM tagvalue WHERE value LIKE ?"
	wildcard := fmt.Sprintf("%%%s%%", text)
	result, err := r.dB.Query(query, wildcard)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}
	defer result.Close()

	tagValues := []*model.TagValue{}
	for result.Next() {
		tagValueTemp := &database.TagValue{}
		err := result.Scan(&tagValueTemp.ID, &tagValueTemp.TagID, &tagValueTemp.AssetID, &tagValueTemp.Value)
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			return nil, err
		}
		var tagValue model.TagValue
		tagValue.ID = tagValueTemp.ID
		tagValue.Value = tagValueTemp.Value

		tagValues = append(tagValues, &tagValue)
	}
	span.SetAttributes(attribute.Int("tagValues.count", len(tagValues)))
	span.SetStatus(codes.Ok, "SearchTagValue completed")
	return tagValues, nil
}
