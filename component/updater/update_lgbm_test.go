package updater

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/metacubex/mihomo/constant"
)

func TestValidateLgbmModelSurvivesAnUnwritableSystemTempDir(t *testing.T) {
	oldHomeDir := constant.Path.HomeDir()
	constant.SetHomeDir(t.TempDir())
	t.Cleanup(func() { constant.SetHomeDir(oldHomeDir) })
	t.Setenv("TMPDIR", filepath.Join(t.TempDir(), "missing"))

	err := validateLgbmModel([]byte("not a LightGBM model"))
	if err == nil {
		t.Fatal("validateLgbmModel accepted garbage data")
	}
	if !strings.Contains(err.Error(), "invalid LightGBM model file") {
		t.Fatalf("validateLgbmModel = %v, want the model parse error instead of a temp file error", err)
	}
}
