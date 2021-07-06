package main

import (
	"reflect"
	"testing"
)

//func TestMain(t *testing.T) {
//
//}

func TestPreprocessURL(t *testing.T) {
	urls := []string{"google.com", "http://someweb.com"}
	expectedURLs := []string{"http://google.com", "http://someweb.com"}
	processedURLs := preprocessURL(urls)

	if !reflect.DeepEqual(expectedURLs, processedURLs) {
		t.Fatalf("expected: %v, got: %v", expectedURLs, processedURLs)
	}
}
