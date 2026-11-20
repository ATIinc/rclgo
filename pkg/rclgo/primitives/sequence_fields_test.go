package primitives

import (
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
	"reflect"
	"regexp"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	// Fields written by the generated <T>__Sequence_to_C (primitiveTypesTemplate
	// in pkg/gogen/templates.go).
	primitiveSequenceFields = []string{"data", "size", "capacity", "is_rosidl_buffer", "owns_rosidl_buffer"}
	// Fields written by Char__Sequence_to_C in char.go.
	charSequenceFields = []string{"data", "size", "capacity", "is_rosidl_buffer", "owns_rosidl_buffer"}
	// Fields written by String__Sequence_to_C and U16String__Sequence_to_C.
	stringSequenceFields = []string{"data", "size", "capacity", "is_rosidl_buffer", "owns_rosidl_buffer"}
)

// sequenceConverterFields maps each exported C sequence alias to the fields its
// __Sequence_to_C converter writes.
var sequenceConverterFields = map[string]struct {
	typ    reflect.Type
	fields []string
}{
	"CBool__Sequence":      {reflect.TypeFor[CBool__Sequence](), primitiveSequenceFields},
	"CByte__Sequence":      {reflect.TypeFor[CByte__Sequence](), primitiveSequenceFields},
	"CChar__Sequence":      {reflect.TypeFor[CChar__Sequence](), charSequenceFields},
	"CFloat32__Sequence":   {reflect.TypeFor[CFloat32__Sequence](), primitiveSequenceFields},
	"CFloat64__Sequence":   {reflect.TypeFor[CFloat64__Sequence](), primitiveSequenceFields},
	"CInt8__Sequence":      {reflect.TypeFor[CInt8__Sequence](), primitiveSequenceFields},
	"CInt16__Sequence":     {reflect.TypeFor[CInt16__Sequence](), primitiveSequenceFields},
	"CInt32__Sequence":     {reflect.TypeFor[CInt32__Sequence](), primitiveSequenceFields},
	"CInt64__Sequence":     {reflect.TypeFor[CInt64__Sequence](), primitiveSequenceFields},
	"CUint8__Sequence":     {reflect.TypeFor[CUint8__Sequence](), primitiveSequenceFields},
	"CUint16__Sequence":    {reflect.TypeFor[CUint16__Sequence](), primitiveSequenceFields},
	"CUint32__Sequence":    {reflect.TypeFor[CUint32__Sequence](), primitiveSequenceFields},
	"CUint64__Sequence":    {reflect.TypeFor[CUint64__Sequence](), primitiveSequenceFields},
	"CString__Sequence":    {reflect.TypeFor[CString__Sequence](), stringSequenceFields},
	"CU16String__Sequence": {reflect.TypeFor[CU16String__Sequence](), stringSequenceFields},
}

var sequenceAliasName = regexp.MustCompile(`^C\w+__Sequence$`)

// exportedSequenceAliases parses this package's sources so a newly generated
// sequence type cannot be left out of sequenceConverterFields.
func exportedSequenceAliases(t *testing.T) []string {
	t.Helper()
	paths, err := filepath.Glob("*.go")
	require.NoError(t, err)
	fset := token.NewFileSet()
	var names []string
	for _, path := range paths {
		file, err := parser.ParseFile(fset, path, nil, parser.SkipObjectResolution)
		require.NoError(t, err)
		ast.Inspect(file, func(n ast.Node) bool {
			if spec, ok := n.(*ast.TypeSpec); ok && sequenceAliasName.MatchString(spec.Name.Name) {
				names = append(names, spec.Name.Name)
			}
			return true
		})
	}
	slices.Sort(names)
	return names
}

// A field the converter does not write keeps whatever the destination memory
// held, which for malloc'd structs is garbage (see Lyrical's rosidl buffer flags).
func TestSequenceToCWritesEveryField(t *testing.T) {
	aliases := exportedSequenceAliases(t)
	require.NotEmpty(t, aliases)
	for _, name := range aliases {
		_, ok := sequenceConverterFields[name]
		assert.Truef(t, ok, "%s has no entry in sequenceConverterFields; add the fields its __Sequence_to_C writes", name)
	}

	for name, c := range sequenceConverterFields {
		assert.Containsf(t, aliases, name, "%s is in sequenceConverterFields but not exported by the package", name)
		var actual []string
		for f := range c.typ.Fields() {
			if f.Name != "_" {
				actual = append(actual, f.Name)
			}
		}
		assert.ElementsMatchf(t, c.fields, actual,
			"%s (C struct %s) fields differ from those its __Sequence_to_C converter writes; update the converter to write every field, then this list", name, c.typ.Name())
	}
}
