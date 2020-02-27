package describe

import "testing"

func TestLoad(t *testing.T) {
	desc, err := Load("describe_example.yaml")
	if err != nil {
		t.Error(err)
	}
	_ = desc
}
