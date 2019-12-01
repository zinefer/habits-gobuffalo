package models

import (
	"time"

	"encoding/json"
	
	"github.com/gobuffalo/nulls"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofrs/uuid"

	_validators "habits/validators"
)

// User model for our users
type User struct {
	ID         uuid.UUID    `json:"id" db:"id"`
	Name       string       `json:"name" db:"name"`
	Nickname   string       `json:"nickname" db:"nickname"`
	Email      nulls.String `json:"email" db:"email"`
	Provider   string       `json:"provider" db:"provider"`
	ProviderID string       `json:"provider_id" db:"provider_id"`
	CreatedAt  time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time    `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (u User) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// Users is not required by pop and may be deleted
type Users []User

// String is not required by pop and may be deleted
func (u Users) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// HasUniqueNickname returns true if this users nickname is unique
func (u *User) HasUniqueNickname(tx *pop.Connection) bool {
	validator := getNicknameValidator("users", "Nickname", u.Nickname, u.ID, tx)
	exists := validate.Validate(validator).HasAny()
	return !exists
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
func (u *User) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		getNicknameValidator("users", "Nickname", u.Nickname, u.ID, tx),
		&validators.StringIsPresent{Field: u.Nickname, Name: "Nickname"},
		&validators.StringIsPresent{Field: u.Name, Name: "Name"},
		&validators.StringIsPresent{Field: u.Provider, Name: "Provider"},
		&validators.StringIsPresent{Field: u.ProviderID, Name: "ProviderID"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
func (u *User) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
func (u *User) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

func getNicknameValidator(table string, field string, value string, id uuid.UUID, tx *pop.Connection) validate.Validator {
	return &_validators.FieldIsUnique{Table: table, Field: field, Value: value, ID: id, TX: tx}
}
