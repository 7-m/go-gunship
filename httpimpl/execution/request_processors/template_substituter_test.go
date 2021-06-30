package request_processors

import (
	"gunship/execution"
	"net/http"
	"testing"
)

func TestTemplate_Before(t1 *testing.T) {
	ctx := map[string]string{ "id" : "123", "authid" : "eyblablabla"}
	var compiler execution.RequestProcessor
	compiler = &templateCompiler{}
	request, err := http.NewRequest("GET","https://www.example.com/api/{id}/profile",nil)
	request.Header.Add("Authorization", "{authid}")

	if err != nil {
		t1.Fatalf("error while parsing url %v", err)
	}

	// fixme changing below from request to nil, fix by using compiled request isntead of http.reqauest
	compiler.ProcessRequest(nil,nil, map[string]interface{}{"template": ctx})

	if request.URL.String() != "https://www.example.com/api/123/profile" || request.Header.Get("Authorization") != "eyblablabla"{
		t1.Fatalf(" got %v ", request.URL)
	}
}

func Test_replace(t *testing.T) {
	type args struct {
		s    string
		vars map[string]string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "success case",
			args:    args{
				s:    "/api/{userid}/profile/{a}",
				vars: map[string]string{"userid":"100uid11", "a":"eyf315652t4v"},
			},
			want:    "/api/100uid11/profile/eyf315652t4v",
			wantErr: false,
		},
		{
			name: "fail case missing close braces",
			args: args{
				s:    "/api/{user/profile",
				vars: nil,
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "pass case",
			args: args{
				s:    "/api/userid}",
				vars: nil,
			},
			want:    "/api/userid}",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := replace(tt.args.s, tt.args.vars)
			if (err != nil) != tt.wantErr {
				t.Errorf("replace() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("replace() got = %v, want %v", got, tt.want)
			}
		})
	}
}