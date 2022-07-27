package linkchecker_test 

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/heying-haevg/linkchecker"
)

func TestGetStatusReturns200ForUrlWithOkStatus(t *testing.T) {	
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	url := ts.URL

	got, err := linkchecker.GetStatus(url)	
	if (err != nil) {
		t.Fatal(err)
	}
	if got != http.StatusOK {
		t.Errorf("expected 200, got: %v", got)
	}
}

func TestGetStatusReturns404ForUrlWithNotFoundStatus(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	url := ts.URL

	got, err := linkchecker.GetStatus(url)	
	if (err != nil) {
		t.Fatal(err)
	}
	if got != http.StatusNotFound {
		t.Errorf("expected 404, got: %v", got)
	}
}

func TestGetStatusWithInvalidUrlReturnsError(t *testing.T) {
	_, err := linkchecker.GetStatus("invalidUrl")
	if (err == nil)	{
		t.Fatal("want an error for an invalid url")
	}
}