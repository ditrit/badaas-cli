// Code generated by badaas-cli v0.0.0, DO NOT EDIT.
package conditions

import (
	customtype "github.com/ditrit/badaas-cli/cmd/gen/conditions/tests/customtype"
	orm "github.com/ditrit/badaas/orm"
	"time"
)

func CustomTypeId(operator orm.Operator[orm.UUID]) orm.WhereCondition[customtype.CustomType] {
	return orm.FieldCondition[customtype.CustomType, orm.UUID]{
		FieldIdentifier: orm.IDFieldID,
		Operator:        operator,
	}
}
func CustomTypeCreatedAt(operator orm.Operator[time.Time]) orm.WhereCondition[customtype.CustomType] {
	return orm.FieldCondition[customtype.CustomType, time.Time]{
		FieldIdentifier: orm.CreatedAtFieldID,
		Operator:        operator,
	}
}
func CustomTypeUpdatedAt(operator orm.Operator[time.Time]) orm.WhereCondition[customtype.CustomType] {
	return orm.FieldCondition[customtype.CustomType, time.Time]{
		FieldIdentifier: orm.UpdatedAtFieldID,
		Operator:        operator,
	}
}
func CustomTypeDeletedAt(operator orm.Operator[time.Time]) orm.WhereCondition[customtype.CustomType] {
	return orm.FieldCondition[customtype.CustomType, time.Time]{
		FieldIdentifier: orm.DeletedAtFieldID,
		Operator:        operator,
	}
}

var customTypeCustomFieldID = orm.FieldIdentifier{Field: "Custom"}

func CustomTypeCustom(operator orm.Operator[customtype.MultiString]) orm.WhereCondition[customtype.CustomType] {
	return orm.FieldCondition[customtype.CustomType, customtype.MultiString]{
		FieldIdentifier: customTypeCustomFieldID,
		Operator:        operator,
	}
}

var CustomTypePreloadAttributes = orm.NewPreloadCondition[customtype.CustomType](customTypeCustomFieldID)
