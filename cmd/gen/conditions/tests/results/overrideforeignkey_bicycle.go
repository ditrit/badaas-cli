// Code generated by badaas-cli v0.0.0, DO NOT EDIT.
package conditions

import (
	overrideforeignkey "github.com/ditrit/badaas-cli/cmd/gen/conditions/tests/overrideforeignkey"
	orm "github.com/ditrit/badaas/orm"
	"time"
)

func BicycleId(operator orm.Operator[orm.UUID]) orm.WhereCondition[overrideforeignkey.Bicycle] {
	return orm.FieldCondition[overrideforeignkey.Bicycle, orm.UUID]{
		FieldIdentifier: orm.IDFieldID,
		Operator:        operator,
	}
}
func BicycleCreatedAt(operator orm.Operator[time.Time]) orm.WhereCondition[overrideforeignkey.Bicycle] {
	return orm.FieldCondition[overrideforeignkey.Bicycle, time.Time]{
		FieldIdentifier: orm.CreatedAtFieldID,
		Operator:        operator,
	}
}
func BicycleUpdatedAt(operator orm.Operator[time.Time]) orm.WhereCondition[overrideforeignkey.Bicycle] {
	return orm.FieldCondition[overrideforeignkey.Bicycle, time.Time]{
		FieldIdentifier: orm.UpdatedAtFieldID,
		Operator:        operator,
	}
}
func BicycleDeletedAt(operator orm.Operator[time.Time]) orm.WhereCondition[overrideforeignkey.Bicycle] {
	return orm.FieldCondition[overrideforeignkey.Bicycle, time.Time]{
		FieldIdentifier: orm.DeletedAtFieldID,
		Operator:        operator,
	}
}
func BicycleOwner(conditions ...orm.Condition[overrideforeignkey.Person]) orm.IJoinCondition[overrideforeignkey.Bicycle] {
	return orm.JoinCondition[overrideforeignkey.Bicycle, overrideforeignkey.Person]{
		Conditions:         conditions,
		RelationField:      "Owner",
		T1Field:            "OwnerSomethingID",
		T1PreloadCondition: BicyclePreloadAttributes,
		T2Field:            "ID",
	}
}

var BicyclePreloadOwner = BicycleOwner(PersonPreloadAttributes)
var bicycleOwnerSomethingIdFieldID = orm.FieldIdentifier{Field: "OwnerSomethingID"}

func BicycleOwnerSomethingId(operator orm.Operator[string]) orm.WhereCondition[overrideforeignkey.Bicycle] {
	return orm.FieldCondition[overrideforeignkey.Bicycle, string]{
		FieldIdentifier: bicycleOwnerSomethingIdFieldID,
		Operator:        operator,
	}
}

var BicyclePreloadAttributes = orm.NewPreloadCondition[overrideforeignkey.Bicycle](bicycleOwnerSomethingIdFieldID)
var BicyclePreloadRelations = []orm.Condition[overrideforeignkey.Bicycle]{BicyclePreloadOwner}
