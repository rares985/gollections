package gollections

import "testing"

func compareErrors(t *testing.T, a, b error) {
	t.Helper()
	if a == nil && b == nil {
		return
	}
	if a == nil && b != nil {
		t.Fatalf("Expected %v but got none", b.Error())
	}

	if a != nil && b == nil {
		t.Fatalf("Expected no error but got %v", a.Error())
	}

	if a.Error() != b.Error() {
		t.Fatalf("Expected an error like %v, but got %v", a.Error(), b.Error())
	}
}
