package request_processors

import (
	"gunship"
	"gunship/httpimpl/execution"
	"testing"
)

func TestTemplate_Before(t1 *testing.T) {
	ctx := map[string]string{"id": "123", "auth": "eyblablabla", "EMPID0": "EMP1001"}
	var compiler gunship.ExecutionRequestProcessor
	compiler = NewTemplateCompiler()
	compiledRquest := &execution.HttpCompiledRequest{
		Method:  "GET",
		BaseUrl: "https://www.example.com",
		Path:    "/api/{id}/profile",
		Body:    `{"empId" : "{EMPID0}"}`,
		Headers: map[string][]string{"authorization": {"{auth}"}},
	}

	compiler.ProcessRequest(compiledRquest, nil, map[string]interface{}{"template": ctx})

	expected := "/api/123/profile"
	if compiledRquest.Path != expected {
		t1.Fatalf("expected %v, got %v", expected, compiledRquest.Path)
	}
	expected = "eyblablabla"
	if compiledRquest.Headers["authorization"][0] != expected {
		t1.Fatalf("expected %v, got %v", expected, compiledRquest.Headers["authorization"][0])
	}
	expected = `{"empId" : "EMP1001"}`
	if compiledRquest.Body != expected {
		t1.Fatalf("expected %v, got %v", expected, compiledRquest.Body)
	}
}

func Test_replace(t *testing.T) {
	type args struct {
		s    string
		vars map[string]string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success case",
			args: args{
				s:    "/api/{userid}/profile/{a}",
				vars: map[string]string{"userid": "100uid11", "a": "eyf315652t4v"},
			},
			want: "/api/100uid11/profile/eyf315652t4v",
		},
		{
			name: "pass case",
			args: args{
				s:    "/api/{user/profile",
				vars: nil,
			},
			want: "/api/{user/profile",
		},
		{
			name: "pass case",
			args: args{
				s:    "/api/userid}",
				vars: nil,
			},
			want: "/api/userid}",
		},
		{
			name: "pass case",
			args: args{
				s:    "{abcdabcd",
				vars: nil,
			},
			want: "{abcdabcd",
		},
		{
			name: "pass case",
			args: args{
				s:    "{abcd{var}",
				vars: map[string]string{"var": "123"},
			},
			want: "{abcd123",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := replace(tt.args.s, tt.args.vars)
			if got != tt.want {
				t.Errorf("replace() got = %v, want %v", got, tt.want)
			}
		})
	}
}
