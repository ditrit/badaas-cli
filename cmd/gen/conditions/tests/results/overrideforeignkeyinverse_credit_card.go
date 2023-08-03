// Code generated by badaas-cli v0.0.0, DO NOT EDIT.
package conditions

import (
	overrideforeignkeyinverse "github.com/ditrit/badaas-cli/cmd/gen/conditions/tests/overrideforeignkeyinverse"
	orm "github.com/ditrit/badaas/orm"
	"time"
)

func CreditCardId(operator orm.Operator[orm.UUID]) orm.WhereCondition[overrideforeignkeyinverse.CreditCard] {
	return orm.FieldCondition[overrideforeignkeyinverse.CreditCard, orm.UUID]{
		FieldIdentifier: orm.IDFieldID,
		Operator:        operator,
	}
}
func CreditCardCreatedAt(operator orm.Operator[time.Time]) orm.WhereCondition[overrideforeignkeyinverse.CreditCard] {
	return orm.FieldCondition[overrideforeignkeyinverse.CreditCard, time.Time]{
		FieldIdentifier: orm.CreatedAtFieldID,
		Operator:        operator,
	}
}
func CreditCardUpdatedAt(operator orm.Operator[time.Time]) orm.WhereCondition[overrideforeignkeyinverse.CreditCard] {
	return orm.FieldCondition[overrideforeignkeyinverse.CreditCard, time.Time]{
		FieldIdentifier: orm.UpdatedAtFieldID,
		Operator:        operator,
	}
}
func CreditCardDeletedAt(operator orm.Operator[time.Time]) orm.WhereCondition[overrideforeignkeyinverse.CreditCard] {
	return orm.FieldCondition[overrideforeignkeyinverse.CreditCard, time.Time]{
		FieldIdentifier: orm.DeletedAtFieldID,
		Operator:        operator,
	}
}

var creditCardUserReferenceFieldID = orm.FieldIdentifier{Field: "UserReference"}

func CreditCardUserReference(operator orm.Operator[orm.UUID]) orm.WhereCondition[overrideforeignkeyinverse.CreditCard] {
	return orm.FieldCondition[overrideforeignkeyinverse.CreditCard, orm.UUID]{
		FieldIdentifier: creditCardUserReferenceFieldID,
		Operator:        operator,
	}
}
