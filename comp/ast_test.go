package comp

import (
	"testing"
)

func TestAstValues(t *testing.T) {
	p := NewParamString("foo")
	if p.GetValue() != "foo" {
		t.Errorf("Expected Value to be 'foo'")
	}

	q := NewParamInt(100)
	if q.GetValue() != 100 {
		t.Errorf("Expected Value to be 100")
	}
	q.SetValue(99)
	if q.GetValue() != 99 {
		t.Errorf("Expected Value to be 99")
	}
}

// REMOVE THIS
func TestAstCheckDef(t *testing.T) {
	check := AstCheckDef{
		FnName:         "testFn",
		ActualParams:   "param1, param2",
		Children:       []AstCheckDef{{FnName: "child1"}, {FnName: "child2"}},
		ContentAddress: "/test/path",
		IsNegated:      true,
		IsRetrying:     false,
		IsGroup:        true,
	}

	if check.FnName != "testFn" {
		t.Errorf("Expected FnName to be 'testFn', got '%s'", check.FnName)
	}

	if check.ActualParams != "param1, param2" {
		t.Errorf("Expected ActualParams to be 'param1, param2', got '%s'", check.ActualParams)
	}

	if len(check.Children) != 2 {
		t.Errorf("Expected 2 children, got %d", len(check.Children))
	}

	if check.ContentAddress != "/test/path" {
		t.Errorf("Expected ContentAddress to be '/test/path', got '%s'", check.ContentAddress)
	}

	if !check.IsNegated {
		t.Error("Expected IsNegated to be true")
	}

	if check.IsRetrying {
		t.Error("Expected IsRetrying to be false")
	}

	if !check.IsGroup {
		t.Error("Expected IsGroup to be true")
	}
}
