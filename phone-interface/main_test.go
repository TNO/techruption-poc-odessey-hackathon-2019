package main

import (
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestResult(t *testing.T) {
	results = make(map[string]*Result)
	req := httptest.NewRequest("POST", "/setresult", strings.NewReader("id=a&result=b\n"))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	setResultHandler(w, req)
	req = httptest.NewRequest("GET", "/getresult?id=a", nil)
	w = httptest.NewRecorder()
	getResultHandler(w, req)
	body, _ := ioutil.ReadAll(w.Result().Body)
	if strings.TrimSpace(string(body)) != "b" {
		t.Errorf("Got %v, wanted b", string(body))
	}
}
