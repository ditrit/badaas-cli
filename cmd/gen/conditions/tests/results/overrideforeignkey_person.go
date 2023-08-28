// Code generated by badaas-cli v0.0.0, DO NOT EDIT.
package conditions

import (
	overrideforeignkey "github.com/ditrit/badaas-cli/cmd/gen/conditions/tests/overrideforeignkey"
	orm "github.com/ditrit/badaas/orm"
	condition "github.com/ditrit/badaas/orm/condition"
	model "github.com/ditrit/badaas/orm/model"
	query "github.com/ditrit/badaas/orm/query"
	"reflect"
	"time"
)

var personType = reflect.TypeOf(*new(overrideforeignkey.Person))

func (personConditions personConditions) IdIs() orm.FieldIs[overrideforeignkey.Person, model.UUID] {
	return orm.FieldIs[overrideforeignkey.Person, model.UUID]{FieldID: personConditions.ID}
}
func (personConditions personConditions) CreatedAtIs() orm.FieldIs[overrideforeignkey.Person, time.Time] {
	return orm.FieldIs[overrideforeignkey.Person, time.Time]{FieldID: personConditions.CreatedAt}
}
func (personConditions personConditions) UpdatedAtIs() orm.FieldIs[overrideforeignkey.Person, time.Time] {
	return orm.FieldIs[overrideforeignkey.Person, time.Time]{FieldID: personConditions.UpdatedAt}
}
func (personConditions personConditions) DeletedAtIs() orm.FieldIs[overrideforeignkey.Person, time.Time] {
	return orm.FieldIs[overrideforeignkey.Person, time.Time]{FieldID: personConditions.DeletedAt}
}

type personConditions struct {
	ID        query.FieldIdentifier[model.UUID]
	CreatedAt query.FieldIdentifier[time.Time]
	UpdatedAt query.FieldIdentifier[time.Time]
	DeletedAt query.FieldIdentifier[time.Time]
}

var Person = personConditions{
	CreatedAt: query.FieldIdentifier[time.Time]{
		Field:     "CreatedAt",
		ModelType: personType,
	},
	DeletedAt: query.FieldIdentifier[time.Time]{
		Field:     "DeletedAt",
		ModelType: personType,
	},
	ID: query.FieldIdentifier[model.UUID]{
		Field:     "ID",
		ModelType: personType,
	},
	UpdatedAt: query.FieldIdentifier[time.Time]{
		Field:     "UpdatedAt",
		ModelType: personType,
	},
}

// Preload allows preloading the Person when doing a query
func (personConditions personConditions) Preload() condition.Condition[overrideforeignkey.Person] {
	return condition.NewPreloadCondition[overrideforeignkey.Person](personConditions.ID, personConditions.CreatedAt, personConditions.UpdatedAt, personConditions.DeletedAt)
}
