package common

import "testing"

func TestSubsDeleteByName(t *testing.T) {
	a := &Sub{
		Sub: "test",
		ID:  "1",
	}

	b := []Sub{*a}

	b = DeleteByName(b, "test")

	t.Log(b)
}
