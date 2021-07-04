package testutils

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	template2 "gunship/httpimpl/execution"
	"gunship/utils"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"strconv"
)

func GetExhchanges1() []*template2.HttpRawExchange {
	// login test
	login := template2.RawRequestBuilder().
		SetBaseUrl("http://example.com").
		SetMethod("POST").
		SetPath("/api/v1/login").
		Build()
	loginresponse := template2.RawResponseBuilder().
		SetBody(`{"authorization" : "ey123"}`).
		Build()
	lognxchg := &template2.HttpRawExchange{login, loginresponse}
	createperson := template2.RawRequestBuilder().
		SetBaseUrl("http://example.com").
		SetMethod("POST").
		SetPath("/api/v1/person").
		SetBody(`{"name" : "elrich", "age":"34", "address" : "silicon valley"}`).
		AddHeader("Authorization", "ey123").
		Build()
	createpersonResponse := template2.RawResponseBuilder().
		SetBody(`{"personId" : "42"}`).
		Build()
	createxchng := &template2.HttpRawExchange{createperson, createpersonResponse}
	getperson := template2.RawRequestBuilder().
		SetBaseUrl("http://example.com").
		SetMethod("GET").
		SetPath("/api/v1/person/42").
		AddHeader("Authorization", "ey123").
		Build()
	getpersonResponse := template2.RawResponseBuilder().
		SetBody(`{"name" : "elrich", "age":"34", "address" : "silicon valley"}`).
		Build()
	getpersonxchng := &template2.HttpRawExchange{getperson, getpersonResponse}

	replaceperson := template2.RawRequestBuilder().
		SetBaseUrl("http://example.com").
		SetMethod("PUT").
		SetPath("/api/v1/person/42").
		AddHeader("Authorization", "ey123").
		SetBody(`{"name" : "elrich", "age" : "34", "address" : "tibet"}`).
		Build()
	replacepersonResponse := template2.RawResponseBuilder().
		SetBody(`{"name" : "elrich", "age" : "34", "address" : "tibet"}`).
		Build()
	replacexchng := &template2.HttpRawExchange{replaceperson, replacepersonResponse}
	return []*template2.HttpRawExchange{lognxchg, createxchng, getpersonxchng, replacexchng}
}

func CreateServer1() (*httptest.Server, *http.Client) {
	// init server
	// server test
	router := httprouter.New()
	router.PUT("/api/v1/person/:id", AuthHandler(UpdatePersonHandler))
	router.GET("/api/v1/person/:id", AuthHandler(GetPersonHandler))
	router.POST("/api/v1/person", AuthHandler(CreatePersonHandler))
	router.POST("/api/v1/login", LoginHandler)

	tes := httptest.NewServer(router)
	cli := tes.Client()
	cli.Jar, _ = cookiejar.New(nil)
	return tes, cli
}

type person struct {
	name    string
	age     int
	address string
}

// person_id -> person
var personDb = map[string]*person{}

func CreatePersonHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	all, err := ioutil.ReadAll(request.Body)
	utils.Panic(err, "error while reading body")

	p := new(person)
	err = json.Unmarshal(all, p)
	utils.Panic(err, "error while unmarshalling json")

	personId := strconv.Itoa(rand.Int())
	personDb[personId] = p

	writer.WriteHeader(http.StatusOK)
	response, err := json.MarshalIndent(struct {
		personId string
	}{
		personId: personId,
	}, "", " ")
	utils.Panic(err, "error while marshalling json")
	_, err = writer.Write(response)
	utils.Panic(err, "error while writing json response")

}
func UpdatePersonHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	all, err := ioutil.ReadAll(request.Body)
	utils.Panic(err, "error while reading body")

	replacement := new(person)
	err = json.Unmarshal(all, replacement)
	utils.Panic(err, "error while unmarshalling json")

	if _, ok := personDb[params.ByName("id")]; !ok {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	personDb[params.ByName("id")] = replacement

	writer.WriteHeader(http.StatusOK)

}

var auths = map[string]bool{}

func LoginHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	itoa := strconv.Itoa(rand.Int())
	auths[itoa] = true
	auth := utils.MustMarshal(struct {
		Authorization string
	}{
		Authorization: itoa,
	})

	writer.WriteHeader(http.StatusOK)
	_, err := writer.Write(auth)
	utils.Panic(err, "error in writing response")
}

func GetPersonHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	if p, ok := personDb[params.ByName("id")]; !ok {
		writer.WriteHeader(http.StatusNotFound)
	} else {
		_, err := writer.Write(utils.MustMarshal(p))
		utils.Panic(err, "error writing person to response")
	}
}

func AuthHandler(nextHandler httprouter.Handle) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

		if _, ok := auths[request.Header["Authorization"][0]]; !ok {
			writer.WriteHeader(401)
			return
		} else {
			nextHandler(writer, request, params)
		}
	}

}
