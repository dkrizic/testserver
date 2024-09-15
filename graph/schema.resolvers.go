package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/dkrizic/testserver/graph/model"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// Identity is the resolver for the identity field.
func (r *accessResolver) Identity(ctx context.Context, obj *model.Access) (model.Identity, error) {
	slog.InfoContext(ctx, "Identity(byAccess)", "id", obj.ID)
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(
		attribute.String("db.system", "mysql"),
		attribute.String("db.operation.name", "select"))
	span.SetAttributes(attribute.String("id", obj.ID))

	query := "SELECT id,email FROM user WHERE id = (SELECT identity_id FROM access WHERE id = ?)"
	span.SetAttributes(
		attribute.String("db.query.text", query),
		attribute.String("db.parameter.id", obj.ID))
	result, err := r.dB.QueryContext(ctx, query, obj.ID)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	var user model.User
	if result.Next() {
		err := result.Scan(&user.ID, &user.Email)
		if err != nil {
			slog.ErrorContext(ctx, "Error scanning row", "error", err, "id", obj.ID, "type", "user")
			return nil, err
		}
		return &user, nil
	}

	query = "SELECT id,name FROM `group` WHERE id = (SELECT identity_id FROM access WHERE id = ?)"
	span.SetAttributes(
		attribute.String("db.query.text", query),
		attribute.String("db.parameter.id", obj.ID))
	result, err = r.dB.QueryContext(ctx, query, obj.ID)
	if err != nil {
		return nil, nil
	}
	defer result.Close()

	var group model.Group
	if result.Next() {
		err := result.Scan(&group.ID, &group.Name)
		if err != nil {
			slog.ErrorContext(ctx, "Error scanning row", "error", err, "identity_id", obj.Identity.GetID(), "type", "group")
			return nil, err
		}
		return &group, nil
	}

	return nil, fmt.Errorf("No group or identitiy for asset id %s found", obj.ID)
}

// Asset is the resolver for the asset field.
func (r *accessResolver) Asset(ctx context.Context, obj *model.Access) (*model.Asset, error) {
	slog.InfoContext(ctx, "Asset(byAccess)", "id", obj.ID)
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(
		attribute.String("db.system", "mysql"),
		attribute.String("db.operation.name", "select"))
	span.SetAttributes(
		attribute.String("id", obj.ID))
	query := "SELECT id,name FROM asset WHERE id = (SELECT asset_id FROM access WHERE id = ?)"
	span.SetAttributes(
		attribute.String("db.query.text", query),
		attribute.String("db.parameter.id", obj.ID))
	result, err := r.dB.QueryContext(ctx, query, obj.ID)
	if err != nil {
		return nil, nil
	}
	defer result.Close()

	var asset model.Asset
	if result.Next() {
		err := result.Scan(&asset.ID, &asset.Name)
		if err != nil {
			slog.ErrorContext(ctx, "Error scanning row", "error", err, "asset_id", obj.Asset.ID, "identity_id", obj.Identity.GetID(), "type", "asset")
			return nil, err
		}
		return &asset, nil
	} else {
		return nil, nil
	}
}

// Accesses is the resolver for the accesses field.
func (r *assetResolver) Accesses(ctx context.Context, obj *model.Asset, skip *int, limit *int) ([]*model.Access, error) {
	slog.InfoContext(ctx, "Accesses(forAsset)", "id", obj.ID, "skip", skip, "limit", limit)
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(
		attribute.String("db.system", "mysql"),
		attribute.String("db.operation.name", "select"),
		attribute.Int("skip", *skip),
		attribute.Int("limit", *limit))
	span.SetAttributes(attribute.String("id", obj.ID))
	query := "SELECT id, permission FROM access WHERE asset_id  = ? limit ?,?"
	span.SetAttributes(
		attribute.String("db.query.text", query),
		attribute.String("db.parameter.id", obj.ID))
	result, err := r.dB.QueryContext(ctx, query, obj.ID, skip, limit)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	accesses := []*model.Access{}
	for result.Next() {
		var access model.Access
		err := result.Scan(&access.ID, &access.Permission)
		if err != nil {
			slog.ErrorContext(ctx, "Error scanning row", "error", err, "id", obj.ID, "type", "access")
			return nil, err
		}
		accesses = append(accesses, &access)
	}
	return accesses, nil
}

// Files is the resolver for the files field.
func (r *assetResolver) Files(ctx context.Context, obj *model.Asset, skip *int, limit *int) ([]*model.File, error) {
	slog.InfoContext(ctx, "Files(forAsset)", "id", obj.ID, "skip", skip, "limit", limit)
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(
		attribute.String("db.system", "mysql"),
		attribute.String("db.operation.name", "select"),
		attribute.Int("skip", *skip),
		attribute.Int("limit", *limit))
	span.SetAttributes(attribute.String("id", obj.ID))
	query := "SELECT id,name,size,mimetype FROM file WHERE asset_id = ? limit ?,?"
	span.SetAttributes(
		attribute.String("db.query.text", query),
		attribute.String("db.parameter.id", obj.ID))
	result, err := r.dB.QueryContext(ctx, query, obj.ID, skip, limit)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	files := []*model.File{}
	for result.Next() {
		var file model.File
		err := result.Scan(&file.ID, &file.Name, &file.Size, &file.MimeType)
		if err != nil {
			slog.ErrorContext(ctx, "Error scanning row", "error", err, "id", obj.ID, "type", "file")
			return nil, err
		}
		files = append(files, &file)
	}
	return files, nil
}

// Tags is the resolver for the tags field.
func (r *assetResolver) Tags(ctx context.Context, obj *model.Asset, skip *int, limit *int) ([]model.Tag, error) {
	slog.InfoContext(ctx, "Tags(forAsset)", "id", obj.ID, "skip", skip, "limit", limit)
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(
		attribute.String("db.system", "mysql"),
		attribute.String("db.operation.name", "select"),
		attribute.Int("skip", *skip),
		attribute.Int("limit", *limit))
	span.SetAttributes(attribute.String("id", obj.ID))
	query := "SELECT tag.id,tag.name,tag.parent_tag_id,tag.value,tag.discriminator FROM tag,asset_tag WHERE tag.id = asset_tag.tag_id and asset_tag.asset_id = ? limit ?,?"
	span.SetAttributes(
		attribute.String("db.query.text", query),
		attribute.String("db.parameter.id", obj.ID))
	result, err := r.dB.QueryContext(ctx, query, obj.ID, skip, limit)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	tags := []model.Tag{}
	for result.Next() {
		var it InternalTag
		err := result.Scan(&it.ID, &it.Name, &it.Parent, &it.Value, &it.Discriminator)
		if err != nil {
			slog.ErrorContext(ctx, "Error scanning row", "error", err, "id", obj.ID, "type", "tag")
			return nil, err
		}
		slog.DebugContext(ctx, "Scanned row", slog.Group("row", "id", it.ID, "name", it.Name, "parent", it.Parent, "value", it.Value, "discriminator", it.Discriminator))
		tag, err := it.AsTag()
		if err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}
	return tags, nil
}

// TagCategory is the resolver for the tagCategory field.
func (r *dynamicTagResolver) TagCategory(ctx context.Context, obj *model.DynamicTag) (*model.DynamicTagCategory, error) {
	slog.InfoContext(ctx, "TagCategory(byDynamicTag)", "id", obj.ID)
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(
		attribute.String("db.system", "mysql"),
		attribute.String("db.operation.name", "select"))
	span.SetAttributes(attribute.String("id", obj.ID))
	query := "SELECT id,name,format FROM tagcategory WHERE id = (SELECT tagcategory_id FROM tag WHERE id = ?)"
	span.SetAttributes(
		attribute.String("db.query.text", query),
		attribute.String("db.parameter.id", obj.ID))
	result, err := r.dB.QueryContext(ctx, query, obj.ID)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	var dtc model.DynamicTagCategory
	if result.Next() {
		err := result.Scan(&dtc.ID, &dtc.Name, &dtc.Format)
		if err != nil {
			slog.ErrorContext(ctx, "Error scanning row", "error", err, "id", obj.ID, "type", "tagcategory")
			return nil, err
		}
		slog.Debug("Entity", slog.Group("entity", "id", dtc.ID, "name", dtc.Name, "format", dtc.Format))
		return &dtc, nil
	} else {
		return nil, nil
	}
}

// ParentTag is the resolver for the parentTag field.
func (r *dynamicTagResolver) ParentTag(ctx context.Context, obj *model.DynamicTag) (model.Tag, error) {
	slog.InfoContext(ctx, "ParentTag(byDynamicTag)", "id", obj.ID)
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(
		attribute.String("db.system", "mysql"),
		attribute.String("db.operation.name", "select"))
	span.SetAttributes(attribute.String("id", obj.ID))
	query := "SELECT id,name,parent_tag_id,value,discriminator FROM tag WHERE id = (SELECT parent_tag_id FROM tag WHERE id = ?)"
	span.SetAttributes(
		attribute.String("db.query.text", query),
		attribute.String("db.parameter.id", obj.ID))
	result, err := r.dB.QueryContext(ctx, query, obj.ID)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	var it InternalTag
	if result.Next() {
		err := result.Scan(&it.ID, &it.Name, &it.Parent, &it.Value, &it.Discriminator)
		if err != nil {
			slog.ErrorContext(ctx, "Error scanning row", "error", err, "id", obj.ID, "type", "tag")
			return nil, err
		}
		return it.AsTag()
	} else {
		return nil, nil
	}
}

// ChildTags is the resolver for the childTags field.
func (r *dynamicTagResolver) ChildTags(ctx context.Context, obj *model.DynamicTag, skip *int, limit *int) ([]model.Tag, error) {
	slog.InfoContext(ctx, "ChildTags(byDynamicTag)", "id", obj.ID, "skip", skip, "limit", limit)
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(
		attribute.String("db.system", "mysql"),
		attribute.String("db.operation.name", "select"),
		attribute.Int("skip", *skip),
		attribute.Int("limit", *limit))
	span.SetAttributes(attribute.String("id", obj.ID))
	query := "SELECT id,name,parent_tag_id,value,discriminator FROM tag WHERE parent_tag_id = ?"
	span.SetAttributes(
		attribute.String("db.query.text", query),
		attribute.String("db.parameter.id", obj.ID))
	result, err := r.dB.QueryContext(ctx, query, obj.ID)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	tags := []model.Tag{}
	for result.Next() {
		var it InternalTag
		err := result.Scan(&it.ID, &it.Name, &it.Parent, &it.Value, &it.Discriminator)
		if err != nil {
			slog.ErrorContext(ctx, "Error scanning row", "error", err, "id", obj.ID, "type", "tag")
			return nil, err
		}
		tag, err := it.AsTag()
		slog.DebugContext(ctx, "Entity", slog.Group("entity", "id", tag.GetID()))
		if err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}
	return tags, nil
}

// ParentTagCategory is the resolver for the parentTagCategory field.
func (r *dynamicTagCategoryResolver) ParentTagCategory(ctx context.Context, obj *model.DynamicTagCategory) (model.TagCategory, error) {
	slog.InfoContext(ctx, "ParentTagCategory(byDynamicTagCategory)", "id", obj.ID)
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(
		attribute.String("db.system", "mysql"),
		attribute.String("db.operation.name", "select"))
	span.SetAttributes(attribute.String("id", obj.ID))
	query := "SELECT p.id,p.name,p.parent,p.discriminator, p.format, p.open FROM tagcategory p, tagcategory this WHERE this.parent = p.id and this.id = ?"
	span.SetAttributes(
		attribute.String("db.query.text", query),
		attribute.String("db.parameter.id", obj.ID))
	result, err := r.dB.QueryContext(ctx, query, obj.ID)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	var itc InternalTagCategory
	if result.Next() {
		err := result.Scan(&itc.ID, &itc.Name, &itc.Parent, &itc.Discriminator, &itc.Format, &itc.Open)
		if err != nil {
			slog.ErrorContext(ctx, "Error scanning row", "error", err, "id", obj.ID, "type", "tagcategory")
			return nil, err
		}
		return itc.AsTagCategory()
	} else {
		return nil, nil
	}
}

// RootTags is the resolver for the rootTags field.
func (r *dynamicTagCategoryResolver) RootTags(ctx context.Context, obj *model.DynamicTagCategory, skip *int, limit *int) ([]model.Tag, error) {
	slog.InfoContext(ctx, "RootTags(byDynamicTagCategory)", "id", obj.ID, "skip", skip, "limit", limit)
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(
		attribute.String("db.system", "mysql"),
		attribute.String("db.operation.name", "select"),
		attribute.Int("skip", *skip),
		attribute.Int("limit", *limit))
	span.SetAttributes(attribute.String("id", obj.ID))
	query := "SELECT id,name,parent_tag_id,value,discriminator FROM tag WHERE parent_tag_id is null and tagcategory_id = ? limit ?,?"
	span.SetAttributes(
		attribute.String("db.query.text", query),
		attribute.String("db.parameter.id", obj.ID))
	result, err := r.dB.QueryContext(ctx, query, obj.ID, skip, limit)
	if err != nil {
		return nil, nil
	}
	defer result.Close()

	tags := []model.Tag{}
	for result.Next() {
		var it InternalTag
		err := result.Scan(&it.ID, &it.Name, &it.Parent, &it.Value, &it.Discriminator)
		if err != nil {
			slog.ErrorContext(ctx, "Error scanning row", "error", err, "id", obj.ID, "type", "tag")
			return nil, err
		}
		slog.DebugContext(ctx, "Scanned row", slog.Group("row", "id", it.ID, "name", it.Name, "parent", it.Parent, "value", it.Value, "discriminator", it.Discriminator))
		tag, err := it.AsTag()
		if err != nil {
			return tags, err
		}
		tags = append(tags, tag)
	}

	return tags, nil
}

// Users is the resolver for the users field.
func (r *groupResolver) Users(ctx context.Context, obj *model.Group, skip *int, limit *int) ([]*model.User, error) {
	slog.InfoContext(ctx, "Users(forGroup)", "id", obj.ID, "skip", skip, "limit", limit)
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(
		attribute.String("db.system", "mysql"),
		attribute.String("db.operation.name", "select"),
		attribute.Int("skip", *skip),
		attribute.Int("limit", *limit))
	span.SetAttributes(attribute.String("id", obj.ID))
	query := "SELECT user.id,user.email FROM user,user_group where user_group.user_id = user.id and user_group.group_id = ? limit ?,?"
	span.SetAttributes(
		attribute.String("db.query.text", query),
		attribute.String("db.parameter.id", obj.ID))
	result, err := r.dB.QueryContext(ctx, query, obj.ID, skip, limit)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	users := []*model.User{}
	for result.Next() {
		var user model.User
		err := result.Scan(&user.ID, &user.Email)
		if err != nil {
			slog.ErrorContext(ctx, "Error scanning row", "error", err, "id", obj.ID, "type", "user")
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

// Accesses is the resolver for the accesses field.
func (r *groupResolver) Accesses(ctx context.Context, obj *model.Group, skip *int, limit *int) ([]*model.Access, error) {
	slog.InfoContext(ctx, "Accesses(forGroup)", "id", obj.ID, "skip", skip, "limit", limit)
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(
		attribute.String("db.system", "mysql"),
		attribute.String("db.operation.name", "select"),
		attribute.Int("skip", *skip),
		attribute.Int("limit", *limit))
	span.SetAttributes(attribute.String("id", obj.ID))
	query := "SELECT id, permission FROM access WHERE discriminator = 'group' and access.identity_id  = ? limit ?,?"
	span.SetAttributes(
		attribute.String("db.query.text", query),
		attribute.String("db.parameter.id", obj.ID))
	result, err := r.dB.QueryContext(ctx, query, obj.ID, skip, limit)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	accesses := []*model.Access{}
	for result.Next() {
		var access model.Access
		err := result.Scan(&access.ID, &access.Permission)
		if err != nil {
			slog.ErrorContext(ctx, "Error scanning row", "error", err, "id", obj.ID, "type", "access")
			return nil, err
		}
		accesses = append(accesses, &access)
	}
	return accesses, nil
}

// Asset is the resolver for the asset field.
func (r *queryResolver) Asset(ctx context.Context, id *string, skip *int, limit *int) ([]*model.Asset, error) {
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(
		attribute.String("db.system", "mysql"),
		attribute.String("db.operation.name", "select"),
		attribute.Int("skip", *skip),
		attribute.Int("limit", *limit))
	if id != nil {
		span.SetAttributes(attribute.String("id", *id))
	}
	slog.InfoContext(ctx, "Asset", "id", id, "skip", skip, "limit", limit)
	var result *sql.Rows
	var err error
	if id != nil {
		query := "SELECT id,name FROM asset WHERE id = ?"
		span.SetAttributes(
			attribute.String("db.query.text", query),
			attribute.String("db.query.parameter.id", *id))
		result, err = r.dB.QueryContext(ctx, query, *id)
		if err != nil {
			return nil, err
		}
	} else {
		query := "SELECT id,name FROM asset LIMIT ?,?"
		span.SetAttributes(attribute.String("db.query.text", query))
		result, err = r.dB.QueryContext(ctx, query, skip, limit)
		if err != nil {
			return nil, err
		}
	}
	defer result.Close()

	assets := []*model.Asset{}
	for result.Next() {
		var asset model.Asset
		err := result.Scan(&asset.ID, &asset.Name)
		if err != nil {
			slog.ErrorContext(ctx, "Error scanning row", "error", err, "id", id, "type", "asset")
			return nil, err
		}
		assets = append(assets, &asset)
	}
	span.SetAttributes(attribute.Int("assets.count", len(assets)))
	return assets, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id *string, skip *int, limit *int) ([]*model.User, error) {
	slog.InfoContext(ctx, "User", "id", id, "skip", skip, "limit", limit)
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(
		attribute.String("db.system", "mysql"),
		attribute.String("db.operation.name", "select"),
		attribute.Int("skip", *skip),
		attribute.Int("limit", *limit))
	if id != nil {
		span.SetAttributes(attribute.String("id", *id))
	}
	var result *sql.Rows
	var err error
	if id != nil {
		query := "SELECT id,email FROM user WHERE id = ?"
		span.SetAttributes(
			attribute.String("db.query.text", query),
			attribute.String("db.query.parameter.id", *id))
		result, err = r.dB.QueryContext(ctx, query, *id)
		if err != nil {
			return nil, err
		}
	} else {
		query := "SELECT id,email FROM user LIMIT ?,?"
		span.SetAttributes(attribute.String("db.query.text", query))
		result, err = r.dB.QueryContext(ctx, query, *skip, *limit)
		if err != nil {
			return nil, err
		}
	}
	defer result.Close()

	users := []*model.User{}
	for result.Next() {
		var user model.User
		err := result.Scan(&user.ID, &user.Email)
		if err != nil {
			slog.ErrorContext(ctx, "Error scanning row", "error", err, "id", id, "type", "user")
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

// Group is the resolver for the group field.
func (r *queryResolver) Group(ctx context.Context, id *string, skip *int, limit *int) ([]*model.Group, error) {
	slog.InfoContext(ctx, "Group", "id", id, "skip", skip, "limit", limit)
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(
		attribute.String("db.system", "mysql"),
		attribute.String("db.operation.name", "select"),
		attribute.Int("skip", *skip),
		attribute.Int("limit", *limit))
	if id != nil {
		span.SetAttributes(attribute.String("id", *id))
	}
	var result *sql.Rows
	var err error
	if id != nil {
		query := "SELECT id,name FROM group WHERE id = ?"
		span.SetAttributes(
			attribute.String("db.query.text", query),
			attribute.String("db.query.parameter.id", *id))
		result, err = r.dB.QueryContext(ctx, query, *id)
		if err != nil {
			return nil, err
		}
	} else {
		query := "SELECT id,name FROM `group` LIMIT ?,?"
		span.SetAttributes(attribute.String("db.query.text", query))
		result, err = r.dB.QueryContext(ctx, query, *skip, *limit)
		if err != nil {
			return nil, err
		}
	}
	defer result.Close()

	groups := []*model.Group{}
	for result.Next() {
		var group model.Group
		err := result.Scan(&group.ID, &group.Name)
		if err != nil {
			slog.ErrorContext(ctx, "Error scanning row", "error", err, "id", id, "type", "group")
			return nil, err
		}
		groups = append(groups, &group)
	}
	return groups, nil
}

// Identity is the resolver for the identity field.
func (r *queryResolver) Identity(ctx context.Context, skip *int, limit *int) ([]model.Identity, error) {
	slog.InfoContext(ctx, "Identity", "skip", skip, "limit", limit)
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(
		attribute.String("db.system", "mysql"),
		attribute.String("db.operation.name", "select"),
		attribute.Int("skip", *skip),
		attribute.Int("limit", *limit))
	query := "SELECT id,email FROM user LIMIT ?,?"
	span.SetAttributes(attribute.String("db.query.text", query))

	result, err := r.dB.QueryContext(ctx, query, *skip, *limit)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	identities := []model.Identity{}
	for result.Next() {
		var user model.User
		err := result.Scan(&user.ID, &user.Email)
		if err != nil {
			slog.ErrorContext(ctx, "Error scanning row", "error", err, "type", "user")
			return nil, err
		}
		identities = append(identities, user)
	}

	query = "SELECT id,name FROM `group` LIMIT ?,?"
	span.SetAttributes(attribute.String("db.query.text", query))
	result, err = r.dB.QueryContext(ctx, query, *skip, *limit)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	for result.Next() {
		var group model.Group
		err := result.Scan(&group.ID, &group.Name)
		if err != nil {
			slog.ErrorContext(ctx, "Error scanning row", "error", err, "type", "group")
			return nil, err
		}
		identities = append(identities, group)
	}

	return identities, nil
}

// TagCategory is the resolver for the tagCategory field.
func (r *queryResolver) TagCategory(ctx context.Context, id *string) (model.TagCategory, error) {
	slog.InfoContext(ctx, "TagCategory", "id", id)
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(
		attribute.String("db.system", "mysql"),
		attribute.String("db.operation.name", "select"))
	if id != nil {
		span.SetAttributes(attribute.String("id", *id))
	}

	query := "SELECT id,name FROM tagcategory WHERE id = ?"
	span.SetAttributes(attribute.String("db.query.text", query))
	result, err := r.dB.QueryContext(ctx, query, *id)
	if err != nil {
		return model.StaticTagCategory{}, err
	}
	defer result.Close()

	var itc InternalTagCategory
	for result.Next() {
		err := result.Scan(&itc.ID, &itc.Name, &itc.Discriminator, &itc.Parent, &itc.Format, &itc.Open)
		if err != nil {
			slog.ErrorContext(ctx, "Error scanning row", "error", err, "id", id, "type", "tagcategory")
			return nil, err
		}
	}
	return itc.AsTagCategory()
}

// TagCategories is the resolver for the tagCategories field.
func (r *queryResolver) TagCategories(ctx context.Context, skip *int, limit *int) ([]model.TagCategory, error) {
	slog.InfoContext(ctx, "TagCategories", "skip", skip, "limit", limit)
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(
		attribute.String("db.system", "mysql"),
		attribute.String("db.operation.name", "select"),
		attribute.Int("skip", *skip),
		attribute.Int("limit", *limit))
	query := "SELECT id,name,discriminator,parent,format,open FROM tagcategory LIMIT ?,?"
	span.SetAttributes(attribute.String("db.query.text", query))
	result, err := r.dB.QueryContext(ctx, query, *skip, *limit)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	tagCategories := []model.TagCategory{}
	for result.Next() {
		var itc InternalTagCategory
		err := result.Scan(&itc.ID, &itc.Name, &itc.Discriminator, &itc.Parent, &itc.Format, &itc.Open)
		if err != nil {
			return nil, err
		}
		tagCategory, err := itc.AsTagCategory()
		if err != nil {
			return nil, err
		}
		tagCategories = append(tagCategories, tagCategory)
	}

	return tagCategories, nil
}

// Tag is the resolver for the tag field.
func (r *queryResolver) Tag(ctx context.Context, id *string) (model.Tag, error) {
	slog.InfoContext(ctx, "Tag", "id", id)
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(
		attribute.String("db.system", "mysql"),
		attribute.String("db.operation.name", "select"))
	if id != nil {
		span.SetAttributes(attribute.String("id", *id))
	}

	query := "SELECT id,name,parent_tag_id,value,discriminator FROM tag WHERE id = ?"
	span.SetAttributes(attribute.String("db.query.text", query))
	result, err := r.dB.QueryContext(ctx, query, *id)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	var it InternalTag
	for result.Next() {
		err := result.Scan(&it.ID, &it.Name, &it.Parent, &it.Value, &it.Discriminator)
		if err != nil {
			slog.ErrorContext(ctx, "Error scanning row", "error", err, "id", id, "type", "tag")
			return nil, err
		}
	}
	return it.AsTag()
}

// TagCategory is the resolver for the tagCategory field.
func (r *staticTagResolver) TagCategory(ctx context.Context, obj *model.StaticTag) (*model.StaticTagCategory, error) {
	slog.InfoContext(ctx, "TagCategory(byStaticTag)", "id", obj.ID)
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(
		attribute.String("db.system", "mysql"),
		attribute.String("db.operation.name", "select"))
	span.SetAttributes(attribute.String("id", obj.ID))
	query := "SELECT tc.id,tc.name,tc.open FROM tagcategory tc, tag t WHERE tc.id = t.tagcategory_id and t.id = ?"
	span.SetAttributes(
		attribute.String("db.query.text", query),
		attribute.String("db.parameter.id", obj.ID))
	result, err := r.dB.QueryContext(ctx, query, obj.ID)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	var stc model.StaticTagCategory
	if result.Next() {
		err := result.Scan(&stc.ID, &stc.Name, &stc.IsOpen)
		if err != nil {
			slog.ErrorContext(ctx, "Error scanning row", "error", err, "id", obj.ID, "type", "tagcategory")
			return nil, err
		}
		slog.Debug("Scanned row", slog.Group("row", "id", stc.ID, "name", stc.Name, "open", stc.IsOpen))
		return &stc, nil
	} else {
		slog.ErrorContext(ctx, "No rows returned", "id", obj.ID, "type", "tagcategory")
		return nil, nil
	}
}

// ParentTag is the resolver for the parentTag field.
func (r *staticTagResolver) ParentTag(ctx context.Context, obj *model.StaticTag) (model.Tag, error) {
	slog.InfoContext(ctx, "ParentTag(byStaticTag)", "id", obj.ID)
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(
		attribute.String("db.system", "mysql"),
		attribute.String("db.operation.name", "select"))
	span.SetAttributes(attribute.String("id", obj.ID))
	query := "SELECT id,name,parent_tag_id,value,discriminator FROM tag WHERE id = (SELECT parent_tag_id FROM tag WHERE id = ?)"
	span.SetAttributes(
		attribute.String("db.query.text", query),
		attribute.String("db.parameter.id", obj.ID))
	result, err := r.dB.QueryContext(ctx, query, obj.ID)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	var it InternalTag
	if result.Next() {
		err := result.Scan(&it.ID, &it.Name, &it.Parent, &it.Value, &it.Discriminator)
		if err != nil {
			return nil, err
		}
		return it.AsTag()
	} else {
		return nil, nil
	}
}

// ChildTags is the resolver for the childTags field.
func (r *staticTagResolver) ChildTags(ctx context.Context, obj *model.StaticTag, skip *int, limit *int) ([]model.Tag, error) {
	slog.InfoContext(ctx, "ChildTags(byStaticTag)", "id", obj.ID, "skip", skip, "limit", limit)
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(
		attribute.String("db.system", "mysql"),
		attribute.String("db.operation.name", "select"),
		attribute.Int("skip", *skip),
		attribute.Int("limit", *limit))
	span.SetAttributes(attribute.String("id", obj.ID))
	query := "SELECT id,name,parent_tag_id,value,discriminator FROM tag WHERE parent_tag_id = ?"
	span.SetAttributes(
		attribute.String("db.query.text", query),
		attribute.String("db.parameter.id", obj.ID))
	result, err := r.dB.QueryContext(ctx, query, obj.ID)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	tags := []model.Tag{}
	for result.Next() {
		var it InternalTag
		err := result.Scan(&it.ID, &it.Name, &it.Parent, &it.Value, &it.Discriminator)
		if err != nil {
			return nil, err
		}
		slog.DebugContext(ctx, "Scanned row", slog.Group("row", "id", it.ID, "name", it.Name, "parent", it.Parent, "value", it.Value, "discriminator", it.Discriminator))
		tag, err := it.AsTag()
		if err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}
	return tags, nil
}

// ParentTagCategory is the resolver for the parentTagCategory field.
func (r *staticTagCategoryResolver) ParentTagCategory(ctx context.Context, obj *model.StaticTagCategory) (model.TagCategory, error) {
	slog.InfoContext(ctx, "ParentTagCategory(byStaticTagCategory)", "id", obj.ID)
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(
		attribute.String("db.system", "mysql"),
		attribute.String("db.operation.name", "select"))
	span.SetAttributes(attribute.String("id", obj.ID))
	query := "SELECT p.id,p.name,p.parent,p.discriminator, p.format, p.open FROM tagcategory p, tagcategory this WHERE this.parent = p.id and this.id = ?"
	span.SetAttributes(
		attribute.String("db.query.text", query),
		attribute.String("db.parameter.id", obj.ID))
	result, err := r.dB.QueryContext(ctx, query, obj.ID)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	if result.Next() {
		var itc InternalTagCategory
		err := result.Scan(&itc.ID, &itc.Name, &itc.Parent, &itc.Discriminator, &itc.Format, &itc.Open)
		if err != nil {
			return nil, err
		}
		return itc.AsTagCategory()
	} else {
		return nil, nil
	}
}

// RootTags is the resolver for the rootTags field.
func (r *staticTagCategoryResolver) RootTags(ctx context.Context, obj *model.StaticTagCategory, skip *int, limit *int) ([]model.Tag, error) {
	slog.InfoContext(ctx, "RootTags(byStaticTageCategory)", "id", obj.ID, "skip", skip, "limit", limit)
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(
		attribute.String("db.system", "mysql"),
		attribute.String("db.operation.name", "select"),
		attribute.Int("skip", *skip),
		attribute.Int("limit", *limit))
	span.SetAttributes(attribute.String("id", obj.ID))
	query := "SELECT id,name,parent_tag_id,value,discriminator FROM tag WHERE parent_tag_id is null and tagcategory_id = ? limit ?,?"
	span.SetAttributes(
		attribute.String("db.query.text", query),
		attribute.String("db.parameter.id", obj.ID))
	result, err := r.dB.QueryContext(ctx, query, obj.ID, skip, limit)
	if err != nil {
		return nil, nil
	}
	defer result.Close()

	tags := []model.Tag{}
	for result.Next() {
		var it InternalTag
		err := result.Scan(&it.ID, &it.Name, &it.Parent, &it.Value, &it.Discriminator)
		if err != nil {
			return nil, err
		}
		slog.DebugContext(ctx, "Scanned row", slog.Group("row", "id", it.ID, "name", it.Name, "parent", it.Parent, "value", it.Value, "discriminator", it.Discriminator))
		tag, err := it.AsTag()
		if err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}
	return tags, nil
}

// Groups is the resolver for the groups field.
func (r *userResolver) Groups(ctx context.Context, obj *model.User, skip *int, limit *int) ([]*model.Group, error) {
	slog.InfoContext(ctx, "Groups(forUser)", "id", obj.ID, "skip", skip, "limit", limit)
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(
		attribute.String("db.system", "mysql"),
		attribute.String("db.operation.name", "select"),
		attribute.Int("skip", *skip),
		attribute.Int("limit", *limit))
	span.SetAttributes(attribute.String("id", obj.ID))
	query := "SELECT `group`.id,`group`.name FROM `group`,user_group where user_group.group_id = `group`.id and user_group.user_id = ? limit ?,?"
	span.SetAttributes(
		attribute.String("db.query.text", query),
		attribute.String("db.parameter.id", obj.ID))
	result, err := r.dB.QueryContext(ctx, query, obj.ID, skip, limit)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	groups := []*model.Group{}
	for result.Next() {
		var group model.Group
		err := result.Scan(&group.ID, &group.Name)
		if err != nil {
			return nil, err
		}
		groups = append(groups, &group)
	}
	return groups, nil
}

// Accesses is the resolver for the accesses field.
func (r *userResolver) Accesses(ctx context.Context, obj *model.User, skip *int, limit *int) ([]*model.Access, error) {
	slog.InfoContext(ctx, "Accesses(forUser)", "id", obj.ID, "skip", skip, "limit", limit)
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(
		attribute.String("db.system", "mysql"),
		attribute.String("db.operation.name", "select"),
		attribute.Int("skip", *skip),
		attribute.Int("limit", *limit))
	span.SetAttributes(attribute.String("id", obj.ID))
	query := "SELECT id,permission FROM access WHERE discriminator='user' AND access.identity_id = ? limit ?,?"
	span.SetAttributes(
		attribute.String("db.query.text", query),
		attribute.String("db.parameter.id", obj.ID))
	result, err := r.dB.QueryContext(ctx, query, obj.ID, skip, limit)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	accesses := []*model.Access{}
	for result.Next() {
		var access model.Access
		err := result.Scan(&access.ID, &access.Permission)
		if err != nil {
			return nil, err
		}
		accesses = append(accesses, &access)
	}
	return accesses, nil
}

// Access returns AccessResolver implementation.
func (r *Resolver) Access() AccessResolver { return &accessResolver{r} }

// Asset returns AssetResolver implementation.
func (r *Resolver) Asset() AssetResolver { return &assetResolver{r} }

// DynamicTag returns DynamicTagResolver implementation.
func (r *Resolver) DynamicTag() DynamicTagResolver { return &dynamicTagResolver{r} }

// DynamicTagCategory returns DynamicTagCategoryResolver implementation.
func (r *Resolver) DynamicTagCategory() DynamicTagCategoryResolver {
	return &dynamicTagCategoryResolver{r}
}

// Group returns GroupResolver implementation.
func (r *Resolver) Group() GroupResolver { return &groupResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// StaticTag returns StaticTagResolver implementation.
func (r *Resolver) StaticTag() StaticTagResolver { return &staticTagResolver{r} }

// StaticTagCategory returns StaticTagCategoryResolver implementation.
func (r *Resolver) StaticTagCategory() StaticTagCategoryResolver {
	return &staticTagCategoryResolver{r}
}

// User returns UserResolver implementation.
func (r *Resolver) User() UserResolver { return &userResolver{r} }

type accessResolver struct{ *Resolver }
type assetResolver struct{ *Resolver }
type dynamicTagResolver struct{ *Resolver }
type dynamicTagCategoryResolver struct{ *Resolver }
type groupResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type staticTagResolver struct{ *Resolver }
type staticTagCategoryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
