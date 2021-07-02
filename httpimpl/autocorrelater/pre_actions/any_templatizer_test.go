package pre_actions

import (
	testutils2 "gunship/httpimpl/testutils"
	"testing"
)

func TestAnyTemplater_ProcessRequest(t *testing.T) {
	var anyTemplater anyTemplater

	exchanges := testutils2.GetExhchanges1()
	ctx := map[string]map[string]string{"personId" : {"42" : "id0"}}
	anyTemplater.ProcessRequest(exchanges[2].Request, ctx)

	if exchanges[2].Request.Path != "/api/v1/person/{id0}" {
		t.Fail()
	}

}
