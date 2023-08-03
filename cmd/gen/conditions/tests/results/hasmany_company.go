// Code generated by badaas-cli v0.0.0, DO NOT EDIT.
package conditions

import (
	hasmany "github.com/ditrit/badaas-cli/cmd/gen/conditions/tests/hasmany"
	orm "github.com/ditrit/badaas/orm"
	"time"
)

func CompanyId(operator orm.Operator[orm.UUID]) orm.WhereCondition[hasmany.Company] {
	return orm.FieldCondition[hasmany.Company, orm.UUID]{
		Field:    "ID",
		Operator: operator,
	}
}
func CompanyCreatedAt(operator orm.Operator[time.Time]) orm.WhereCondition[hasmany.Company] {
	return orm.FieldCondition[hasmany.Company, time.Time]{
		Field:    "CreatedAt",
		Operator: operator,
	}
}
func CompanyUpdatedAt(operator orm.Operator[time.Time]) orm.WhereCondition[hasmany.Company] {
	return orm.FieldCondition[hasmany.Company, time.Time]{
		Field:    "UpdatedAt",
		Operator: operator,
	}
}
func CompanyDeletedAt(operator orm.Operator[time.Time]) orm.WhereCondition[hasmany.Company] {
	return orm.FieldCondition[hasmany.Company, time.Time]{
		Field:    "DeletedAt",
		Operator: operator,
	}
}
