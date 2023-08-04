// Code generated by badaas-cli v0.0.0, DO NOT EDIT.
package conditions

import (
	overrideforeignkey "github.com/ditrit/badaas-cli/cmd/gen/conditions/tests/overrideforeignkey"
	orm "github.com/ditrit/badaas/orm"
	"reflect"
	"time"
)

var personType = reflect.TypeOf(*new(overrideforeignkey.Person))
var PersonIdField = orm.FieldIdentifier[orm.UUID]{
	Field:     "ID",
	ModelType: personType,
}

func PersonId(operator orm.Operator[orm.UUID]) orm.WhereCondition[overrideforeignkey.Person] {
	return orm.FieldCondition[overrideforeignkey.Person, orm.UUID]{
		FieldIdentifier: PersonIdField,
		Operator:        operator,
	}
}

var PersonCreatedAtField = orm.FieldIdentifier[time.Time]{
	Field:     "CreatedAt",
	ModelType: personType,
}

func PersonCreatedAt(operator orm.Operator[time.Time]) orm.WhereCondition[overrideforeignkey.Person] {
	return orm.FieldCondition[overrideforeignkey.Person, time.Time]{
		FieldIdentifier: PersonCreatedAtField,
		Operator:        operator,
	}
}

var PersonUpdatedAtField = orm.FieldIdentifier[time.Time]{
	Field:     "UpdatedAt",
	ModelType: personType,
}

func PersonUpdatedAt(operator orm.Operator[time.Time]) orm.WhereCondition[overrideforeignkey.Person] {
	return orm.FieldCondition[overrideforeignkey.Person, time.Time]{
		FieldIdentifier: PersonUpdatedAtField,
		Operator:        operator,
	}
}

var PersonDeletedAtField = orm.FieldIdentifier[time.Time]{
	Field:     "DeletedAt",
	ModelType: personType,
}

func PersonDeletedAt(operator orm.Operator[time.Time]) orm.WhereCondition[overrideforeignkey.Person] {
	return orm.FieldCondition[overrideforeignkey.Person, time.Time]{
		FieldIdentifier: PersonDeletedAtField,
		Operator:        operator,
	}
}

var PersonPreloadAttributes = orm.NewPreloadCondition[overrideforeignkey.Person](PersonIdField, PersonCreatedAtField, PersonUpdatedAtField, PersonDeletedAtField)
