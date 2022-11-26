// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package group

const (
	// Label holds the string label denoting the group type in the database.
	Label = "group"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeUsers holds the string denoting the users edge name in mutations.
	EdgeUsers = "users"
	// EdgeTags holds the string denoting the tags edge name in mutations.
	EdgeTags = "tags"
	// EdgeJoinedUsers holds the string denoting the joined_users edge name in mutations.
	EdgeJoinedUsers = "joined_users"
	// EdgeGroupTags holds the string denoting the group_tags edge name in mutations.
	EdgeGroupTags = "group_tags"
	// Table holds the table name of the group in the database.
	Table = "groups"
	// UsersTable is the table that holds the users relation/edge. The primary key declared below.
	UsersTable = "user_groups"
	// UsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UsersInverseTable = "users"
	// TagsTable is the table that holds the tags relation/edge. The primary key declared below.
	TagsTable = "group_tags"
	// TagsInverseTable is the table name for the Tag entity.
	// It exists in this package in order to avoid circular dependency with the "tag" package.
	TagsInverseTable = "tags"
	// JoinedUsersTable is the table that holds the joined_users relation/edge.
	JoinedUsersTable = "user_groups"
	// JoinedUsersInverseTable is the table name for the UserGroup entity.
	// It exists in this package in order to avoid circular dependency with the "usergroup" package.
	JoinedUsersInverseTable = "user_groups"
	// JoinedUsersColumn is the table column denoting the joined_users relation/edge.
	JoinedUsersColumn = "group_id"
	// GroupTagsTable is the table that holds the group_tags relation/edge.
	GroupTagsTable = "group_tags"
	// GroupTagsInverseTable is the table name for the GroupTag entity.
	// It exists in this package in order to avoid circular dependency with the "grouptag" package.
	GroupTagsInverseTable = "group_tags"
	// GroupTagsColumn is the table column denoting the group_tags relation/edge.
	GroupTagsColumn = "group_id"
)

// Columns holds all SQL columns for group fields.
var Columns = []string{
	FieldID,
	FieldName,
}

var (
	// UsersPrimaryKey and UsersColumn2 are the table columns denoting the
	// primary key for the users relation (M2M).
	UsersPrimaryKey = []string{"user_id", "group_id"}
	// TagsPrimaryKey and TagsColumn2 are the table columns denoting the
	// primary key for the tags relation (M2M).
	TagsPrimaryKey = []string{"tag_id", "group_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultName holds the default value on creation for the "name" field.
	DefaultName string
)
