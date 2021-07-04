package har

import "time"

type Root struct {
	Log Log `json:"log"`
}
type Creator struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}
type Browser struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}
type Pagetimings struct {
	Oncontentload int `json:"onContentLoad"`
	Onload        int `json:"onLoad"`
}
type Pages struct {
	Starteddatetime time.Time `json:"startedDateTime"`
	ID              string    `json:"id"`
	Title           string    `json:"title"`
	//Pagetimings     Pagetimings `json:"pageTimings"`
}
type Headers struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
type Cookies struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
type Querystring struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Content struct {
	Mimetype string `json:"mimeType"`
	Size     int    `json:"size"`
	Text     string `json:"text"`
}
type Response struct {
	Status      int       `json:"status"`
	Statustext  string    `json:"statusText"`
	Httpversion string    `json:"httpVersion"`
	Headers     []Headers `json:"headers"`
	Cookies     []Cookies `json:"cookies"`
	Content     Content   `json:"content"`
	Redirecturl string    `json:"redirectURL"`
	Headerssize int       `json:"headersSize"`
	Bodysize    int       `json:"bodySize"`
}
type Cache struct {
}
type Timings struct {
	Blocked float32 `json:"blocked"`
	DNS     float32 `json:"dns"`
	Connect float32 `json:"connect"`
	Ssl     float32 `json:"ssl"`
	Send    float32 `json:"send"`
	Wait    float32 `json:"wait"`
	Receive float32 `json:"receive"`
}
type Postdata struct {
	Mimetype string        `json:"mimeType"`
	Params   []interface{} `json:"params"`
	Text     string        `json:"text"`
}
type Request struct {
	Bodysize    int           `json:"bodySize"`
	Method      string        `json:"method"`
	URL         string        `json:"url"`
	Httpversion string        `json:"httpVersion"`
	Headers     []Headers     `json:"headers"`
	Cookies     []Cookies     `json:"cookies"`
	Querystring []Querystring `json:"queryString"`
	Headerssize int           `json:"headersSize"`
	Postdata    Postdata      `json:"postData"`
}
type Entries struct {
	//Pageref         string    `json:"pageref"`
	Starteddatetime time.Time `json:"startedDateTime"`
	Request         Request   `json:"request,omitempty"`
	Response        Response  `json:"response"`
	//Cache           Cache     `json:"cache"`
	Timings Timings `json:"timings"`
	Time    float32 `json:"time"`
	//Securitystate   string    `json:"_securityState"`
	Serveripaddress string `json:"serverIPAddress"`
	Connection      string `json:"connection"`
}
type Log struct {
	//Version string    `json:"version"`
	//Creator Creator   `json:"creator"`
	//Browser Browser   `json:"browser"`
	//Pages   []Pages   `json:"pages"`
	Entries []Entries `json:"entries"`
}
