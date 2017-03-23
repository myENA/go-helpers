package helpers_test

import (
	"testing"

	"github.com/myENA/go-helpers"
)

func TestCleanStringSlice(t *testing.T) {
	s := []string{"ok", " extralead", "extrafollow ", "       "}
	expected := []string{"ok", "extralead", "extrafollow"}
	cleaned := helpers.CleanStringSlice(s)
	t.Run("Length Test", func(t *testing.T) {
		if len(expected) != len(cleaned) {
			t.Logf("Expected \"%d\", saw \"%d\"", len(expected), len(cleaned))
			t.FailNow()
		}
	})
	t.Run("Result Test", func(t *testing.T) {
		if "ok" != cleaned[0] || "extralead" != cleaned[1] || "extrafollow" != cleaned[2] {
			t.Logf("Expected \"%+v\", saw \"%+v\"", expected, cleaned)
			t.Fail()
		}
	})
}

func TestUniqueStringSlice(t *testing.T) {
	s := []string{"ok", "dupe1", "dupe2", "     dupe2", "dupe1 ", "ok2"}
	expected := []string{"ok", "dupe1", "dupe2", "ok2"}
	unique := helpers.UniqueStringSlice(s)
	t.Run("Length Test", func(t *testing.T) {
		if len(expected) != len(unique) {
			t.Logf("Expected \"%d\", saw \"%d\"", len(expected), len(unique))
			t.FailNow()
		}
	})
	t.Run("Result Test", func(t *testing.T) {
		if "ok" != unique[0] || "dupe1" != unique[1] || "dupe2" != unique[2] || "ok2" != unique[3] {
			t.Logf("Expected \"%+v\", saw \"%+v\"", expected, unique)
			t.Fail()
		}
	})
}

func TestCombineStringSlices(t *testing.T) {
	s1 := []string{"ok", "ok", "val1", "val2", "      val3"}
	s2 := []string{"ok2", "ok", "val3", "val2"}
	expected := []string{"ok", "val1", "val2", "val3", "ok2"}
	combined, additions := helpers.CombineStringSlices(s1, s2)
	t.Run("Length Test", func(t *testing.T) {
		if len(expected) != len(combined) {
			t.Logf("Expected \"%d\", saw \"%d\"", len(expected), len(combined))
			t.FailNow()
		}
	})
	t.Run("Addition Count Test", func(t *testing.T) {
		if 1 != additions {
			t.Logf("Expected \"1\", saw \"%d\"", additions)
			t.FailNow()
		}
	})
	t.Run("Result Test", func(t *testing.T) {
		if "ok" != combined[0] || "val1" != combined[1] || "val2" != combined[2] || "val3" != combined[3] || "ok2" != combined[4] {
			t.Logf("Expected \"%+v\", saw \"%+v\"", expected, combined)
			t.Fail()
		}
	})
}

func TestRemoveStringsFromSlice(t *testing.T) {
	root := []string{"ok", "removeme1", "removeme2", "ok2"}
	remove := []string{"removeme1", "removeme2"}
	expected := []string{"ok", "ok2"}
	removed, removedCount := helpers.RemoveStringsFromSlice(root, remove)
	t.Run("Length Test", func(t *testing.T) {
		if len(expected) != len(removed) {
			t.Logf("Expected \"%d\", saw \"%d\"", len(expected), len(removed))
			t.FailNow()
		}
	})
	t.Run("Remove Count Test", func(t *testing.T) {
		if 2 != removedCount {
			t.Logf("Expected \"2\", saw \"%d\"", removedCount)
			t.FailNow()
		}
	})
	t.Run("Result Test", func(t *testing.T) {
		if "ok" != removed[0] || "ok2" != removed[1] {
			t.Logf("Expected \"%+v\", saw \"%+v\"", expected, removed)
			t.Fail()
		}
	})
}
