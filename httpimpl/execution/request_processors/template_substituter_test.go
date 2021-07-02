package request_processors

import (
	"gunship"
	"gunship/httpimpl/execution"
	"testing"
)

func TestTemplate_Before(t1 *testing.T) {
	ctx := map[string]string{ "id" : "123", "auth" : "eyblablabla"}
	var compiler gunship.ExecutionRequestProcessor
	compiler = NewTemplateCompiler()
	compiledRquest := &execution.HttpCompiledRequest{
		Method:              "GET",
		BaseUrl:             "https://www.example.com",
		Path:                "/api/{id}/profile",
		Body:                "",
		Headers: map[string][]string{"authorization" :{ "{auth}"}},
	}

	compiler.ProcessRequest(compiledRquest,nil, map[string]interface{}{"template": ctx})

	if compiledRquest.Path != "/api/123/profile" || compiledRquest.Headers["authorization"][0] != "eyblablabla"{
		t1.Fail()
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