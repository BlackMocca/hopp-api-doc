package methods

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/gofrs/uuid"
	"github.com/gosimple/slug"
	"github.com/knadh/stuffbin"
	"github.com/pkg/browser"
	"github.com/spf13/cast"
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
					folders[folderIndex].Requests[reqIndex].PrepareText()
				}
			}
		}
	}
	for colIndex := range colls {
		if len(colls[colIndex].Requests) > 0 {
			for reqIndex := range colls[colIndex].Requests {
				colls[colIndex].Requests[reqIndex].PrepareText()
			}
		}
		prepareRequests(colls[colIndex].Folders)
		prepareData(colls[colIndex].Folders)
	}
}

func getSlug(data interface{}) string {
	if v, ok := data.(Collection); ok && v.Name != "" {
		return slug.Make(fmt.Sprintf("Folder: %s", v.Name))
	}
	if v, ok := data.(Requests); ok && v.Method != "" && v.Name != "" {
		path := strings.ReplaceAll(v.Name, "/", "")
		path = strings.ReplaceAll(path, ":", "")
		return slug.Make(fmt.Sprintf("[%s] %s", v.Method, slug.Make(path)))
	}
	return slug.Make(data.(string))
}

func writeFile(path string, fs stuffbin.FileSystem) error {
	if err := os.MkdirAll(path, 0755); err != nil {
		return err
	}
	var tmpdir = filepath.Join("/tmp", path)
	if err := os.MkdirAll(tmpdir, 0755); err != nil {
		return err
	}

	for _, filename := range fs.List() {
		fc, err := fs.Get(filename)
		if err != nil {
			return err
		}
		defer fc.Close()

		writePath := filepath.Join(tmpdir, filename)
		if err := os.WriteFile(writePath, fc.ReadBytes(), 0664); err != nil {
			return err
		}
	}

	cmd := exec.Command("sh", "-c", fmt.Sprintf(`mv -f %s %s`, filepath.Join(tmpdir, "*"), path))
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// GenerateDocs generates the Documentation site from the hoppscotch-collection.json
func GenerateDocs(output string, exportPathfile string, servePort int, isOpenBrowser bool, templatedir string) error {
	execPath, err := os.Executable() //get Executable Path for StuffBin
	if err != nil {
		return err
	}
	fs, err := initFileSystem(execPath, templatedir) //Init Virtual FS
	if err != nil {
		return err
	}

	colls, err := ReadCollection(exportPathfile)
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
		"getSlug":             getSlug,
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

	sidebarTmpl, err := stuffbin.ParseTemplates(fmap, fs, "/_sidebar.md")
	if err != nil {
		panic(err)
	}
	// f will be used to store rendered templates in memory.
	var sidebarF FileTrunk

	// Execute the template to the file.
	if err = sidebarTmpl.Execute(&sidebarF, colls); err != nil {
		return err
	}

	if err := fs.Add(stuffbin.NewFile("/README.md", &f, f.Bytes())); err != nil {
		return err
	}

	if err := fs.Delete("/_sidebar.md"); err != nil {
		return err
	}
	if err := fs.Add(stuffbin.NewFile("/_sidebar.md", &sidebarF, sidebarF.Bytes())); err != nil {
		return err
	}

	switch output != "" {
	case true:
		output := filepath.Clean(output)

		/* writing file */
		if err := writeFile(output, fs); err != nil {
			return err
		}
		fmt.Println("Writing File Successful!!")
	default:
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			out, err := fs.Read("index.html")
			if err != nil {
				log.Println(err)
			}
			w.Write(out)
		})
		PortStr := ":" + strconv.Itoa(servePort)
		URL := fmt.Sprintf("http://localhost%s", PortStr)

		http.Handle("/docs/", http.StripPrefix("/docs/", fs.FileServer()))

		log.Printf("\033[1;36mServer Listening at %s\033[0m", URL)

		if isOpenBrowser { //Check if User wants to open the Broswer
			browser.OpenURL(URL) // AutoOpen the Broswer
		}

		http.ListenAndServe(PortStr, nil)
	}
	return nil
}
