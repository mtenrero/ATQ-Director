package uuid

import "github.com/satori/go.uuid"

// NewUUID Creates new UUID
func NewUUID() string {
	uniqueid, _ := uuid.NewV4()

	return uniqueid.String()
}

// AppendAlias append a random generated UUID to a given alias
func AppendAlias(alias string) string {
	return alias + "_" + NewUUID()
}

// AppendAliasString append a random generated UUID to a given alias and another given string
func AppendAliasString(alias, additional string) string {
	return alias + "_" + additional + "_" + NewUUID()
}
