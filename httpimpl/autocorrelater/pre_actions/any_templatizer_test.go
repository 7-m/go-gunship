package pre_actions

import (
	testutils2 "gunship/httpimpl/testutils"
	"testing"
)

func TestAnyTemplater_ProcessRequest(t *testing.T) {
	var anyTemplater anyTemplater;

	exchanges := testutils2.GetExhchanges1("")
	ctx := map[string]map[string]string{"id" : {"1234" : "id0"}}
	anyTemplater.ProcessRequest(exchanges[1].Request, ctx)

	if exchanges[1].Request.Path != "/api/v1/{id0}" {
		t.Fail()
	}

}
