package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gorilla/mux"
)

func TestGuestbookMessageCanBeAddedAndRead(t *testing.T) {
	previousLists := lists
	lists = map[string][]string{}
	t.Cleanup(func() {
		lists = previousLists
	})

	pushRequest := mux.SetURLVars(
		httptest.NewRequest(http.MethodGet, "/rpush/guestbook/local-test", nil),
		map[string]string{"key": "guestbook", "value": "local-test"},
	)
	pushResponse := httptest.NewRecorder()
	ListPushHandler(pushResponse, pushRequest)
	if pushResponse.Code != http.StatusOK {
		t.Fatalf("push status = %d, want %d", pushResponse.Code, http.StatusOK)
	}

	listRequest := mux.SetURLVars(
		httptest.NewRequest(http.MethodGet, "/lrange/guestbook", nil),
		map[string]string{"key": "guestbook"},
	)
	listResponse := httptest.NewRecorder()
	ListRangeHandler(listResponse, listRequest)
	if listResponse.Code != http.StatusOK {
		t.Fatalf("list status = %d, want %d", listResponse.Code, http.StatusOK)
	}

	var entries []string
	if err := json.Unmarshal(listResponse.Body.Bytes(), &entries); err != nil {
		t.Fatalf("decode list response: %v", err)
	}
	if want := []string{"local-test"}; !reflect.DeepEqual(entries, want) {
		t.Fatalf("list entries = %v, want %v", entries, want)
	}
}
