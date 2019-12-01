package validators

import (
	"fmt"
	"log"
	"strings"

	"github.com/gobuffalo/validate"
	"github.com/gofrs/uuid"
)

// FieldIsUnique is a gobuffalo/validator that ensures a Value is unique 
// per Table.Field. ID is used to ignore a possible self.
type FieldIsUnique struct {
	Table string
	Field string
	Value string
	ID    uuid.UUID
	TX    popConnection
}

// IsValid queries the database and adds an error if Value is already in use
func (v *FieldIsUnique) IsValid(errors *validate.Errors) {
	v.Field = strings.ToLower(v.Field)
	unique, values := v.uniqueQuery()
	var q popQuery = v.TX.Where(unique, values...)
	exists, err := q.Exists(v.Table)
	if err != nil {
		log.Fatalf(err.Error())
	}
	if exists {
		errors.Add(v.Field, fmt.Sprintf("%s is already in use", v.Value))
	}
}

func (v *FieldIsUnique) uniqueQuery() (string, []interface{}) {
	fieldEq := fmt.Sprintf("%s = ?", v.Field)
	if v.ID == uuid.Nil {
		return fieldEq, []interface{}{v.Value}
	}
	return fieldEq + " and id != ?", []interface{}{v.Value, v.ID}
}

type popConnection interface {
	Where(string, ...interface{}) popQuery
}

type popQuery interface {
	Exists(model interface{}) (bool, error)
}