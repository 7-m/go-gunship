package post_actions

import (
	template2 "gunship/httpimpl/execution"
	"testing"
)

func TestJsonCorrelator_After(t *testing.T) {
	ctx := map[string]map[string]string{}

	cor := NewJsonCorrelator(map[string]string{"empId": "id"})

	json1 := `
			{
				"empId" : 123
			}
			`
	resp1 := template2.NewRawResponse(nil, json1)

	cor.ProcessResponse(resp1, ctx)

	if val := ctx["id"]["123"]; val != "id0" {
		t.Fail()
	}

	json2 := `
			{
				"empId" : 789
			}
			`
	resp2 := template2.NewRawResponse(nil, json2)
	cor.ProcessResponse(resp2, ctx)

	if val := ctx["id"]["789"]; val != "id1" {
		t.Fail()
	}
}
