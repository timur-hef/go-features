package reflection

import (
	"reflect"
	"testing"
)

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		expected := "Chris"
		var got []string

		x := struct {
			Name string
		}{expected}

		walk(x, func(input string) {
			got = append(got, input)
		})

		if len(got) != 1 {
			t.Errorf("wrong number of function calls, got %d want %d", len(got), 1)
		}

		if got[0] != expected {
			t.Errorf("got %q, want %q", got[0], expected)
		}
	})
	t.Run("part of the struct", func(t *testing.T) {
		cases := []struct {
			Name          string
			Input         interface{}
			ExpectedCalls []string
		}{
			{
				Name: "struct with one string field",
				Input: struct {
					Name     string
					PostCode int
					City     string
				}{"Chris", 1172, "Buenos Aires"},
				ExpectedCalls: []string{"Chris", "Buenos Aires"},
			},
		}

		for _, test := range cases {
			t.Run(test.Name, func(t *testing.T) {
				var got []string
				walk(test.Input, func(input string) {
					got = append(got, input)
				})

				if !reflect.DeepEqual(got, test.ExpectedCalls) {
					t.Errorf("got %v, want %v", got, test.ExpectedCalls)
				}
			})
		}
	})

	t.Run("struct", func(t *testing.T) {
		oneStructCase := struct {
			Name  string
			Input interface{}
		}{
			Name: "String Field",
			Input: struct {
				Name     string
				PostCode int
				City     string
			}{"Chris", 1172, "Buenos Aires"},
		}

		ExpectedCalls := []string{"String Field", "Chris", "Buenos Aires"}

		var got []string
		walk(oneStructCase, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, ExpectedCalls) {
			t.Errorf("got %v, want %v", got, ExpectedCalls)
		}
	})

	t.Run("with function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{33, "Berlin"}, Profile{34, "Katowice"}
		}

		var got []string
		want := []string{"Berlin", "Katowice"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

}
