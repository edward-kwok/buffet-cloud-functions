package cmd

import (
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestWebhook(t *testing.T) {
	tests := []struct {
		body string
		want string
	}{
		{body: `{"ok":true,"result":[{"update_id":523349956,
		"message":{"message_id":51,"from":{"id":303262877,"first_name":"YourName"},"chat":{"id":303262877,"first_name":"YourName","type":"private"},"date":1486829360,"text":"Hello"}}]}`, want: "Hello, World!"},
	}
	for _, test := range tests {
		req := httptest.NewRequest("POST", "/", strings.NewReader(test.body))
		req.Header.Add("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		Webhook(rr, req)

		out, err := ioutil.ReadAll(rr.Result().Body)

		if rr.Code != 200 {
			t.Fatalf("It is not OK")
		}

		if err != nil {
			t.Fatalf("ReadAll: %v", err)
		}
		if got := string(out); got != test.want {
			t.Errorf("HelloHTTP(%q) = %q, want %q", test.body, got, test.want)
		}

	}

}
