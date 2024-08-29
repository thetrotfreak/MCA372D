package user

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	users = []User{
		User{
			ID:       1,
			Username: "adent",
		},
		User{
			ID:       2,
			Username: "tmacmillan",
		},
	}

	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()
	expect, err := json.Marshal(users)
	if err != nil {
		t.Fatal(err)
	}

	Handler(rec, req)

	res := rec.Result()
	got, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(expect, got) {
		t.Fail()
	}

}
