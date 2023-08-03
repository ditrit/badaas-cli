package conditions

import (
	"go/types"

	"github.com/dave/jennifer/jen"
	"github.com/ettle/strcase"

	"github.com/ditrit/badaas-cli/cmd/log"
)

const (
	// badaas/orm/condition.go
	badaasORMCondition      = "Condition"
	badaasORMFieldCondition = "FieldCondition"
	badaasORMWhereCondition = "WhereCondition"
	badaasORMJoinCondition  = "JoinCondition"
	// badaas/orm/operator.go
	badaasORMOperator = "Operator"
)

type Condition struct {
	codes   []jen.Code
	param   *JenParam
	destPkg string
}

func NewCondition(destPkg string, objectType Type, field Field) *Condition {
	condition := &Condition{
		param:   NewJenParam(),
		destPkg: destPkg,
	}
	condition.generate(objectType, field)
	return condition
}

// Generate the condition between the object and the field
func (condition *Condition) generate(objectType Type, field Field) {
	switch fieldType := field.GetType().(type) {
	case *types.Basic:
		// the field is a basic type (string, int, etc)
		// adapt param to that type and generate a WhereCondition
		condition.param.ToBasicKind(fieldType)
		condition.generateWhere(
			objectType,
			field,
		)
	case *types.Named:
		// the field is a named type (user defined structs)
		condition.generateForNamedType(
			objectType,
			field,
		)
	case *types.Pointer:
		// the field is a pointer
		condition.generate(
			objectType,
			field.ChangeType(fieldType.Elem()),
		)
	case *types.Slice:
		// the field is a slice
		// adapt param to slice and
		// generate code for the type of the elements of the slice
		condition.param.ToSlice()
		condition.generateForSlice(
			objectType,
			field.ChangeType(fieldType.Elem()),
		)
	default:
		log.Logger.Debugf("struct field type not handled: %T", fieldType)
	}
}

// Generate condition between the object and the field when the field is a slice
func (condition *Condition) generateForSlice(objectType Type, field Field) {
	switch elemType := field.GetType().(type) {
	case *types.Basic:
		// slice of basic types ([]string, []int, etc.)
		// the only one supported directly by gorm is []byte
		// but the others can be used after configuration in some dbs
		condition.generate(
			objectType,
			field,
		)
	case *types.Pointer:
		// slice of pointers, generate code for a slice of the pointed type
		condition.generateForSlice(
			objectType,
			field.ChangeType(elemType.Elem()),
		)
	default:
		log.Logger.Debugf("struct field list elem type not handled: %T", elemType)
	}
}

// Generate condition between object and field when the field is a named type (user defined struct)
func (condition *Condition) generateForNamedType(objectType Type, field Field) {
	_, err := field.Type.BadaasModelStruct()

	if err == nil {
		// field is a badaas model
		condition.generateForBadaasModel(objectType, field)
	} else if field.Type.IsSQLNullableType() {
		// field is a sql nullable type (sql.NullBool, sql.NullInt, etc.)
		condition.param.SQLToBasicType(field.Type)
		condition.generateWhere(
			objectType,
			field,
		)
	} else if field.Type.IsGormCustomType() || field.TypeString() == "time.Time" {
		// field is a Gorm Custom type (implements Scanner and Valuer interfaces)
		// or a named type supported by gorm (time.Time)
		condition.param.ToCustomType(condition.destPkg, field.Type)
		condition.generateWhere(
			objectType,
			field,
		)
	} else {
		log.Logger.Debugf("struct field type not handled: %s", field.TypeString())
	}
}

// Generate condition between object and field when the field is a Badaas Model
func (condition *Condition) generateForBadaasModel(objectType Type, field Field) {
	hasFK, _ := objectType.HasFK(field)
	if hasFK {
		// belongsTo relation
		condition.generateJoinWithFK(
			objectType,
			field,
		)
	} else {
		// hasOne relation
		condition.generateJoinWithoutFK(
			objectType,
			field,
		)
	}
}

// Generate a WhereCondition between object and field
func (condition *Condition) generateWhere(objectType Type, field Field) {
	objectTypeQual := jen.Qual(
		getRelativePackagePath(condition.destPkg, objectType),
		objectType.Name(),
	)

	fieldCondition := jen.Qual(
		badaasORMPath, badaasORMFieldCondition,
	).Types(
		objectTypeQual,
		condition.param.GenericType(),
	)

	whereCondition := jen.Qual(
		badaasORMPath, badaasORMWhereCondition,
	).Types(
		objectTypeQual,
	)

	conditionName := getConditionName(objectType, field)
	log.Logger.Debugf("Generated %q", conditionName)

	conditionValues := jen.Dict{
		jen.Id("Operator"): jen.Id("operator"),
	}
	columnName := field.getColumnName()

	if columnName != "" {
		conditionValues[jen.Id("Column")] = jen.Lit(columnName)
	} else {
		conditionValues[jen.Id("Field")] = jen.Lit(field.Name)
	}

	columnPrefix := field.ColumnPrefix
	if columnPrefix != "" {
		conditionValues[jen.Id("ColumnPrefix")] = jen.Lit(columnPrefix)
	}

	condition.codes = append(
		condition.codes,
		jen.Func().Id(
			conditionName,
		).Params(
			condition.param.Statement(),
		).Add(
			whereCondition,
		).Block(
			jen.Return(
				fieldCondition.Clone().Values(conditionValues),
			),
		),
	)
}

// Generate a JoinCondition between the object and field's object
// when object has a foreign key to the field's object
func (condition *Condition) generateJoinWithFK(objectType Type, field Field) {
	condition.generateJoin(
		objectType,
		field,
		field.getFKAttribute(),
		field.getFKReferencesAttribute(),
	)
}

// Generate a JoinCondition between the object and field's object
// when object has not a foreign key to the field's object
// (so the field's object has it)
func (condition *Condition) generateJoinWithoutFK(objectType Type, field Field) {
	condition.generateJoin(
		objectType,
		field,
		field.getFKReferencesAttribute(),
		field.getRelatedTypeFKAttribute(objectType.Name()),
	)
}

// Generate a JoinCondition
func (condition *Condition) generateJoin(objectType Type, field Field, t1Field, t2Field string) {
	t1 := jen.Qual(
		getRelativePackagePath(condition.destPkg, objectType),
		objectType.Name(),
	)

	t2 := jen.Qual(
		getRelativePackagePath(condition.destPkg, field.Type),
		field.TypeName(),
	)

	conditionName := getConditionName(objectType, field)
	log.Logger.Debugf("Generated %q", conditionName)

	ormT1Condition := jen.Qual(
		badaasORMPath, badaasORMCondition,
	).Types(t1)
	ormT2Condition := jen.Qual(
		badaasORMPath, badaasORMCondition,
	).Types(t2)
	ormJoinCondition := jen.Qual(
		badaasORMPath, badaasORMJoinCondition,
	).Types(
		t1, t2,
	)

	condition.codes = append(
		condition.codes,
		jen.Func().Id(
			conditionName,
		).Params(
			jen.Id("conditions").Op("...").Add(ormT2Condition),
		).Add(
			ormT1Condition,
		).Block(
			jen.Return(
				ormJoinCondition.Values(jen.Dict{
					jen.Id("T1Field"):       jen.Lit(t1Field),
					jen.Id("T2Field"):       jen.Lit(t2Field),
					jen.Id("RelationField"): jen.Lit(field.Name),
					jen.Id("Conditions"):    jen.Id("conditions"),
				}),
			),
		),
	)
}

// Generate condition names
func getConditionName(typeV Type, field Field) string {
	return typeV.Name() + strcase.ToPascal(field.ColumnPrefix) + strcase.ToPascal(field.Name)
}

// Avoid importing the same package as the destination one
func getRelativePackagePath(destPkg string, typeV Type) string {
	srcPkg := typeV.Pkg()
	if srcPkg == nil || srcPkg.Name() == destPkg {
		return ""
	}

	return srcPkg.Path()
}
