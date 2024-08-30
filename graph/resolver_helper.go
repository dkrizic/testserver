package graph

import (
	"context"
	"github.com/dkrizic/testserver/database"
	"github.com/dkrizic/testserver/graph/model"
	"log/slog"
)

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
