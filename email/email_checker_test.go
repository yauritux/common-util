package email

import (
	"fmt"
	"testing"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		in  string
		out string
		err bool
	}{
		// Invalid format.
		{in: "", err: true},
		{in: "email@", err: true},
		{in: "email@x", err: true},
		{in: "email@@example.com", err: true},
		{in: ".email@example.com", err: true},
		{in: "email.@example.com", err: true},
		{in: "email..test@example.com", err: true},
		{in: ".email..test.@example.com", err: true},
		{in: "email@at@example.com", err: true},
		{in: "some whitespace@example.com", err: true},
		{in: "email@whitespace example.com", err: true},
		{in: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa@example.com", err: true},
		{in: "email@aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa.com", err: true},

		// Unresolvable domain.
		{in: "tux+180182@wrong.example.com", err: true},

		// Valid.
		{in: "yauritux@gmail.com"},
		{in: "yauri.attamimi@automate.id"},
		{in: "email+extra@example.com"},
		{in: "EMAIL@aol.co.uk"},
		{in: "EMAIL+EXTRA@aol.co.uk"},
	}

	for _, tt := range tests {
		err := Validate(tt.in)
		if err != nil {
			if !tt.err {
				t.Errorf(`"%s": unexpected error \"%v\"`, tt.in, err)
			}
			continue
		}
		if tt.err && err == nil {
			t.Errorf(`"%s": expected error`, tt.in)
			continue
		}
	}
}

func ExampleValidate() {
	err := Validate("My+Email@wrong.example.com")
	if err != nil {
		fmt.Println("Email is not valid.")

		if err == ErrInvalidFormat {
			fmt.Println("Wrong format.")
		}

		if err == ErrUnresolvableHost {
			fmt.Println("Unresolvable host.")
		}
	}
	// Output:
	// Email is not valid.
	// Unresolvable host.
}

func TestNormalize(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{in: "email@EXAMPLE.COM. ", out: "email@example.com"},
		{in: " Email+Me@example.com. ", out: "email+me@example.com"},
	}

	for _, tt := range tests {
		normalized := Normalize(tt.in)
		if normalized != tt.out {
			t.Errorf(`%v: got "%v", want "%v"`, tt.in, normalized, tt.out)
		}
	}
}

func ExampleNormalize() {
	fmt.Println(Normalize(" Email+Me@example.com. "))
	// Output: email+me@example.com
}
