package vanilla

import "testing"

func TestJSON(t *testing.T) {
	//var data = make(map[string]interface{})
	expecting := []struct {
		original string
		result   string
	}{
		{
			"ab!", `"ab!"`,
		},
		{
			"done#", `"done#"`,
		},
		{
			"no Spaces", `"no Spaces"`,
		},
	}

	for _, v := range expecting {
		t.Run(v.original, func(t *testing.T) {
			result := JSON(v.original)
			if result != v.result {
				t.Error("Expected ", v.result, "but got", result, "instead")
			}
		})
	}
}

func TestAlphaNumeric(t *testing.T) {
	expecting := []struct {
		original string
		result   string
	}{
		{
			"ab!", "ab",
		},
		{
			"done#", "done",
		},
		{
			"no Spaces", "noSpaces",
		},
	}

	for _, v := range expecting {
		t.Run(v.original, func(t *testing.T) {
			result := AlphaNumeric(v.original)
			if result != v.result {
				t.Error("Expected ", v.result, "but got", result, "instead")
			}
		})
	}
}
