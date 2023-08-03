// Code generated by badaas-cli v0.0.0, DO NOT EDIT.
package conditions

import (
	hasone "github.com/ditrit/badaas-cli/cmd/gen/conditions/tests/hasone"
	orm "github.com/ditrit/badaas/orm"
	"time"
)

func CountryId(operator orm.Operator[orm.UUID]) orm.WhereCondition[hasone.Country] {
	return orm.FieldCondition[hasone.Country, orm.UUID]{
		Field:    "ID",
		Operator: operator,
	}
}
func CountryCreatedAt(operator orm.Operator[time.Time]) orm.WhereCondition[hasone.Country] {
	return orm.FieldCondition[hasone.Country, time.Time]{
		Field:    "CreatedAt",
		Operator: operator,
	}
}
func CountryUpdatedAt(operator orm.Operator[time.Time]) orm.WhereCondition[hasone.Country] {
	return orm.FieldCondition[hasone.Country, time.Time]{
		Field:    "UpdatedAt",
		Operator: operator,
	}
}
func CountryDeletedAt(operator orm.Operator[time.Time]) orm.WhereCondition[hasone.Country] {
	return orm.FieldCondition[hasone.Country, time.Time]{
		Field:    "DeletedAt",
		Operator: operator,
	}
}
func CountryCapital(conditions ...orm.Condition[hasone.City]) orm.Condition[hasone.Country] {
	return orm.JoinCondition[hasone.Country, hasone.City]{
		Conditions: conditions,
		T1Field:    "ID",
		T2Field:    "CountryID",
	}
}
