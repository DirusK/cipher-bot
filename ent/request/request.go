// Code generated by entc, DO NOT EDIT.

package request

import (
	"fmt"
	"time"
)

const (
	// Label holds the string label denoting the request type in the database.
	Label = "request"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldAlgorithm holds the string denoting the algorithm field in the database.
	FieldAlgorithm = "algorithm"
	// FieldKeyMode holds the string denoting the key_mode field in the database.
	FieldKeyMode = "key_mode"
	// FieldManualKeyValidation holds the string denoting the manual_key_validation field in the database.
	FieldManualKeyValidation = "manual_key_validation"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the request in the database.
	Table = "requests"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "requests"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
)

// Columns holds all SQL columns for request fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldStatus,
	FieldType,
	FieldAlgorithm,
	FieldKeyMode,
	FieldManualKeyValidation,
	FieldUserID,
}

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
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)

// Status defines the type for the "status" enum field.
type Status string

// StatusActive is the default value of the Status enum.
const DefaultStatus = StatusActive

// Status values.
const (
	StatusActive  Status = "active"
	StatusDone    Status = "done"
	StatusExpired Status = "expired"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusActive, StatusDone, StatusExpired:
		return nil
	default:
		return fmt.Errorf("request: invalid enum value for status field: %q", s)
	}
}

// Type defines the type for the "type" enum field.
type Type string

// Type values.
const (
	TypeEncryption Type = "encryption"
	TypeDecryption Type = "decryption"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeEncryption, TypeDecryption:
		return nil
	default:
		return fmt.Errorf("request: invalid enum value for type field: %q", _type)
	}
}

// Algorithm defines the type for the "algorithm" enum field.
type Algorithm string

// Algorithm values.
const (
	AlgorithmAES Algorithm = "AES"
	AlgorithmRC4 Algorithm = "RC4"
)

func (a Algorithm) String() string {
	return string(a)
}

// AlgorithmValidator is a validator for the "algorithm" field enum values. It is called by the builders before save.
func AlgorithmValidator(a Algorithm) error {
	switch a {
	case AlgorithmAES, AlgorithmRC4:
		return nil
	default:
		return fmt.Errorf("request: invalid enum value for algorithm field: %q", a)
	}
}

// KeyMode defines the type for the "key_mode" enum field.
type KeyMode string

// KeyMode values.
const (
	KeyModeAuto   KeyMode = "auto"
	KeyModeManual KeyMode = "manual"
)

func (km KeyMode) String() string {
	return string(km)
}

// KeyModeValidator is a validator for the "key_mode" field enum values. It is called by the builders before save.
func KeyModeValidator(km KeyMode) error {
	switch km {
	case KeyModeAuto, KeyModeManual:
		return nil
	default:
		return fmt.Errorf("request: invalid enum value for key_mode field: %q", km)
	}
}
