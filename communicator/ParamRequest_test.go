package communicator

import (
	"testing"
	"time"
)

func TestAddParameterString(t *testing.T) {

	paramList := RequestParams{}
	AddRequestParameter(&paramList, "Test", "foo")

	reqParam, err := NewRequestParam("Test", "foo")
	if err != nil {
		t.Fatalf("TestAddParameterString : %v", err)
	}

	expectedList := RequestParams{}
	expectedList = append(expectedList, *reqParam)

	if requestParamsCmp(paramList, expectedList) == false {
		t.Fatal("TestAddParameterString failed on compare")
	}
}

func TestAddParameterInteger(t *testing.T) {

	paramList := RequestParams{}
	AddRequestParameter(&paramList, "Test", 42)

	reqParam, err := NewRequestParam("Test", "42")
	if err != nil {
		t.Fatalf("TestAddParameterInteger : %v", err)
	}

	expectedList := RequestParams{}
	expectedList = append(expectedList, *reqParam)

	if requestParamsCmp(paramList, expectedList) == false {
		t.Fatal("TestAddParameterInteger failed on compare")
	}
}

func TestAddParameterInteger64(t *testing.T) {

	paramList := RequestParams{}
	AddRequestParameter(&paramList, "Test", (1 << 42))

	reqParam, err := NewRequestParam("Test", "4398046511104")
	if err != nil {
		t.Fatalf("TestAddParameterInteger64 : %v", err)
	}

	expectedList := RequestParams{}
	expectedList = append(expectedList, *reqParam)

	if requestParamsCmp(paramList, expectedList) == false {
		t.Fatal("TestAddParameterInteger64 failed on compare")
	}
}

func TestAddParameterBool(t *testing.T) {

	paramList := RequestParams{}
	AddRequestParameter(&paramList, "Test", true)

	reqParam, err := NewRequestParam("Test", "true")
	if err != nil {
		t.Fatalf("TestAddParameterBool : %v", err)
	}

	expectedList := RequestParams{}
	expectedList = append(expectedList, *reqParam)

	if requestParamsCmp(paramList, expectedList) == false {
		t.Fatal("TestAddParameterBoolean failed on compare")
	}
}

func TestAddParameterStringSlice(t *testing.T) {
	controlCharacterSpam := `\\/!@#%$%^&*\"' foo '\"/\\`
	brackets := `[{(<!--bar-->)}]`

	stringSlice := []string{controlCharacterSpam, brackets, ""}

	paramList := RequestParams{}
	AddRequestParameter(&paramList, "Test", stringSlice)

	reqParam1, _ := NewRequestParam("Test", controlCharacterSpam)
	reqParam2, _ := NewRequestParam("Test", brackets)

	expectedList := RequestParams{}
	expectedList = append(expectedList, *reqParam1)
	expectedList = append(expectedList, *reqParam2)

	if requestParamsCmp(paramList, expectedList) == false {
		t.Fatal("TestAddParameterStringSlice failed on compare")
	}
}

func TestAddParameterIntegerSlice(t *testing.T) {

	integerSlice := []int{2, 42, 127}

	paramList := RequestParams{}
	AddRequestParameter(&paramList, "Test", integerSlice)

	reqParam1, _ := NewRequestParam("Test", "2")
	reqParam2, _ := NewRequestParam("Test", "42")
	reqParam3, _ := NewRequestParam("Test", "127")

	expectedList := RequestParams{}
	expectedList = append(expectedList, *reqParam1)
	expectedList = append(expectedList, *reqParam2)
	expectedList = append(expectedList, *reqParam3)

	if requestParamsCmp(paramList, expectedList) == false {
		t.Fatal("TestAddParameterIntegerSlice failed on compare")
	}
}

func TestAddParameterInteger64Slice(t *testing.T) {

	integer64Slice := []int64{2, 42, (1 << 42)}

	paramList := RequestParams{}
	AddRequestParameter(&paramList, "Test", integer64Slice)

	reqParam1, _ := NewRequestParam("Test", "2")
	reqParam2, _ := NewRequestParam("Test", "42")
	reqParam3, _ := NewRequestParam("Test", "4398046511104")

	expectedList := RequestParams{}
	expectedList = append(expectedList, *reqParam1)
	expectedList = append(expectedList, *reqParam2)
	expectedList = append(expectedList, *reqParam3)

	if requestParamsCmp(paramList, expectedList) == false {
		t.Fatal("TestAddParameterInteger64Slice failed on compare")
	}
}

func TestAddParameterBoolSlice(t *testing.T) {

	boolSlice := []bool{true, false, true}

	paramList := RequestParams{}
	AddRequestParameter(&paramList, "Test", boolSlice)

	reqParam1, _ := NewRequestParam("Test", "true")
	reqParam2, _ := NewRequestParam("Test", "false")
	reqParam3, _ := NewRequestParam("Test", "true")

	expectedList := RequestParams{}
	expectedList = append(expectedList, *reqParam1)
	expectedList = append(expectedList, *reqParam2)
	expectedList = append(expectedList, *reqParam3)

	if requestParamsCmp(paramList, expectedList) == false {
		t.Fatal("TestAddParameterInteger64Slice failed on compare")
	}
}

func TestAddParameterInvalid(t *testing.T) {

	paramList := RequestParams{}

	err := AddRequestParameter(&paramList, "invalid", time.Now())

	if err == nil {
		t.Fatal("TestAddParameterInvalid failed on nil-check")
	}
}

func TestAddParameterInvalidSlice(t *testing.T) {
	type initialType string
	type initialTypeSlice []initialType
	type initialTypeSliceSlice []initialTypeSlice

	paramList := RequestParams{}

	err := AddRequestParameter(&paramList, "invalid", initialTypeSliceSlice{initialTypeSlice{initialType("Test")}})

	if err == nil {
		t.Fatal("TestAddParameterInvalidSlice failed on nil-check")
	}
}

type TestAbstractRequest struct {
}

func (ar *TestAbstractRequest) ToRequestParameters() RequestParams {
	return nil
}

func requestParamsCmp(a, b RequestParams) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
