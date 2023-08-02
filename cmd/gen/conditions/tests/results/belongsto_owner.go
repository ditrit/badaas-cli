// Code generated by badaas-cli v0.0.0, DO NOT EDIT.
package conditions

import (
	belongsto "github.com/ditrit/badaas-cli/cmd/gen/conditions/tests/belongsto"
	orm "github.com/ditrit/badaas/orm"
	gorm "gorm.io/gorm"
	"time"
)

func OwnerId(operator orm.Operator[orm.UUID]) orm.WhereCondition[belongsto.Owner] {
	return orm.FieldCondition[belongsto.Owner, orm.UUID]{
		Field:    "ID",
		Operator: operator,
	}
}
func OwnerCreatedAt(operator orm.Operator[time.Time]) orm.WhereCondition[belongsto.Owner] {
	return orm.FieldCondition[belongsto.Owner, time.Time]{
		Field:    "CreatedAt",
		Operator: operator,
	}
}
func OwnerUpdatedAt(operator orm.Operator[time.Time]) orm.WhereCondition[belongsto.Owner] {
	return orm.FieldCondition[belongsto.Owner, time.Time]{
		Field:    "UpdatedAt",
		Operator: operator,
	}
}
func OwnerDeletedAt(operator orm.Operator[gorm.DeletedAt]) orm.WhereCondition[belongsto.Owner] {
	return orm.FieldCondition[belongsto.Owner, gorm.DeletedAt]{
		Field:    "DeletedAt",
		Operator: operator,
	}
}
