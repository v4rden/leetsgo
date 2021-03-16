package ValidPalindrome

import "testing"

func TestValidPalindromeTableDrive(t *testing.T) {
	tests := []struct {
		id     string
		input  string
		result bool
	}{
		{"T1",
			"A man, a plan, a canal: Panama",
			true,
		},
		{
			"T2",
			"race a car",
			false,
		},
	}

	for _, test := range tests {
		t.Run(test.id, func(t *testing.T) {
			ans := isPalindrome(test.input)
			if ans != test.result {
				t.Errorf("wan't %v got %v", test.result, ans)
			}
		})
	}
}
