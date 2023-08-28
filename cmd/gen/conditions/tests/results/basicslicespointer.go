// Code generated by badaas-cli v0.0.0, DO NOT EDIT.
package conditions

import (
	basicslicespointer "github.com/ditrit/badaas-cli/cmd/gen/conditions/tests/basicslicespointer"
	orm "github.com/ditrit/badaas/orm"
	condition "github.com/ditrit/badaas/orm/condition"
	model "github.com/ditrit/badaas/orm/model"
	query "github.com/ditrit/badaas/orm/query"
	"reflect"
	"time"
)

var basicSlicesPointerType = reflect.TypeOf(*new(basicslicespointer.BasicSlicesPointer))

func (basicSlicesPointerConditions basicSlicesPointerConditions) IdIs() orm.FieldIs[basicslicespointer.BasicSlicesPointer, model.UUID] {
	return orm.FieldIs[basicslicespointer.BasicSlicesPointer, model.UUID]{FieldID: basicSlicesPointerConditions.ID}
}
func (basicSlicesPointerConditions basicSlicesPointerConditions) CreatedAtIs() orm.FieldIs[basicslicespointer.BasicSlicesPointer, time.Time] {
	return orm.FieldIs[basicslicespointer.BasicSlicesPointer, time.Time]{FieldID: basicSlicesPointerConditions.CreatedAt}
}
func (basicSlicesPointerConditions basicSlicesPointerConditions) UpdatedAtIs() orm.FieldIs[basicslicespointer.BasicSlicesPointer, time.Time] {
	return orm.FieldIs[basicslicespointer.BasicSlicesPointer, time.Time]{FieldID: basicSlicesPointerConditions.UpdatedAt}
}
func (basicSlicesPointerConditions basicSlicesPointerConditions) DeletedAtIs() orm.FieldIs[basicslicespointer.BasicSlicesPointer, time.Time] {
	return orm.FieldIs[basicslicespointer.BasicSlicesPointer, time.Time]{FieldID: basicSlicesPointerConditions.DeletedAt}
}
func (basicSlicesPointerConditions basicSlicesPointerConditions) BoolIs() orm.FieldIs[basicslicespointer.BasicSlicesPointer, []bool] {
	return orm.FieldIs[basicslicespointer.BasicSlicesPointer, []bool]{FieldID: basicSlicesPointerConditions.Bool}
}
func (basicSlicesPointerConditions basicSlicesPointerConditions) IntIs() orm.FieldIs[basicslicespointer.BasicSlicesPointer, []int] {
	return orm.FieldIs[basicslicespointer.BasicSlicesPointer, []int]{FieldID: basicSlicesPointerConditions.Int}
}
func (basicSlicesPointerConditions basicSlicesPointerConditions) Int8Is() orm.FieldIs[basicslicespointer.BasicSlicesPointer, []int8] {
	return orm.FieldIs[basicslicespointer.BasicSlicesPointer, []int8]{FieldID: basicSlicesPointerConditions.Int8}
}
func (basicSlicesPointerConditions basicSlicesPointerConditions) Int16Is() orm.FieldIs[basicslicespointer.BasicSlicesPointer, []int16] {
	return orm.FieldIs[basicslicespointer.BasicSlicesPointer, []int16]{FieldID: basicSlicesPointerConditions.Int16}
}
func (basicSlicesPointerConditions basicSlicesPointerConditions) Int32Is() orm.FieldIs[basicslicespointer.BasicSlicesPointer, []int32] {
	return orm.FieldIs[basicslicespointer.BasicSlicesPointer, []int32]{FieldID: basicSlicesPointerConditions.Int32}
}
func (basicSlicesPointerConditions basicSlicesPointerConditions) Int64Is() orm.FieldIs[basicslicespointer.BasicSlicesPointer, []int64] {
	return orm.FieldIs[basicslicespointer.BasicSlicesPointer, []int64]{FieldID: basicSlicesPointerConditions.Int64}
}
func (basicSlicesPointerConditions basicSlicesPointerConditions) UIntIs() orm.FieldIs[basicslicespointer.BasicSlicesPointer, []uint] {
	return orm.FieldIs[basicslicespointer.BasicSlicesPointer, []uint]{FieldID: basicSlicesPointerConditions.UInt}
}
func (basicSlicesPointerConditions basicSlicesPointerConditions) UInt8Is() orm.FieldIs[basicslicespointer.BasicSlicesPointer, []uint8] {
	return orm.FieldIs[basicslicespointer.BasicSlicesPointer, []uint8]{FieldID: basicSlicesPointerConditions.UInt8}
}
func (basicSlicesPointerConditions basicSlicesPointerConditions) UInt16Is() orm.FieldIs[basicslicespointer.BasicSlicesPointer, []uint16] {
	return orm.FieldIs[basicslicespointer.BasicSlicesPointer, []uint16]{FieldID: basicSlicesPointerConditions.UInt16}
}
func (basicSlicesPointerConditions basicSlicesPointerConditions) UInt32Is() orm.FieldIs[basicslicespointer.BasicSlicesPointer, []uint32] {
	return orm.FieldIs[basicslicespointer.BasicSlicesPointer, []uint32]{FieldID: basicSlicesPointerConditions.UInt32}
}
func (basicSlicesPointerConditions basicSlicesPointerConditions) UInt64Is() orm.FieldIs[basicslicespointer.BasicSlicesPointer, []uint64] {
	return orm.FieldIs[basicslicespointer.BasicSlicesPointer, []uint64]{FieldID: basicSlicesPointerConditions.UInt64}
}
func (basicSlicesPointerConditions basicSlicesPointerConditions) UIntptrIs() orm.FieldIs[basicslicespointer.BasicSlicesPointer, []uintptr] {
	return orm.FieldIs[basicslicespointer.BasicSlicesPointer, []uintptr]{FieldID: basicSlicesPointerConditions.UIntptr}
}
func (basicSlicesPointerConditions basicSlicesPointerConditions) Float32Is() orm.FieldIs[basicslicespointer.BasicSlicesPointer, []float32] {
	return orm.FieldIs[basicslicespointer.BasicSlicesPointer, []float32]{FieldID: basicSlicesPointerConditions.Float32}
}
func (basicSlicesPointerConditions basicSlicesPointerConditions) Float64Is() orm.FieldIs[basicslicespointer.BasicSlicesPointer, []float64] {
	return orm.FieldIs[basicslicespointer.BasicSlicesPointer, []float64]{FieldID: basicSlicesPointerConditions.Float64}
}
func (basicSlicesPointerConditions basicSlicesPointerConditions) Complex64Is() orm.FieldIs[basicslicespointer.BasicSlicesPointer, []complex64] {
	return orm.FieldIs[basicslicespointer.BasicSlicesPointer, []complex64]{FieldID: basicSlicesPointerConditions.Complex64}
}
func (basicSlicesPointerConditions basicSlicesPointerConditions) Complex128Is() orm.FieldIs[basicslicespointer.BasicSlicesPointer, []complex128] {
	return orm.FieldIs[basicslicespointer.BasicSlicesPointer, []complex128]{FieldID: basicSlicesPointerConditions.Complex128}
}
func (basicSlicesPointerConditions basicSlicesPointerConditions) StringIs() orm.FieldIs[basicslicespointer.BasicSlicesPointer, []string] {
	return orm.FieldIs[basicslicespointer.BasicSlicesPointer, []string]{FieldID: basicSlicesPointerConditions.String}
}
func (basicSlicesPointerConditions basicSlicesPointerConditions) ByteIs() orm.FieldIs[basicslicespointer.BasicSlicesPointer, []uint8] {
	return orm.FieldIs[basicslicespointer.BasicSlicesPointer, []uint8]{FieldID: basicSlicesPointerConditions.Byte}
}

type basicSlicesPointerConditions struct {
	ID         query.FieldIdentifier[model.UUID]
	CreatedAt  query.FieldIdentifier[time.Time]
	UpdatedAt  query.FieldIdentifier[time.Time]
	DeletedAt  query.FieldIdentifier[time.Time]
	Bool       query.FieldIdentifier[[]bool]
	Int        query.FieldIdentifier[[]int]
	Int8       query.FieldIdentifier[[]int8]
	Int16      query.FieldIdentifier[[]int16]
	Int32      query.FieldIdentifier[[]int32]
	Int64      query.FieldIdentifier[[]int64]
	UInt       query.FieldIdentifier[[]uint]
	UInt8      query.FieldIdentifier[[]uint8]
	UInt16     query.FieldIdentifier[[]uint16]
	UInt32     query.FieldIdentifier[[]uint32]
	UInt64     query.FieldIdentifier[[]uint64]
	UIntptr    query.FieldIdentifier[[]uintptr]
	Float32    query.FieldIdentifier[[]float32]
	Float64    query.FieldIdentifier[[]float64]
	Complex64  query.FieldIdentifier[[]complex64]
	Complex128 query.FieldIdentifier[[]complex128]
	String     query.FieldIdentifier[[]string]
	Byte       query.FieldIdentifier[[]uint8]
}

var BasicSlicesPointer = basicSlicesPointerConditions{
	Bool: query.FieldIdentifier[[]bool]{
		Field:     "Bool",
		ModelType: basicSlicesPointerType,
	},
	Byte: query.FieldIdentifier[[]uint8]{
		Field:     "Byte",
		ModelType: basicSlicesPointerType,
	},
	Complex128: query.FieldIdentifier[[]complex128]{
		Field:     "Complex128",
		ModelType: basicSlicesPointerType,
	},
	Complex64: query.FieldIdentifier[[]complex64]{
		Field:     "Complex64",
		ModelType: basicSlicesPointerType,
	},
	CreatedAt: query.FieldIdentifier[time.Time]{
		Field:     "CreatedAt",
		ModelType: basicSlicesPointerType,
	},
	DeletedAt: query.FieldIdentifier[time.Time]{
		Field:     "DeletedAt",
		ModelType: basicSlicesPointerType,
	},
	Float32: query.FieldIdentifier[[]float32]{
		Field:     "Float32",
		ModelType: basicSlicesPointerType,
	},
	Float64: query.FieldIdentifier[[]float64]{
		Field:     "Float64",
		ModelType: basicSlicesPointerType,
	},
	ID: query.FieldIdentifier[model.UUID]{
		Field:     "ID",
		ModelType: basicSlicesPointerType,
	},
	Int: query.FieldIdentifier[[]int]{
		Field:     "Int",
		ModelType: basicSlicesPointerType,
	},
	Int16: query.FieldIdentifier[[]int16]{
		Field:     "Int16",
		ModelType: basicSlicesPointerType,
	},
	Int32: query.FieldIdentifier[[]int32]{
		Field:     "Int32",
		ModelType: basicSlicesPointerType,
	},
	Int64: query.FieldIdentifier[[]int64]{
		Field:     "Int64",
		ModelType: basicSlicesPointerType,
	},
	Int8: query.FieldIdentifier[[]int8]{
		Field:     "Int8",
		ModelType: basicSlicesPointerType,
	},
	String: query.FieldIdentifier[[]string]{
		Field:     "String",
		ModelType: basicSlicesPointerType,
	},
	UInt: query.FieldIdentifier[[]uint]{
		Field:     "UInt",
		ModelType: basicSlicesPointerType,
	},
	UInt16: query.FieldIdentifier[[]uint16]{
		Field:     "UInt16",
		ModelType: basicSlicesPointerType,
	},
	UInt32: query.FieldIdentifier[[]uint32]{
		Field:     "UInt32",
		ModelType: basicSlicesPointerType,
	},
	UInt64: query.FieldIdentifier[[]uint64]{
		Field:     "UInt64",
		ModelType: basicSlicesPointerType,
	},
	UInt8: query.FieldIdentifier[[]uint8]{
		Field:     "UInt8",
		ModelType: basicSlicesPointerType,
	},
	UIntptr: query.FieldIdentifier[[]uintptr]{
		Field:     "UIntptr",
		ModelType: basicSlicesPointerType,
	},
	UpdatedAt: query.FieldIdentifier[time.Time]{
		Field:     "UpdatedAt",
		ModelType: basicSlicesPointerType,
	},
}

// Preload allows preloading the BasicSlicesPointer when doing a query
func (basicSlicesPointerConditions basicSlicesPointerConditions) Preload() condition.Condition[basicslicespointer.BasicSlicesPointer] {
	return condition.NewPreloadCondition[basicslicespointer.BasicSlicesPointer](basicSlicesPointerConditions.ID, basicSlicesPointerConditions.CreatedAt, basicSlicesPointerConditions.UpdatedAt, basicSlicesPointerConditions.DeletedAt, basicSlicesPointerConditions.Bool, basicSlicesPointerConditions.Int, basicSlicesPointerConditions.Int8, basicSlicesPointerConditions.Int16, basicSlicesPointerConditions.Int32, basicSlicesPointerConditions.Int64, basicSlicesPointerConditions.UInt, basicSlicesPointerConditions.UInt8, basicSlicesPointerConditions.UInt16, basicSlicesPointerConditions.UInt32, basicSlicesPointerConditions.UInt64, basicSlicesPointerConditions.UIntptr, basicSlicesPointerConditions.Float32, basicSlicesPointerConditions.Float64, basicSlicesPointerConditions.Complex64, basicSlicesPointerConditions.Complex128, basicSlicesPointerConditions.String, basicSlicesPointerConditions.Byte)
}
