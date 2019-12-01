package validators_test

import (
	"log"
	"testing"

	"github.com/gobuffalo/validate"
	"github.com/stretchr/testify/require"

	"github.com/gofrs/uuid"

	"habits/validators"
)

type MockPopConnection struct{
	exists bool
	err error
}

func (c *MockPopConnection) Where(string, ...interface{}) validators.PopQuery {
	return &MockPopQuery{exists: c.exists, err: c.err}
}

type MockPopQuery struct{
	exists bool
	err error
}

func (q *MockPopQuery) Exists(model interface{}) (bool, error) {
	return q.exists, q.err
}

func Test_IsValid(t *testing.T) {
	r := require.New(t)

	mtx := &MockPopConnection{exists: true, err: nil}

	v := validators.FieldIsUnique{Table: "users", Field: "Name", Value: "potter", ID: uuid.Nil, TX: mtx}
	errors := validate.NewErrors()
	v.IsValid(errors)
	r.Equal(true, errors.HasAny())
	r.Equal(errors.Get("name"), []string{"potter is already in use"})
}

func Test_uniqueQuery(t *testing.T) {
	r := require.New(t)

	uid, err := uuid.NewV4()
	if err != nil {
		log.Fatalf(err.Error())
	}

	var tests = []struct {
		value          string
		id             uuid.UUID
		desiredSQL     string
		shouldReturnID bool
	}{
		{"exotic-ladybug", uuid.Nil, "name = ?", false},
		{"common-snowflake", uid, "name = ? and id != ?", true},
	}

	for _, test := range tests {
		v := &validators.FieldIsUnique{Table: "users", Field: "name", Value: test.value, ID: test.id}
		sql, values := v.UniqueQuery()
		r.Equal(test.desiredSQL, sql)
		if test.shouldReturnID {
			r.Len(values, 2)
			r.Equal([]interface{}{test.value, test.id}, values)
		} else {
			r.Len(values, 1)
			r.Equal([]interface{}{test.value}, values)
		}
	}
}
