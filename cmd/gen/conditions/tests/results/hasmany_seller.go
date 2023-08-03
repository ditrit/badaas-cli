// Code generated by badaas-cli v0.0.0, DO NOT EDIT.
package conditions

import (
	hasmany "github.com/ditrit/badaas-cli/cmd/gen/conditions/tests/hasmany"
	orm "github.com/ditrit/badaas/orm"
	"time"
)

func SellerId(operator orm.Operator[orm.UUID]) orm.WhereCondition[hasmany.Seller] {
	return orm.FieldCondition[hasmany.Seller, orm.UUID]{
		FieldIdentifier: orm.IDFieldID,
		Operator:        operator,
	}
}
func SellerCreatedAt(operator orm.Operator[time.Time]) orm.WhereCondition[hasmany.Seller] {
	return orm.FieldCondition[hasmany.Seller, time.Time]{
		FieldIdentifier: orm.CreatedAtFieldID,
		Operator:        operator,
	}
}
func SellerUpdatedAt(operator orm.Operator[time.Time]) orm.WhereCondition[hasmany.Seller] {
	return orm.FieldCondition[hasmany.Seller, time.Time]{
		FieldIdentifier: orm.UpdatedAtFieldID,
		Operator:        operator,
	}
}
func SellerDeletedAt(operator orm.Operator[time.Time]) orm.WhereCondition[hasmany.Seller] {
	return orm.FieldCondition[hasmany.Seller, time.Time]{
		FieldIdentifier: orm.DeletedAtFieldID,
		Operator:        operator,
	}
}
func SellerCompany(conditions ...orm.Condition[hasmany.Company]) orm.Condition[hasmany.Seller] {
	return orm.JoinCondition[hasmany.Seller, hasmany.Company]{
		Conditions:    conditions,
		RelationField: "Company",
		T1Field:       "CompanyID",
		T2Field:       "ID",
	}
}

var sellerCompanyIdFieldID = orm.FieldIdentifier{Field: "CompanyID"}

func SellerCompanyId(operator orm.Operator[orm.UUID]) orm.WhereCondition[hasmany.Seller] {
	return orm.FieldCondition[hasmany.Seller, orm.UUID]{
		FieldIdentifier: sellerCompanyIdFieldID,
		Operator:        operator,
	}
}

var SellerPreloadAttributes = orm.NewPreloadCondition[hasmany.Seller](sellerCompanyIdFieldID)
