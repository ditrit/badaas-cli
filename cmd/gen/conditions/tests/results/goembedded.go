// Code generated by badaas-cli v0.0.0, DO NOT EDIT.
package conditions

import (
	goembedded "github.com/ditrit/badaas-cli/cmd/gen/conditions/tests/goembedded"
	orm "github.com/ditrit/badaas/orm"
	"reflect"
	"time"
)

var goEmbeddedType = reflect.TypeOf(*new(goembedded.GoEmbedded))
var GoEmbeddedIdField = orm.FieldIdentifier[orm.UIntID]{
	Field:     "ID",
	ModelType: goEmbeddedType,
}

func GoEmbeddedId(operator orm.Operator[orm.UIntID]) orm.WhereCondition[goembedded.GoEmbedded] {
	return orm.FieldCondition[goembedded.GoEmbedded, orm.UIntID]{
		FieldIdentifier: GoEmbeddedIdField,
		Operator:        operator,
	}
}

var GoEmbeddedCreatedAtField = orm.FieldIdentifier[time.Time]{
	Field:     "CreatedAt",
	ModelType: goEmbeddedType,
}

func GoEmbeddedCreatedAt(operator orm.Operator[time.Time]) orm.WhereCondition[goembedded.GoEmbedded] {
	return orm.FieldCondition[goembedded.GoEmbedded, time.Time]{
		FieldIdentifier: GoEmbeddedCreatedAtField,
		Operator:        operator,
	}
}

var GoEmbeddedUpdatedAtField = orm.FieldIdentifier[time.Time]{
	Field:     "UpdatedAt",
	ModelType: goEmbeddedType,
}

func GoEmbeddedUpdatedAt(operator orm.Operator[time.Time]) orm.WhereCondition[goembedded.GoEmbedded] {
	return orm.FieldCondition[goembedded.GoEmbedded, time.Time]{
		FieldIdentifier: GoEmbeddedUpdatedAtField,
		Operator:        operator,
	}
}

var GoEmbeddedDeletedAtField = orm.FieldIdentifier[time.Time]{
	Field:     "DeletedAt",
	ModelType: goEmbeddedType,
}

func GoEmbeddedDeletedAt(operator orm.Operator[time.Time]) orm.WhereCondition[goembedded.GoEmbedded] {
	return orm.FieldCondition[goembedded.GoEmbedded, time.Time]{
		FieldIdentifier: GoEmbeddedDeletedAtField,
		Operator:        operator,
	}
}

var GoEmbeddedIntField = orm.FieldIdentifier[int]{
	Field:     "Int",
	ModelType: goEmbeddedType,
}

func GoEmbeddedInt(operator orm.Operator[int]) orm.WhereCondition[goembedded.GoEmbedded] {
	return orm.FieldCondition[goembedded.GoEmbedded, int]{
		FieldIdentifier: GoEmbeddedIntField,
		Operator:        operator,
	}
}

var GoEmbeddedToBeEmbeddedIntField = orm.FieldIdentifier[int]{
	Field:     "Int",
	ModelType: goEmbeddedType,
}

func GoEmbeddedToBeEmbeddedInt(operator orm.Operator[int]) orm.WhereCondition[goembedded.GoEmbedded] {
	return orm.FieldCondition[goembedded.GoEmbedded, int]{
		FieldIdentifier: GoEmbeddedToBeEmbeddedIntField,
		Operator:        operator,
	}
}

var GoEmbeddedPreloadAttributes = orm.NewPreloadCondition[goembedded.GoEmbedded](GoEmbeddedIdField, GoEmbeddedCreatedAtField, GoEmbeddedUpdatedAtField, GoEmbeddedDeletedAtField, GoEmbeddedIntField, GoEmbeddedToBeEmbeddedIntField)
