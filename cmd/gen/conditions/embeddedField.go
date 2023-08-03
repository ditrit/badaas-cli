package conditions

import (
	"errors"
	"go/types"
)

// Generate conditions for a embedded field using the "generator"
// it will generate a condition for each of the field of the embedded field's type
func generateForEmbeddedField[T any](file *File, field Field, generator CodeGenerator[T]) []T {
	embeddedStructType, ok := field.Type.Underlying().(*types.Struct)
	if !ok {
		panic(errors.New("unreachable! embedded objects are always structs"))
	}

	fields, err := getStructFields(embeddedStructType, field.Tags.getEmbeddedPrefix())
	if err != nil {
		// embedded field's type has not fields
		return []T{}
	}

	return generator.ForEachField(file, fields)
}
