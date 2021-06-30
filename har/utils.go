package har

import (
	"encoding/json"
	template2 "gunship/httpimpl/execution"
	"gunship/utils"
	"io/ioutil"
	"net/url"
	"os"
)

func ReadHarFromFile(path string ) []*template2.HttpRawExchange {
	root := Root{}
	data, err := ioutil.ReadFile(path)
	utils.Panic(err,"error reading file")
	err = json.Unmarshal(data, &root)
	utils.Panic(err, "error un marshalling")
	//fmt.Println(root)

	return MakeExchangeFromHar(root.Log.Entries)

}

func MakeExchangeFromHar(entries []Entries) []*template2.HttpRawExchange {
	exchanges := []*template2.HttpRawExchange{}

	for _,ele := range entries{

		// parse request
		reqUrl, err :=url.Parse(ele.Request.URL)
		utils.Panic(err,"error parsing url")

		reqHeaders := map[string][]string{}
		for _, header := range ele.Request.Headers {
			if header.Name == "Cookie" || header.Name == "Accept-Encoding"{
				continue
			}
			if _, ok := reqHeaders[header.Name]; !ok {
				reqHeaders[header.Name] = []string{}
			}
			reqHeaders[header.Name] = append(reqHeaders[header.Name], header.Value)
		}


		//request := template.RawReqUrlHeaders(ele.Request.Method, reqUrl, reqHeaders)
		request := template2.RawRequestBuilder().
			SetMethod(ele.Request.Method).
			SetFromUrl(reqUrl).
			SetHeaders(reqHeaders).
			SetBody(ele.Request.Postdata.Text).
			Build()


		// parse response
		respHeaders := map[string][]string{}
		for _, header := range ele.Response.Headers {
			//fmt.Printf("%d : %s=%s\n", idx, header.Name, header.Value)
			if _, ok := respHeaders[header.Name]; !ok {
				respHeaders[header.Name] = []string{}
			}
			respHeaders[header.Name] = append(respHeaders[header.Name], header.Value)
		}


		response := template2.NewRawResponse(respHeaders, ele.Response.Content.Text)
		exchanges = append(exchanges, &template2.HttpRawExchange{
			Request:  request,
			Response: response,
		})
	}
	return exchanges

}

type RequestFilerFunc func(exchange *Entries) bool
// HarFilter Reads and filters out useless fields and entries and rewrites file
func HarFilter(path string, filter RequestFilerFunc){
	root := Root{}
	data, err := ioutil.ReadFile(path)
	utils.Panic(err,"error reading file")
	err = json.Unmarshal(data, &root)
	utils.Panic(err, "error un marshalling")
	oldEntries := root.Log.Entries
	newEntries := []Entries{}

	for _, e:= range oldEntries{
		if filter(&e) {
			newEntries = append(newEntries, e )
		}
	}
	root.Log.Entries = newEntries
	file, err := os.Create(path + ".json")
	utils.Panic(err, "erro opening file")

	encoder := json.NewEncoder(file)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent(""," ")
	err = encoder.Encode(root)
	utils.Panic(err, "error encoding json to file")

}

