package auth

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func TestGetAPIKey(t *testing.T){

	tests := []struct{
		key, value, expected, expectedErr string
	}{
		{
			expectedErr: "no authorization header",
		},
		{
			expectedErr : "no authorization header",
            key: "Authorization",
		},
		{
			key: "Authorization",
			value: "-",
			expectedErr: "malformed authorization header",
		},
		{
			key: "Authorization",
			value: "Bearer xxxxxxx",
			expectedErr: "malformed authorization header",
		},
		{
         key: "Authorization",
		 value: "ApiKey xxxxxxxx",
		 expected: "xxxxxxxx",
		 expectedErr: "not expecting an error",
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("TestGetAPIKey Case #%v", i), func(t *testing.T) {
			header := http.Header{}
			header.Add(test.key, test.value)

			output, err := GetAPIKey(header)

			if err != nil {
				if strings.Contains(err.Error(), test.expectedErr){
					return
				}
				t.Errorf("Unexpected: TestGetAPIKey: %v\n", err)
				return
			}

			if output != test.expected {
				t.Errorf("unexpected: TestGetAPIKey:%s", output)
				return
			}
		})
	}
}