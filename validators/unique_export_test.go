package validators

type PopConnection = popConnection
type PopQuery = popQuery

func (v *FieldIsUnique) UniqueQuery() (string, []interface{}) {
	return v.uniqueQuery()
} 