// Code generated by badaas-cli v0.0.0, DO NOT EDIT.
package conditions

import (
	customtype "github.com/ditrit/badaas-cli/cmd/gen/conditions/tests/customtype"
	orm "github.com/ditrit/badaas/orm"
	"time"
)

func CustomTypeId(operator orm.Operator[orm.UUID]) orm.WhereCondition[customtype.CustomType] {
	return orm.FieldCondition[customtype.CustomType, orm.UUID]{
		Field:    "ID",
		Operator: operator,
	}
}
func CustomTypeCreatedAt(operator orm.Operator[time.Time]) orm.WhereCondition[customtype.CustomType] {
	return orm.FieldCondition[customtype.CustomType, time.Time]{
		Field:    "CreatedAt",
		Operator: operator,
	}
}
func CustomTypeUpdatedAt(operator orm.Operator[time.Time]) orm.WhereCondition[customtype.CustomType] {
	return orm.FieldCondition[customtype.CustomType, time.Time]{
		Field:    "UpdatedAt",
		Operator: operator,
	}
}
func CustomTypeDeletedAt(operator orm.Operator[time.Time]) orm.WhereCondition[customtype.CustomType] {
	return orm.FieldCondition[customtype.CustomType, time.Time]{
		Field:    "DeletedAt",
		Operator: operator,
	}
}
func CustomTypeCustom(operator orm.Operator[customtype.MultiString]) orm.WhereCondition[customtype.CustomType] {
	return orm.FieldCondition[customtype.CustomType, customtype.MultiString]{
		Field:    "Custom",
		Operator: operator,
	}
}
