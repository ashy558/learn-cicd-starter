package auth_test

import (
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
)

func TestGetApiKey(t *testing.T) {
	type test struct {
		input     http.Header
		wantValue string
		wantErr   error
	}
	emptyHeaderInput := http.Header{}
	emptyAuthInput := http.Header{}
	emptyAuthInput.Add("Authorization", "")
	malformedAuthInput := http.Header{}
	malformedAuthInput.Add("Authorization", "apikey")
	simpleAuthInput := http.Header{}
	simpleAuthInput.Add("Authorization", "ApiKey asdf12345")

	tests := map[string]test{
		"empty header":          {input: emptyHeaderInput, wantValue: "", wantErr: auth.ErrNoAuthHeaderIncluded},
		"empty auth header":     {input: emptyAuthInput, wantValue: "", wantErr: auth.ErrNoAuthHeaderIncluded},
		"malformed auth header": {input: malformedAuthInput, wantValue: "", wantErr: auth.ErrMalformedAuthorizationHeader},
		"simple auth header":    {input: simpleAuthInput, wantValue: "asdf12345", wantErr: nil},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			gotValue, gotErr := auth.GetAPIKey(tc.input)
			if !cmp.Equal(tc.wantValue, gotValue) || tc.wantErr != gotErr {
				t.Fatalf("[%s] test: expected: %#v,%v got: %#v,%v", name, tc.wantValue, tc.wantErr, gotValue, gotErr)
			}
		})
	}
}
