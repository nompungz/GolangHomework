package nompungz

import (
	"testing"
)

func TestRomanConverter(t *testing.T){
	test := []struct {
		name string
		input int
		expected string
	}{
		{
			name:"test01",
			input :8,
			expected : "VIII",
		},
		{
			name:"test02",
			input :19,
			expected : "XIX",
		},
		{
			name:"test03",
			input :34,
			expected : "XXXIV",
		},
		{
			name:"test04",
			input :43,
			expected : "XLIII",
		},
		{
			name:"test05",
			input :67,
			expected : "LXVII",
		},
		{
			name:"test06",
			input :97,
			expected : "XCVII",
		},
		{
			name:"test07",
			input :100,
			expected : "C",
		},
	}
	for _,tt := range test {
		t.Run(tt.name,func(t *testing.T){
			result := romanConverter(tt.input)
			if result != tt.expected{
				t.Errorf("expect : % #v but got :% #v", tt.expected, result)
			}
		})
		
	}
}