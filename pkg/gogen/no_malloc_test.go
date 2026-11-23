package gogen

import (
	"io/fs"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestNoCMalloc scans every non-test Go source under pkg/, cmd/ and internal/,
// which includes the generator templates in this package.
func TestNoCMalloc(t *testing.T) {
	repo := os.DirFS("../..")
	scanned := 0
	for _, dir := range []string{"pkg", "cmd", "internal"} {
		err := fs.WalkDir(repo, dir, func(path string, d fs.DirEntry, err error) error {
			if err != nil || d.IsDir() || !strings.HasSuffix(path, ".go") || strings.HasSuffix(path, "_test.go") {
				return err
			}
			src, err := fs.ReadFile(repo, path)
			if err != nil {
				return err
			}
			scanned++
			if strings.Contains(string(src), "C.malloc(") {
				t.Errorf("%s: use C.calloc; Go's write barrier reads the old value of pointer slots Go stores into, and uninitialized C memory can crash the GC", path)
			}
			return nil
		})
		require.NoError(t, err)
	}
	require.Positive(t, scanned)
}
