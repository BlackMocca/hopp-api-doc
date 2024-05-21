package methods

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"time"

	"github.com/gofrs/uuid"
	"github.com/knadh/stuffbin"
	"github.com/pkg/browser"
	"github.com/spf13/cast"
	"github.com/urfave/cli"
)

// FileTrunk handles the buffer for generated README.md File
type FileTrunk struct{ bytes.Buffer }

// Name holds the FileName, here README.md
func (f *FileTrunk) Name() string { return "README.md" }

// Size holds the size of the File
func (f *FileTrunk) Size() int64 { return int64(f.Len()) }

// Mode holds the file Mode
func (f *FileTrunk) Mode() os.FileMode { return 0755 }

// ModTime holds creation time of File
func (f *FileTrunk) ModTime() time.Time { return time.Now() }

// IsDir checks if True
func (f *FileTrunk) IsDir() bool { return false }

// Sys - I have no idea
func (f *FileTrunk) Sys() interface{} { return nil }

func isArray(data interface{}) bool {
	switch data.(type) {
	case []interface{}:
		return true
	default:
		return false
	}
}

func getRequestExamples(data interface{}) []ExampleResponse {
	var examples = make([]ExampleResponse, 0)
	if v, ok := data.([]RequestVariable); ok && len(v) > 0 {
		for _, item := range v {
			if len(item.Examples) > 0 {
				examples = append(examples, item.Examples...)
			}
		}
	}
	return examples
}

func getRequestVariables(data interface{}) []RequestVariable {
	if v, ok := data.([]RequestVariable); ok && len(v) > 0 {
		return RequestVariables(v).GetRequestVariables()
	}
	return make([]RequestVariable, 0)
}

func getDataType(data interface{}) string {
	if _, err := cast.ToIntE(data); err == nil {
		return reflect.Int.String()
	}
	if _, err := cast.ToBoolE(data); err == nil {
		return reflect.Bool.String()
	}
	if _, err := cast.ToFloat64E(data); err == nil {
		return "double"
	}
	if vals, err := cast.ToSliceE(data); err == nil {
		if len(vals) > 0 {
			itemDataType := getDataType(vals[0])
			return fmt.Sprintf("%s[%s]", reflect.Array.String(), itemDataType)
		}
		return reflect.Array.String()
	}
	if _, err := cast.ToStringMapE(data); err == nil {
		return "object"
	}
	if v, err := cast.ToStringE(data); err == nil && v != "" {
		if uid := uuid.FromStringOrNil(v); uid != uuid.Nil {
			return "uuid"
		}
		return reflect.String.String()
	}

	return "unknow"
}

func tabStart() string {
	return "<!-- tabs:start -->"
}
func tabEnd() string {
	return "<!-- tabs:end -->"
}

func prepareData(colls []Collection) {
	if len(colls) == 0 {
		return
	}
	var prepareRequests = func(folders []Collection) {
		if len(folders) == 0 {
			return
		}
		for folderIndex := range folders {
			if reqs := folders[folderIndex].Requests; len(reqs) > 0 {
				for reqIndex := range reqs {
					folders[folderIndex].Requests[reqIndex].ReplaceVariableKey()
				}
			}
		}
	}
	for colIndex := range colls {
		if len(colls[colIndex].Requests) > 0 {
			for reqIndex := range colls[colIndex].Requests {
				colls[colIndex].Requests[reqIndex].ReplaceVariableKey()
			}
		}
		prepareRequests(colls[colIndex].Folders)
		prepareData(colls[colIndex].Folders)
	}
}

// GenerateDocs generates the Documentation site from the hoppscotch-collection.json
func GenerateDocs(c *cli.Context) error {
	execPath, err := os.Executable() //get Executable Path for StuffBin
	if err != nil {
		return err
	}
	fs, err := initFileSystem(execPath) //Init Virtual FS
	if err != nil {
		return err
	}

	colls, err := ReadCollection(c.Args().Get(0))
	if err != nil {
		return err
	}
	prepareData(colls)

	// FuncMap for the HTML template
	fmap := map[string]interface{}{
		"html":                func(val string) string { return val },
		"isArray":             isArray,
		"getDataType":         getDataType,
		"tabStart":            tabStart,
		"tabEnd":              tabEnd,
		"getRequestExamples":  getRequestExamples,
		"getRequestVariables": getRequestVariables,
	}

	t, err := stuffbin.ParseTemplates(fmap, fs, "/template.md")
	if err != nil {
		panic(err)
	}
	// f will be used to store rendered templates in memory.
	var f FileTrunk

	// Execute the template to the file.
	if err = t.Execute(&f, colls); err != nil {
		return err
	}

	if err := fs.Add(stuffbin.NewFile("/README.md", &f, f.Bytes())); err != nil {
		return err
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		out, err := fs.Read("templates/index.html")
		if err != nil {
			log.Println(err)
		}
		w.Write(out)
	})
	PortStr := ":" + strconv.Itoa(c.Int("port"))
	URL := fmt.Sprintf("http://localhost%s", PortStr)

	http.Handle("/static/", http.StripPrefix("/static/", fs.FileServer()))

	log.Printf("\033[1;36mServer Listening at %s\033[0m", URL)

	if !c.Bool("browser") { //Check if User wants to open the Broswer
		browser.OpenURL(URL) // AutoOpen the Broswer
	}

	http.ListenAndServe(PortStr, nil)
	return nil
}
