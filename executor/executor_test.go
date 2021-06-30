package executor

import (
	"github.com/tidwall/gjson"
	"testing"
)

func Test_gjson(t *testing.T) {
	blob := `
{
	"cat": "12",
	"subcat": "12009",
	"dims": [
		{
			"id": "10328",
			"dim": "1000 X 50 X 50"
		},
		{
			"id": "10605",
			"dim": "1000 X 50 X 50"
		},
		{
			"id": "10342",
			"dim": "1000 X 50 X 50"
		},
		{
			"id": "10340",
			"dim": "1200 X 50 X 50"
		}
	],
}
`
	if gjson.Get(blob, "dims.1.id").String() != "10605" {
		t.Fail()
	}
}


