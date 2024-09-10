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
	"go.opentelemetry.io/otel/trace"
)

// TagValues is the resolver for the tagValues field.
func (r *assetResolver) TagValues(ctx context.Context, obj *model.Asset) ([]*model.TagValue, error) {
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("id", obj.ID))
	slog.InfoContext(ctx, "TagValues(byAsset)", "id", obj.ID)

	query := "SELECT id,tag_id,asset_id,value FROM tagvalue WHERE asset_id = ?"
	span.SetAttributes(
		attribute.String("db.query.text", query),
		attribute.String("db.parameter.id", obj.ID))
	result, err := r.dB.Query(query, obj.ID)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	tagValues := []*model.TagValue{}
	for result.Next() {
		tagValueTemp := &database.TagValue{}
		err := result.Scan(&tagValueTemp.ID, &tagValueTemp.TagID, &tagValueTemp.AssetID, &tagValueTemp.Value)
		if err != nil {
			return nil, err
		}
		var tagValue model.TagValue
		tagValue.ID = tagValueTemp.ID
		tagValue.Value = tagValueTemp.Value
		tagValues = append(tagValues, &tagValue)
	}

	span.SetAttributes(attribute.Int("tagValues.count", len(tagValues)))
	return tagValues, nil
}

// Users is the resolver for the users field.
func (r *groupResolver) Users(ctx context.Context, obj *model.Group, skip *int, limit *int) ([]*model.User, error) {
	slog.Info("Users(forGroup)", "id", obj.ID, "skip", skip, "limit", limit)
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
	result, err := r.dB.Query(query, obj.ID, skip, limit)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	users := []*model.User{}
	for result.Next() {
		var user model.User
		err := result.Scan(&user.ID, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

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

// AddGroup is the resolver for the addGroup field.
func (r *mutationResolver) AddGroup(ctx context.Context, name string) (*model.Group, error) {
	slog.Info("Add group", "name", name)
	result, err := r.dB.Exec("INSERT INTO `group` (name) VALUES (?)", name)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &model.Group{ID: fmt.Sprintf("%d", id), Name: name}, nil
}

// AddUser is the resolver for the addUser field.
func (r *mutationResolver) AddUser(ctx context.Context, email string) (*model.User, error) {
	slog.Info("Add user", "email", email)
	result, err := r.dB.Exec("INSERT INTO user (email) VALUES (?)", email)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &model.User{ID: fmt.Sprintf("%d", id), Email: email}, nil
}

// AddUsertoGroup is the resolver for the addUsertoGroup field.
func (r *mutationResolver) AddUsertoGroup(ctx context.Context, userID string, groupID string) (*model.Group, error) {
	slog.Info("Add user to group", "userID", userID, "groupID", groupID)
	result, err := r.dB.Exec("INSERT INTO user_group (user_id,group_id) VALUES (?,?)", userID, groupID)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &model.Group{ID: fmt.Sprintf("%d", id)}, nil
}

// RemoveUserFromGroup is the resolver for the removeUserFromGroup field.
func (r *mutationResolver) RemoveUserFromGroup(ctx context.Context, userID string, groupID string) (*model.Group, error) {
	slog.Info("Remove user from group", "userID", userID, "groupID", groupID)
	_, err := r.dB.Exec("DELETE FROM user_group WHERE user_id = ? AND group_id = ?", userID, groupID)
	if err != nil {
		return nil, err
	}

	slog.Debug("Loading group", "id", groupID)
	result, err := r.dB.Query("SELECT id,name FROM `group` WHERE id = ?", groupID)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	var group model.Group
	for result.Next() {
		err := result.Scan(&group.ID, &group.Name)
		if err != nil {
			return nil, err
		}
	}

	return &group, nil
}

// AssignPermission is the resolver for the assignPermission field.
func (r *mutationResolver) AssignPermission(ctx context.Context, identityID string, assetID string, permission model.Permission) (*model.Assignment, error) {
	panic(fmt.Errorf("not implemented: AssignPermission - assignPermission"))
}

// RemovePermission is the resolver for the removePermission field.
func (r *mutationResolver) RemovePermission(ctx context.Context, identityID string, assetID string) (*model.Assignment, error) {
	panic(fmt.Errorf("not implemented: RemovePermission - removePermission"))
}

// Tag is the resolver for the tag field.
func (r *queryResolver) Tag(ctx context.Context, id *string, skip *int, limit *int) ([]*model.Tag, error) {
	span := trace.SpanFromContext(ctx)
	if id != nil {
		span.SetAttributes(attribute.String("id", *id))
	}
	slog.Info("Tag", "id", id, "skip", skip, "limit", limit)
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
	return tags, nil
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
	slog.Info("Asset", "id", id, "skip", skip, "limit", limit)
	var result *sql.Rows
	var err error
	if id != nil {
		query := "SELECT id,name FROM asset WHERE id = ?"
		span.SetAttributes(
			attribute.String("db.query.text", query),
			attribute.String("db.query.parameter.id", *id))
		result, err = r.dB.Query(query, *id)
		if err != nil {
			return nil, err
		}
	} else {
		query := "SELECT id,name FROM asset LIMIT ?,?"
		span.SetAttributes(attribute.String("db.query.text", query))
		result, err = r.dB.Query(query, skip, limit)
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
			return nil, err
		}
		assets = append(assets, &asset)
	}
	span.SetAttributes(attribute.Int("assets.count", len(assets)))
	return assets, nil
}

// TagValue is the resolver for the tagValue field.
func (r *queryResolver) TagValue(ctx context.Context, id *string, skip *int, limit *int) ([]*model.TagValue, error) {
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(
		attribute.String("db.system", "mysql"),
		attribute.String("db.operation.name", "select"),
		attribute.Int("skip", *skip),
		attribute.Int("limit", *limit))
	if id != nil {
		span.SetAttributes(attribute.String("id", *id))
	}
	slog.Info("TagValue", "id", id, "skip", *skip, "limit", *limit)
	var result *sql.Rows
	var err error
	if id != nil {
		query := "SELECT id,tag_id,asset_id,value FROM tagvalue WHERE id = ? LIMIT ?,?"
		span.SetAttributes(
			attribute.String("db.query.text", query),
			attribute.String("db.query.parameter.id", *id))
		result, err = r.dB.Query(query, *id, *skip, *limit)
		if err != nil {
			return nil, err
		}
	} else {
		query := "SELECT id,tag_id,asset_id,value FROM tagvalue limit ?,?"
		span.SetAttributes(attribute.String("db.query.text", query))
		result, err = r.dB.Query(query, *skip, *limit)
		if err != nil {
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
		tagValue.Value = tagValueTemp.Value

		if err != nil {
			return nil, err
		}
		tagValues = append(tagValues, &tagValue)
	}

	span.SetAttributes(attribute.Int("tagValues.count", len(tagValues)))
	return tagValues, nil
}

// Search is the resolver for the search field.
func (r *queryResolver) Search(ctx context.Context, input model.Search, skip *int, limit *int) (*model.SearchResult, error) {
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(
		attribute.String("query", input.Text),
		attribute.Bool("searchAssetName", input.SearchAssetName),
		attribute.Bool("searchTagName", input.SearchTagName),
		attribute.Bool("searchTagValue", input.SearchTagValue),
		attribute.Int("skip", *skip),
		attribute.Int("limit", *limit))

	slog.Info("Search",
		"text", input.Text,
		"searchAssetName", input.SearchAssetName,
		"searchTagName", input.SearchTagName,
		"searchTagValue", input.SearchTagValue)
	searchResult := &model.SearchResult{}
	if input.SearchAssetName {
		assets, err := r.searchAssetName(ctx, input.Text, *skip, *limit)
		if err != nil {
			return nil, err
		}
		searchResult.Asset = assets
	}
	if input.SearchTagName {
		tags, err := r.searchTagName(ctx, input.Text, *skip, *limit)
		if err != nil {
			return nil, err
		}
		searchResult.Tag = tags
	}
	if input.SearchTagValue {
		tagValues, err := r.searchTagValue(ctx, input.Text, *skip, *limit)
		if err != nil {
			return nil, err
		}
		searchResult.TagValue = tagValues
	}

	return searchResult, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id *string, skip *int, limit *int) ([]*model.User, error) {
	slog.Info("User", "id", id, "skip", skip, "limit", limit)
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
		result, err = r.dB.Query(query, *id)
		if err != nil {
			return nil, err
		}
	} else {
		query := "SELECT id,email FROM user LIMIT ?,?"
		span.SetAttributes(attribute.String("db.query.text", query))
		result, err = r.dB.Query(query, *skip, *limit)
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
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

// Group is the resolver for the group field.
func (r *queryResolver) Group(ctx context.Context, id *string, skip *int, limit *int) ([]*model.Group, error) {
	slog.Info("Group", "id", id, "skip", skip, "limit", limit)
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
		result, err = r.dB.Query(query, *id)
		if err != nil {
			return nil, err
		}
	} else {
		query := "SELECT id,name FROM `group` LIMIT ?,?"
		span.SetAttributes(attribute.String("db.query.text", query))
		result, err = r.dB.Query(query, *skip, *limit)
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

	result, err := r.dB.Query(query, *skip, *limit)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	identities := []model.Identity{}
	for result.Next() {
		var user model.User
		err := result.Scan(&user.ID, &user.Email)
		if err != nil {
			return nil, err
		}
		identities = append(identities, user)
	}

	query = "SELECT id,name FROM `group` LIMIT ?,?"
	span.SetAttributes(attribute.String("db.query.text", query))
	result, err = r.dB.Query(query, *skip, *limit)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	for result.Next() {
		var group model.Group
		err := result.Scan(&group.ID, &group.Name)
		if err != nil {
			return nil, err
		}
		identities = append(identities, group)
	}

	return identities, nil
}

// Assets is the resolver for the assets field.
func (r *tagResolver) Assets(ctx context.Context, obj *model.Tag, skip *int, limit *int) ([]*model.Asset, error) {
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(
		attribute.String("id", obj.ID),
		attribute.Int("limit", *limit),
		attribute.Int("skip", *skip))
	slog.DebugContext(ctx, "Assets(byTag)", "id", obj.ID, "skip", skip, "limit", limit)
	span.SetAttributes(
		attribute.String("db.system", "mysql"),
		attribute.String("db.operation.name", "select"))

	query := "SELECT DISTINCT asset.id,asset.name FROM asset,tagvalue WHERE asset.id = tagvalue.asset_id and tagvalue.tag_id = ? limit ?,?"
	span.SetAttributes(
		attribute.String("db.query.text", query),
		attribute.String("db.parameter.id", obj.ID))
	result, err := r.dB.Query(query, obj.ID, skip, limit)
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
		assets = append(assets, &asset)
	}
	span.SetAttributes(attribute.Int("assets.count", len(assets)))
	return assets, nil
}

// Tag is the resolver for the tag field.
func (r *tagValueResolver) Tag(ctx context.Context, obj *model.TagValue) (*model.Tag, error) {
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(
		attribute.String("db.system", "mysql"),
		attribute.String("db.operation.name", "select"),
		attribute.String("id", obj.ID))
	slog.DebugContext(ctx, "Tag(byTagValue)", "id", obj.ID)

	query := "SELECT tag.id,tag.name FROM tag,tagvalue where tagvalue.tag_id = tag.id and tagvalue.id = ?"
	span.SetAttributes(
		attribute.String("db.query.text", query),
		attribute.String("db.parameter.id", obj.ID))
	result, err := r.dB.Query(query, obj.ID)
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

// Asset is the resolver for the asset field.
func (r *tagValueResolver) Asset(ctx context.Context, obj *model.TagValue) (*model.Asset, error) {
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("id", obj.ID))
	slog.DebugContext(ctx, "Asset(byTagValue)", "id", obj.ID)
	span.SetAttributes(
		attribute.String("db.system", "mysql"),
		attribute.String("db.operation.name", "select"))

	query := "SELECT asset.id,asset.name FROM asset,tagvalue where tagvalue.asset_id = asset.id and tagvalue.id = ?"
	span.SetAttributes(
		attribute.String("db.query.text", query),
		attribute.String("db.parameter.id", obj.ID))
	result, err := r.dB.Query(query, obj.ID)
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

// Groups is the resolver for the groups field.
func (r *userResolver) Groups(ctx context.Context, obj *model.User, skip *int, limit *int) ([]*model.Group, error) {
	slog.Info("Groups(forUser)", "id", obj.ID, "skip", skip, "limit", limit)
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
	result, err := r.dB.Query(query, obj.ID, skip, limit)
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

// Asset returns AssetResolver implementation.
func (r *Resolver) Asset() AssetResolver { return &assetResolver{r} }

// Group returns GroupResolver implementation.
func (r *Resolver) Group() GroupResolver { return &groupResolver{r} }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Tag returns TagResolver implementation.
func (r *Resolver) Tag() TagResolver { return &tagResolver{r} }

// TagValue returns TagValueResolver implementation.
func (r *Resolver) TagValue() TagValueResolver { return &tagValueResolver{r} }

// User returns UserResolver implementation.
func (r *Resolver) User() UserResolver { return &userResolver{r} }

type assetResolver struct{ *Resolver }
type groupResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type tagResolver struct{ *Resolver }
type tagValueResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
