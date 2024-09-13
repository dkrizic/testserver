// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Identity interface {
	IsIdentity()
}

type Tag interface {
	IsTag()
	GetID() string
}

type TagCategory interface {
	IsTagCategory()
	GetID() string
	GetName() string
	GetParentTagCategory() TagCategory
	GetChildTagCategories() []TagCategory
	GetRootTags() []Tag
}

type Access struct {
	Identity   Identity   `json:"identity"`
	Permission Permission `json:"permission"`
	Asset      *Asset     `json:"asset"`
}

type Asset struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Permissions []*Access `json:"permissions,omitempty"`
	Files       []*File   `json:"files,omitempty"`
	Tags        []Tag     `json:"tags,omitempty"`
}

type DynamicTag struct {
	ID          string              `json:"id"`
	TagCategory *DynamicTagCategory `json:"tagCategory"`
	Value       string              `json:"value"`
}

func (DynamicTag) IsTag()             {}
func (this DynamicTag) GetID() string { return this.ID }

type DynamicTagCategory struct {
	ID                 string        `json:"id"`
	Name               string        `json:"name"`
	ParentTagCategory  TagCategory   `json:"parentTagCategory,omitempty"`
	ChildTagCategories []TagCategory `json:"childTagCategories,omitempty"`
	DynamicTags        []*DynamicTag `json:"dynamicTags,omitempty"`
	Format             *string       `json:"format,omitempty"`
	RootTags           []Tag         `json:"rootTags,omitempty"`
}

func (DynamicTagCategory) IsTagCategory()                         {}
func (this DynamicTagCategory) GetID() string                     { return this.ID }
func (this DynamicTagCategory) GetName() string                   { return this.Name }
func (this DynamicTagCategory) GetParentTagCategory() TagCategory { return this.ParentTagCategory }
func (this DynamicTagCategory) GetChildTagCategories() []TagCategory {
	if this.ChildTagCategories == nil {
		return nil
	}
	interfaceSlice := make([]TagCategory, 0, len(this.ChildTagCategories))
	for _, concrete := range this.ChildTagCategories {
		interfaceSlice = append(interfaceSlice, concrete)
	}
	return interfaceSlice
}
func (this DynamicTagCategory) GetRootTags() []Tag {
	if this.RootTags == nil {
		return nil
	}
	interfaceSlice := make([]Tag, 0, len(this.RootTags))
	for _, concrete := range this.RootTags {
		interfaceSlice = append(interfaceSlice, concrete)
	}
	return interfaceSlice
}

type File struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Asset *Asset `json:"asset"`
}

type Group struct {
	ID     string   `json:"id"`
	Name   string   `json:"name"`
	Users  []*User  `json:"users,omitempty"`
	Assets []*Asset `json:"assets,omitempty"`
}

func (Group) IsIdentity() {}

type Query struct {
}

type StaticTag struct {
	ID              string             `json:"id"`
	Name            string             `json:"name"`
	TagCategory     *StaticTagCategory `json:"tagCategory"`
	ParentStaticTag *StaticTag         `json:"parentStaticTag,omitempty"`
}

func (StaticTag) IsTag()             {}
func (this StaticTag) GetID() string { return this.ID }

type StaticTagCategory struct {
	ID                 string        `json:"id"`
	Name               string        `json:"name"`
	ParentTagCategory  TagCategory   `json:"parentTagCategory,omitempty"`
	ChildTagCategories []TagCategory `json:"childTagCategories,omitempty"`
	StaticTags         []*StaticTag  `json:"staticTags,omitempty"`
	IsOpen             bool          `json:"isOpen"`
	RootTags           []Tag         `json:"rootTags,omitempty"`
}

func (StaticTagCategory) IsTagCategory()                         {}
func (this StaticTagCategory) GetID() string                     { return this.ID }
func (this StaticTagCategory) GetName() string                   { return this.Name }
func (this StaticTagCategory) GetParentTagCategory() TagCategory { return this.ParentTagCategory }
func (this StaticTagCategory) GetChildTagCategories() []TagCategory {
	if this.ChildTagCategories == nil {
		return nil
	}
	interfaceSlice := make([]TagCategory, 0, len(this.ChildTagCategories))
	for _, concrete := range this.ChildTagCategories {
		interfaceSlice = append(interfaceSlice, concrete)
	}
	return interfaceSlice
}
func (this StaticTagCategory) GetRootTags() []Tag {
	if this.RootTags == nil {
		return nil
	}
	interfaceSlice := make([]Tag, 0, len(this.RootTags))
	for _, concrete := range this.RootTags {
		interfaceSlice = append(interfaceSlice, concrete)
	}
	return interfaceSlice
}

type User struct {
	ID     string   `json:"id"`
	Email  string   `json:"email"`
	Groups []*Group `json:"groups,omitempty"`
	Assets []*Asset `json:"assets,omitempty"`
}

func (User) IsIdentity() {}

type Permission string

const (
	PermissionRead  Permission = "READ"
	PermissionWrite Permission = "WRITE"
)

var AllPermission = []Permission{
	PermissionRead,
	PermissionWrite,
}

func (e Permission) IsValid() bool {
	switch e {
	case PermissionRead, PermissionWrite:
		return true
	}
	return false
}

func (e Permission) String() string {
	return string(e)
}

func (e *Permission) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Permission(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Permission", str)
	}
	return nil
}

func (e Permission) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type PermissionQuery string

const (
	PermissionQueryRead        PermissionQuery = "READ"
	PermissionQueryWrite       PermissionQuery = "WRITE"
	PermissionQueryReadOrWrite PermissionQuery = "READ_OR_WRITE"
)

var AllPermissionQuery = []PermissionQuery{
	PermissionQueryRead,
	PermissionQueryWrite,
	PermissionQueryReadOrWrite,
}

func (e PermissionQuery) IsValid() bool {
	switch e {
	case PermissionQueryRead, PermissionQueryWrite, PermissionQueryReadOrWrite:
		return true
	}
	return false
}

func (e PermissionQuery) String() string {
	return string(e)
}

func (e *PermissionQuery) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PermissionQuery(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PermissionQuery", str)
	}
	return nil
}

func (e PermissionQuery) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
