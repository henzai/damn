package docker

import "testing"

func TestHasContainerByID(t *testing.T) {
	_, bool := HasContainerByID("cbac2e202ec4")
	if bool == false {
		t.Fatalf("failed test")
	}
}
