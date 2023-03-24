package helper

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type TestStruct struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestReadFromRequestBody(t *testing.T) {
	testData := TestStruct{Name: "ryan", Age: 25}
	jsonData, _ := json.Marshal(testData)

	req, err := http.NewRequest("POST", "/test", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatal(err)
	}

	var result TestStruct
	ReadFromRequestBody(req, &result)

	if result.Name != testData.Name {
		t.Errorf("Expected %s but got %s", testData.Name, result.Name)
	}

	if result.Age != testData.Age {
		t.Errorf("Expected %d but got %d", testData.Age, result.Age)
	}
}

func TestWriteToResponseBody(t *testing.T) {
	testData := TestStruct{Name: "ryan", Age: 25}

	rr := httptest.NewRecorder()
	WriteToResponseBody(rr, testData)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"name":"ryan","age":25}`
	if strings.TrimRight(rr.Body.String(), "\n") != expected {
		t.Errorf("Expected body %s, but got %s", expected, rr.Body.String())
	}
}
