package methods

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/spf13/cast"
)

// Collection hold the structure of the basic `postwoman-collection.json`
type Collection struct {
	// Name of the Whole Collection
	Name string `json:"name"`
	// Folders JSON Type
	// Folders []Folders `json:"folders"`
	Folders []Collection `json:"folders"`
	// Requests inside the Collection
	Requests []Requests `json:"requests"`

	// Set Data into Collection
	Property FolderProperties `json:"data"`
}

// Folders can be organized to Folders
type Folders struct {
	// Folder name
	Name string `json:"name"`
	// Requests inside the Folder
	Requests []Requests `json:"requests"`

	// Set Data into Collection
	Property FolderProperties `json:"data"`
}

// Headers are the Request Headers
type Headers struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Requests are the Request Model in JSON
type Requests struct {
	// Base URL of the Request
	URL string `json:"endpoint"`
	// Path is the enpoint path
	// URL+PATH = Full URL
	Path string `json:"path"`
	// Request Method - GET,POST,PUT,PATCH,DELETE
	Method string `json:"method"`
	// Authentication Type - Bearer Token or Basic Auth
	Auth AuthType `json:"auth"`
	// Username for Basic Auth
	User string `json:"httpUser"`
	// Password for Basic Auth
	Pass              string `json:"httpPassword"`
	PasswordFieldType string `json:"passwordFieldType"`
	// Request Headers if any- Key,Value pairs
	Headers []Headers `json:"headers"`
	// Params for Get Requests
	Params []interface{} `json:"params"`
	// Body Params for POST requests and forth
	Bparams []interface{} `json:"bodyParams"`
	// Raw Input. Not Formatted JSON
	RawParams string `json:"rawParams"`
	// If RawInputs are used or Not
	RawInput bool `json:"rawInput"`
	// Content Type of Request
	Body             Body              `json:"body"`
	RequestType      string            `json:"requestType"`
	PreRequestScript string            `json:"preRequestScript"`
	TestScript       string            `json:"testScript"`
	RequestVariable  []RequestVariable `json:"requestVariables"`
	// Label of Collection
	Label string `json:"label"`
	// Name of the Request
	Name string `json:"name"`
	// Number of Collection
	Collection int `json:"collection"`
}

func (r *Requests) ReplaceVariableKey() {
	if len(r.RequestVariable) == 0 {
		return
	}

	if reqs := RequestVariables(r.RequestVariable).GetRequestVariables(); len(reqs) > 0 {
		for _, req := range reqs {
			if req.Value != "" {
				var regex, err = regexp.Compile(req.Value)
				if err == nil && regex.MatchString(r.URL) {
					r.URL = regex.ReplaceAllString(r.URL, fmt.Sprintf(":%s", req.Key))
				}
			}
		}
	}
}

type FolderProperties struct {
	Auth    AuthType  `json:"auth"`
	Headers []Headers `json:"headers"`
}

func (f *FolderProperties) UnmarshalJSON(data []byte) error {
	var dataJSON string
	if err := json.Unmarshal(data, &dataJSON); err != nil {
		return err
	}

	if json.Valid([]byte(dataJSON)) {
		var tmp = struct {
			Auth    AuthType  `json:"auth"`
			Headers []Headers `json:"headers"`
		}{}
		if err := json.Unmarshal([]byte(dataJSON), &tmp); err != nil {
			fmt.Println(err)
			return err
		}
		f.Auth = tmp.Auth
		f.Headers = tmp.Headers
	}
	return nil
}

type AuthType struct {
	Key   string `json:"key"`
	AddTo string `json:"addTo"`
	Value string `json:"value"`

	// value enum => none,inherit,bearer
	AuthType string `json:"authType"`

	AuthActive bool `json:"authActive"`

	// JWT Token on auth type bearer
	Token string `json:"token"`

	// for basic auth
	Username string `json:"username"`
	Password string `json:"password"`
}

// BodyParams include the Body Parameters
type BodyParams struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Body struct {
	Body        interface{} `json:"body"`
	ContentType string      `json:"contentType"`
}

func (b *Body) UnmarshalJSON(data []byte) error {
	type bodyData struct {
		Body        interface{} `json:"body"`
		ContentType string      `json:"contentType"`
	}
	var tmp = bodyData{}

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	switch {
	case strings.Contains("multipart/form-data", tmp.ContentType):
		if tmp.Body != nil {
			if body, ok := tmp.Body.([]interface{}); ok {
				for index, valM := range body {
					isFile := valM.(map[string]interface{})["isFile"].(bool)
					_, valsIsArray := valM.(map[string]interface{})["value"].([]interface{})

					if valsIsArray && isFile {
						tmp.Body.([]interface{})[index].(map[string]interface{})["value"] = "Web Browser File"
						continue
					}
				}
			}
		}
	case strings.Contains("application/x-www-form-urlencoded", tmp.ContentType):
		if tmp.Body != nil {
			str := tmp.Body.(string)
			var replaceBody = []interface{}{}
			for _, line := range strings.Split(str, "\n") {
				if index := strings.Index(line, ": "); index != -1 {
					key := line[0:index]
					value := strings.Trim(strings.TrimPrefix(line[index:len(line)-1], ": "), `"`)
					value = strings.ReplaceAll(value, `\"`, `"`)

					replaceBody = append(replaceBody, map[string]interface{}{
						"key":   key,
						"value": value,
					})
				}
			}
			tmp.Body = replaceBody
		}
	default:
		b.Body = "API DOC support content type json,form-data,urlencoded only"
	}

	b.Body = tmp.Body
	b.ContentType = tmp.ContentType
	return nil
}

type RequestVariable struct {
	Key      string            `json:"key"`
	Value    string            `json:"value"`
	Examples []ExampleResponse `json:"-"`
}

type RequestVariables []RequestVariable

func (r RequestVariables) GetRequestVariables() []RequestVariable {
	var requests = make([]RequestVariable, 0)
	if len(r) > 0 {
		for _, item := range r {
			if len(item.Examples) == 0 {
				requests = append(requests, item)
			}
		}
	}
	return requests
}

func newExampleResponseFromText(key string, response string) *ExampleResponse {
	splits := strings.Split(key, "_")
	if len(splits) > 2 && splits[0] == "EXAMPLE" {
		status := cast.ToInt(splits[1])
		name := strings.Join(splits[2:], " ")

		if status > 0 && name != "" {
			return &ExampleResponse{
				Status:   status,
				Name:     name,
				Response: response,
			}
		}
	}
	return nil
}

func (b *RequestVariable) UnmarshalJSON(data []byte) error {
	type bodyData struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}
	var tmp = bodyData{}

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	b.Key = tmp.Key
	b.Value = tmp.Value
	if len(b.Examples) == 0 {
		b.Examples = make([]ExampleResponse, 0)
	}

	if tmp.Key != "" {
		if exampleResponse := newExampleResponseFromText(tmp.Key, tmp.Value); exampleResponse != nil {
			b.Examples = append(b.Examples, *exampleResponse)
		}
	}

	return nil
}

type ExampleResponse struct {
	Status   int
	Name     string
	Response string
}
