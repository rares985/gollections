package gollections

import "testing"

func compareErrors(t *testing.T, a, b error) bool {
	t.Helper()
	if a == nil && b == nil {
		return true
	}
	if a == nil && b != nil {
		t.Fatalf("Expected %v but got none", b.Error())
	}

	if a != nil && b == nil {
		t.Fatalf("Expected no error but got %v", a.Error())
	}

	return a.Error() == b.Error()
}
