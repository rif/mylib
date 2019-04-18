package utils

import "testing"

func TestToJSON(t *testing.T) {
	m := map[string]int{
		"first":  1,
		"second": 2,
	}
	jsonRep := ToJSON(m)
	expect := `{"first":1,"second":2}`
	if jsonRep != expect {
		t.Error("got bad json result: ", jsonRep)
	}
}

func TestToJSONStruct(t *testing.T) {
	js := &struct {
		Name        string
		Score       float64
		privateInfo string
	}{
		Name:        "Maria",
		Score:       9.78,
		privateInfo: "hope it will stay private",
	}
	jsonRep := ToJSON(js)
	expect := `{"Name":"Maria","Score":9.78}`
	if jsonRep != expect {
		t.Error("got bad json result: ", jsonRep)
	}
}

func TestToJSONStructMeta(t *testing.T) {
	js := &struct {
		Name        string  `json:"full_name"`
		Score       float64 `json:"score"`
		privateInfo string  `json:"private_info"`
	}{
		Name:        "Maria",
		Score:       9.78,
		privateInfo: "hope it will stay private",
	}
	jsonRep := ToJSON(js)
	expect := `{"full_name":"Maria","score":9.78}`
	if jsonRep != expect {
		t.Error("got bad json result: ", jsonRep)
	}
}

func BenchmarkToJSONStruct(b *testing.B) {
	js := &struct {
		Name        string  `json:"full_name"`
		Score       float64 `json:"score"`
		privateInfo string  `json:"private_info"`
	}{
		Name:        "Maria",
		Score:       9.78,
		privateInfo: "hope it will stay private",
	}
	for i := 0; i < b.N; i++ {
		ToJSON(js)
	}
}
