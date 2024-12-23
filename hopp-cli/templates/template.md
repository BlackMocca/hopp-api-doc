{{ range .  }}

# Folder: {{.Name}}


<!-- Start Root Folder Property -->
{{- if ne .Property.Auth.AuthType "" -}}
**Authentication**: `{{ .Property.Auth.AuthType | html }}`
    {{- if eq .Property.Auth.AuthType "bearer" -}}
        <table>
        <tr>
        <th>AddTo</th>
        <th>Key</th>
        <th>Value</th>
        <th>Description</th>
        </tr>
        <tr>
        <td>{{ .Property.Auth.AddTo | html }}</td>
        <td>Authorization</td>
        <td>`{{ .Property.Auth.AuthType | html}}{{" "}}{{- .Property.Auth.Token | html}}`</td>
        <td>{{ .Property.Auth.Description | html }}</td>
        </tr>
        </table>
    {{- end -}}
    {{- if eq .Property.Auth.AuthType "api-key" -}}
        <table>
        <tr>
        <th>AddTo</th>
        <th>Key</th>
        <th>Value</th>
        <th>Description</th>
        </tr>
        <tr>
        <td>{{ .Property.Auth.AddTo | html }}</td>
        <td>`{{ .Property.Auth.Key | html}}`</td>
        <td>`{{- .Property.Auth.Value | html}}`</td>
        <td>{{ .Property.Auth.Description | html }}</td>
        </tr>
        </table>
    {{- end -}}
    {{- if eq .Property.Auth.AuthType "basic" -}}
        <table>
        <tr>
        <th>AddTo</th>
        <th>Username</th>
        <th>Password</th>
        <th>Description</th>
        </tr>
        <tr>
        <td>{{ .Property.Auth.AddTo | html }}</td>
        <td>`{{ .Property.Auth.Username | html}}`</td>
        <td>`{{- .Property.Auth.Password | html}}`</td>
        <td>{{ .Property.Auth.Description | html }}</td>
        </tr>
        </table>
    {{- end -}}
{{- end -}}

{{- if .Property.Headers -}}
**Headers**: 
    <table>
    <tr>
    <th>Key</th>
    <th>Value</th>
    <th>Description</th>
    </tr>
    {{- range .Property.Headers -}}
    <tr>
    <td>`{{- .Key | html -}}`</td>
    <td>`{{- .Value | html -}}`</td>
    <td>{{- .Description | html -}}</td>
    </tr>
    {{- end -}}
    </table>
{{- end -}}
<!-- END Root Folder Property -->


<!-- Start Request Root -->
{{- range .Requests}}

---
## [{{ .Method }}] {{ .Name }} :id={{ getSlug . }}

**Method**: {{ .Method}}

**RequestURL**:  `{{ .URL | html }}{{ .Path}}`

{{ if .Headers}}
**Headers**: 
    <table>
    <tr>
    <th>Key</th>
    <th>Value</th>
    <th>Description</th>
    </tr>
    {{- range .Headers}}
    <tr>
    <td>{{- .Key | html}}</td>
    <td>`{{- .Value | html }}`</td>
    <td>{{- .Description | html }}</td>
    </tr>
    {{- end -}}
    </table>
{{- end}}
{{- if .Params}}
**Query Params**: 
    <table>
    <tr>
    <th>Key</th>
    <th>Type</th>
    <th>Value</th>
    <th>Description</th>
    </tr>
    {{- range .Params}}
    <tr>
    <td>`{{- .key | html}}`</td>
    <td><code>{{- getDataType .value}}</code></td>
    <td>{{- if .value }}`{{ .value | html }}`{{ else }} {{ end }}</td>
    <td>{{- .description | html}}</td>
    </tr>
    {{-  end}}
    </table>
{{- end -}}
{{ if ne .Auth.AuthType "None"}}
**Authentication Type**: `{{.Auth.AuthType | html}}` 
    
{{ if .Auth.Token}}
**BearerToken**: `{{ .Auth.Token | html}}`

{{ end}}

{{ if and .User .Pass}}
Username: `{{ .User}}`  
Password: `{{ .Pass}}`
{{ end}}
{{ end}}
{{ if and .RawParams (ne .RawParams "{}")}}
**RawParams**:

```json
{{ .RawParams | html}}
```

{{ end}}


**ContentType**: `{{ .Body.ContentType}}`


**Request Body**:
{{- if or (eq .Body.ContentType "application/json") ( eq .Body.ContentType "application/json; charset=utf-8") }}
```json
{{- if .Body.Body }}
{{ .Body.Body | prettyFormat | html }}
{{- end }}
```
{{ end -}}

{{- if eq .Body.ContentType "multipart/form-data" -}}
    <table>
    <tr>
    <th>Key</th>
    <th>Type</th>
    <th>Value</th>
    </tr>
    {{- if isArray .Body.Body }}
        {{- range .Body.Body }}
        <tr>
        <td><code>{{- .key | html }}</code></td>
        <td><code>{{- if .isFile }} @file {{ else }} {{ getDataType .value }} {{- end }}</code></td>
        <td><code>{{- if isHTMLData .value}} `{{- .value | html}}` {{ else }} {{- .value | prettyFormat | html}} {{- end}}</code></td>
        </tr>
        {{- end }}
    {{- end }}
    </table>
{{- end -}}

{{- if eq .Body.ContentType "application/x-www-form-urlencoded" -}}
    <table>
    <tr>
    <th>Key</th>
    <th>Type</th>
    <th>Value</th>
    </tr>
    {{- if isArray .Body.Body }}
        {{- range .Body.Body }}
        <tr>
        <td><code>{{- .key | html }}</code></td>
        <td><code>{{- getDataType .value}}</code></td>
        <td><code>{{- if isHTMLData .value}} `{{- .value | html}}` {{ else }} {{- .value | prettyFormat | html}} {{- end}}</code></td>
        </tr>
        {{- end }}
    {{- end }}
    </table>
{{- end -}}

{{- if (contains .Body.ContentType "xml") -}}
{{- if .Body.Body }}
```xml 
{{ .Body.Body | prettyFormat | html }}
```
{{- end }}
{{- end -}}


<!-- Start Folder/Request Request Variable -->
{{- if getRequestExamples .RequestVariable .ExampleResponses -}}
{{ tabStart | html }}
**Example Response**

{{- range getRequestExamples .RequestVariable .ExampleResponses -}}
#### **{{ .Status }} {{.Name | html}}**

<!-- Level 1 Example Request Variable -->
<h3>Request</h3>
<hr>

{{- if exists (getRequestVariables .OriginalRequest.RequestVariable) }}
<strong>Variable:</strong>
    <table>
        <tr>
        <th>Type</th>
        <th>Value</th> 
        </tr>
        {{- range getRequestVariables .OriginalRequest.RequestVariable -}}
        <tr>
        <td><code>{{ .Key | html}}</code></td>
        <td><code>{{ .Value | html }}</code></td>
        </tr>
        {{- end -}}
    </table>
{{- end }}

<!-- Level 1 Example Request QueryParams -->
{{- if exists (.OriginalRequest.Params) }}
**QueryParams:**
    <table>
    <tr>
    <th>Key</th>
    <th>Type</th>
    <th>Value</th>
    <th>Description</th>
    </tr>
    {{- range .OriginalRequest.Params -}}
    <tr>
    <td>`{{- .key | html}}`</td>
    <td><code>{{- getDataType .value}}</code></td>
    <td>{{- if .value }}`{{ .value | html }}`{{ else }} {{ end }}</td>
    <td>{{- .description | html}}</td>
    </tr>
    {{-  end -}}
    </table>
{{- end }}
    

<!-- Level 1 Example Request QueryParams -->
{{- if exists (.OriginalRequest.Headers) }}
**Header:**
    <table>
    <tr>
    <th>Key</th>
    <th>Value</th>
    <th>Description</th>
    </tr>
    {{- range .OriginalRequest.Headers -}}
    <tr>
    <td>`{{- .Key | html -}}`</td>
    <td>`{{- .Value | html -}}`</td>
    <td>{{- .Description | html -}}</td>
    </tr>
    {{- end -}}
    </table>
{{- end }}

<!-- Level 1 Example Request Body -->
**Body:**
{{ if or (eq .OriginalRequest.Body.ContentType "application/json") ( eq .OriginalRequest.Body.ContentType "application/json; charset=utf-8") }}
```json
{{ if .OriginalRequest.Body.Body }}
{{ .OriginalRequest.Body.Body | prettyFormat | html }}
{{ end }}
```
{{ end }}

{{- if eq .OriginalRequest.Body.ContentType "multipart/form-data" -}}
    <table>
    <tr>
    <th>Key</th>
    <th>Type</th>
    <th>Value</th>
    </tr>
    {{- if isArray .OriginalRequest.Body.Body }}
        {{- range .OriginalRequest.Body.Body }}
        <tr>
        <td><code>{{- .key | html }}</code></td>
        <td><code>{{- if .isFile }} @file {{ else }} {{ getDataType .value }} {{- end }}</code></td>
        <td><code>{{- if isHTMLData .value}} `{{- .value | html}}` {{ else }} {{- .value | prettyFormat | html}} {{- end}}</code></td>
        </tr>
        {{- end }}
    {{- end }}
    </table>
<p></p>
{{- end -}}

{{- if eq .OriginalRequest.Body.ContentType "application/x-www-form-urlencoded" -}}
    <table>
    <tr>
    <th>Key</th>
    <th>Type</th>
    <th>Value</th>
    </tr>
    {{- if isArray .OriginalRequest.Body.Body }}
        {{- range .OriginalRequest.Body.Body }}
        <tr>
        <td><code>{{- .key | html }}</code></td>
        <td><code>{{- getDataType .value}}</code></td>
        <td><code>{{- if isHTMLData .value}} `{{- .value | html}}` {{ else }} {{- .value | prettyFormat | html}} {{- end}}</code></td>
        </tr>
        {{- end }}
    {{- end }}
    </table>
<p></p>
{{- end -}}

{{- if (contains .OriginalRequest.Body.ContentType "xml") -}}
{{- if .OriginalRequest.Body.Body }}
```xml
{{ .OriginalRequest.Body.Body | prettyFormat | html }}
```
{{- end }}
{{- end -}}


<!-- Level 1 Example Response Header -->
<h3>Response</h3>
<hr>

**Header:**
    <table>
        <tr>
            <th>Key</th>
            <th>Value</th>
            <th>Description</th>
        </tr>
        {{- range .Headers -}}
        <tr>
            <td>{{- .Key | html -}}</td>
            <td>{{- .Value | html -}}</td>
            <td>{{- .Description | html -}}</td>
        </tr>
        {{- end -}}
    </table>


<!-- Level 1 Example Response Body -->
**Body:**
```json
{{ .Response | prettyFormat | html }}
```

{{- end -}}
{{ tabEnd | html }}
{{- end -}}
<!-- End Folder/Request Request Variable -->

{{- if .PreRequestScript }}
**Pre Request Script**:

```js
{{ .PreRequestScript | html}}
```
{{- end }}

{{- if .TestScript }}
**Test Script**:

```js
{{ .TestScript | html}}
```
{{- end }}

{{ end }}

---
<!-- End Request Root -->


<!-- Range Folder And Request -->
{{-  if .Folders}}
{{- range .Folders}}

---

## Folder: {{.Name}}
<!-- Start Level 2 Sub Folder Property -->
{{- if ne .Property.Auth.AuthType "" -}}
**Authentication**: `{{ .Property.Auth.AuthType | html }}`
    {{- if eq .Property.Auth.AuthType "bearer" -}}
        <table>
        <tr>
        <th>AddTo</th>
        <th>Key</th>
        <th>Value</th>
        <th>Description</th>
        </tr>
        <tr>
        <td>{{ .Property.Auth.AddTo | html }}</td>
        <td>Authorization</td>
        <td>`{{ .Property.Auth.AuthType | html}}{{" "}}{{- .Property.Auth.Token | html}}`</td>
        <td>{{ .Property.Auth.Description | html}}</td>
        </tr>
        </table>
    {{- end -}}
    {{- if eq .Property.Auth.AuthType "api-key" -}}
        <table>
        <tr>
        <th>AddTo</th>
        <th>Key</th>
        <th>Value</th>
        <th>Description</th>
        </tr>
        <tr>
        <td>{{ .Property.Auth.AddTo | html }}</td>
        <td>`{{ .Property.Auth.Key | html}}`</td>
        <td>`{{- .Property.Auth.Value | html}}`</td>
        <td>{{ .Property.Auth.Description | html}}</td>
        </tr>
        </table>
    {{- end -}}
    {{- if eq .Property.Auth.AuthType "basic" -}}
        <table>
        <tr>
        <th>AddTo</th>
        <th>Username</th>
        <th>Password</th>
        <th>Description</th>
        </tr>
        <tr>
        <td>{{ .Property.Auth.AddTo | html }}</td>
        <td>`{{ .Property.Auth.Username | html}}`</td>
        <td>`{{- .Property.Auth.Password | html}}`</td>
        <td>{{ .Property.Auth.Description | html}}</td>
        </tr>
        </table>
    {{- end -}}
<p></p>
{{- end -}}

{{- if .Property.Headers -}}
**Headers**: 
    <table>
    <tr>
    <th>Key</th>
    <th>Value</th>
    <th>Description</th>
    </tr>
    {{- range .Property.Headers -}}
    <tr>
    <td>`{{- .Key | html -}}`</td>
    <td>`{{- .Value | html -}}`</td>
    <td>{{- .Description | html -}}</td>
    </tr>
    {{- end -}}
    </table>
{{- end -}}
<!-- End Level 2 Sub Folder Property -->

<!-- Start Range Level 3 Folder And Request -->
{{-  if .Folders}}
{{- range .Folders}}

---

### Folder: {{.Name}}
<!-- Start Level 3 Sub Folder Property -->
{{- if ne .Property.Auth.AuthType "" -}}
**Authentication**: `{{ .Property.Auth.AuthType | html }}`
    {{- if eq .Property.Auth.AuthType "bearer" -}}
        <table>
        <tr>
        <th>AddTo</th>
        <th>Key</th>
        <th>Value</th>
        <th>Description</th>
        </tr>
        <tr>
        <td>{{ .Property.Auth.AddTo | html }}</td>
        <td>Authorization</td>
        <td>`{{ .Property.Auth.AuthType | html}}{{" "}}{{- .Property.Auth.Token | html}}`</td>
        <td>{{ .Property.Auth.Description | html}}</td>
        </tr>
        </table>
    {{- end -}}
    {{- if eq .Property.Auth.AuthType "api-key" -}}
        <table>
        <tr>
        <th>AddTo</th>
        <th>Key</th>
        <th>Value</th>
        <th>Description</th>
        </tr>
        <tr>
        <td>{{ .Property.Auth.AddTo | html }}</td>
        <td>`{{ .Property.Auth.Key | html}}`</td>
        <td>`{{- .Property.Auth.Value | html}}`</td>
        <td>{{ .Property.Auth.Description | html}}</td>
        </tr>
        </table>
    {{- end -}}
    {{- if eq .Property.Auth.AuthType "basic" -}}
        <table>
        <tr>
        <th>AddTo</th>
        <th>Username</th>
        <th>Password</th>
        <th>Description</th>
        </tr>
        <tr>
        <td>{{ .Property.Auth.AddTo | html }}</td>
        <td>`{{ .Property.Auth.Username | html}}`</td>
        <td>`{{- .Property.Auth.Password | html}}`</td>
        <td>{{ .Property.Auth.Description | html}}</td>
        </tr>
        </table>
    {{- end -}}
<p></p>
{{- end -}}

{{- if .Property.Headers -}}
**Headers**: 
    <table>
    <tr>
    <th>Key</th>
    <th>Value</th>
    <th>Description</th>
    </tr>
    {{- range .Property.Headers -}}
    <tr>
    <td>`{{- .Key | html -}}`</td>
    <td>`{{- .Value | html -}}`</td>
    <td>{{ .Description | html}}</td>
    </tr>
    {{- end -}}
    </table>
{{- end -}}
<!-- End Level 3 Sub Folder Property -->

<!-- Start Level 3 Request -->
{{-  range .Requests -}}


---

#### [{{ .Method }}] {{.Name}} :id={{ getSlug . }}

**Method**: {{.Method}}

**RequestURL**: `{{ .URL | html }}{{ .Path }}`

{{if .Headers}}
**Headers**: 
    <table>
    <tr>
    <th>Key</th>
    <th>Value</th>
    <th>Description</th>
    </tr>
    {{- range .Headers}}
    <tr>
    <td>{{- .Key | html}}</td>
    <td>`{{- .Value | html}}`</td>
    <td>{{- .Description | html}}</td>
    </tr>
    {{- end -}}
    </table>
{{- end}}
{{- if .Params}}
**Query Params**: 
    <table>
    <tr>
    <th>Key</th>
    <th>Type</th>
    <th>Value</th>
    <th>Description</th>
    </tr>
    {{- range .Params}}
    <tr>
    <td>`{{- .key | html}}`</td>
    <td><code>{{- getDataType .value}}</code></td>
    <td>{{- if .value }}`{{ .value | html }}`{{ else }} {{ end }}</td>
    <td>{{- .description | html}}</td>
    </tr>
    {{-  end}}
    </table>
{{- end -}}

{{ if ne .Auth.AuthType "None"}}
**Authentication Type**: `{{.Auth.AuthType | html}}` 
    
{{ if .Auth.Token}}
**BearerToken**: `{{ .Auth.Token | html}}`

{{ end}}

{{- if and .User .Pass}}

**Username**: `{{  .User}}`
**Password**: `{{  .Pass}}`

{{ end -}}

{{ end}}

{{- if and .RawParams (ne .RawParams "{}")}}
**RawParams:**

```json
{{ .RawParams | html}}
```
{{ end -}}


**ContentType**: `{{ .Body.ContentType}}` 

{{- if getRequestVariables .RequestVariable }}
<p><strong>Request Variable:</strong></p>
<table>
    <tr>
    <th>Type</th>
    <th>Value</th>
    </tr>
    {{- range getRequestVariables .RequestVariable }}
    <tr>
    <td><code>{{ .Key | html}}</code></td>
    <td><code>{{ .Value | html }}</code></td>
    </tr>
    {{- end }}
</table>
{{- end }}

**Request Body**:
{{ if or (eq .Body.ContentType "application/json") ( eq .Body.ContentType "application/json; charset=utf-8") }}
```json
{{ if .Body.Body }}
{{ .Body.Body | prettyFormat | html }}
{{ end }}
```
{{ end }}

{{- if eq .Body.ContentType "multipart/form-data" -}}
    <table>
    <tr>
    <th>Key</th>
    <th>Type</th>
    <th>Value</th>
    </tr>
    {{- if isArray .Body.Body }}
        {{- range .Body.Body }}
        <tr>
        <td><code>{{- .key | html }}</code></td>
        <td><code>{{- if .isFile }} @file {{ else }} {{ getDataType .value }} {{- end }}</code></td>
        <td><code>{{- if isHTMLData .value}} `{{- .value | html}}` {{ else }} {{- .value | prettyFormat | html}} {{- end}}</code></td>
        </tr>
        {{- end }}
    {{- end }}
    </table>
<p></p>
{{- end -}}

{{- if eq .Body.ContentType "application/x-www-form-urlencoded" -}}
    <table>
    <tr>
    <th>Key</th>
    <th>Type</th>
    <th>Value</th>
    </tr>
    {{- if isArray .Body.Body }}
        {{- range .Body.Body }}
        <tr>
        <td><code>{{- .key | html }}</code></td>
        <td><code>{{- getDataType .value}}</code></td>
        <td><code>{{- if isHTMLData .value}} `{{- .value | html}}` {{ else }} {{- .value | prettyFormat | html}} {{- end}}</code></td>
        </tr>
        {{- end }}
    {{- end }}
    </table>
<p></p>
{{- end -}}

{{- if (contains .Body.ContentType "xml") -}}
{{- if .Body.Body }}
```xml
{{ .Body.Body | prettyFormat | html }}
```
{{- end }}
{{- end -}}


<!-- Start Folder/Request Request Variable -->
{{- if getRequestExamples .RequestVariable .ExampleResponses -}}
{{ tabStart | html }}

**Example Response**

{{- range getRequestExamples .RequestVariable .ExampleResponses -}}
#### **{{ .Status }} {{.Name | html}}**


<!-- Level 3 Example Request Variable -->
<h3>Request</h3>
<hr>

{{- if exists (getRequestVariables .OriginalRequest.RequestVariable) }}
<strong>Variable:</strong>
    <table>
        <tr>
        <th>Type</th>
        <th>Value</th> 
        </tr>
        {{- range getRequestVariables .OriginalRequest.RequestVariable -}}
        <tr>
        <td><code>{{ .Key | html}}</code></td>
        <td><code>{{ .Value | html }}</code></td>
        </tr>
        {{- end -}}
    </table>
{{- end }}

<!-- Level 3 Example Request QueryParams -->
{{- if exists (.OriginalRequest.Params) }}
**QueryParams:**
    <table>
    <tr>
    <th>Key</th>
    <th>Type</th>
    <th>Value</th>
    <th>Description</th>
    </tr>
    {{- range .OriginalRequest.Params -}}
    <tr>
    <td>`{{- .key | html}}`</td>
    <td><code>{{- getDataType .value}}</code></td>
    <td>{{- if .value }}`{{ .value | html }}`{{ else }} {{ end }}</td>
    <td>{{- .description | html}}</td>
    </tr>
    {{-  end -}}
    </table>
{{- end }}
    

<!-- Level 3 Example Request QueryParams -->
{{- if exists (.OriginalRequest.Headers) }}
**Header:**
    <table>
    <tr>
    <th>Key</th>
    <th>Value</th>
    <th>Description</th>
    </tr>
    {{- range .OriginalRequest.Headers -}}
    <tr>
    <td>`{{- .Key | html -}}`</td>
    <td>`{{- .Value | html -}}`</td>
    <td>{{- .Description | html -}}</td>
    </tr>
    {{- end -}}
    </table>
{{- end }}

<!-- Level 3 Example Request Body -->
**Body:**
{{ if or (eq .OriginalRequest.Body.ContentType "application/json") ( eq .OriginalRequest.Body.ContentType "application/json; charset=utf-8") }}
```json
{{ if .OriginalRequest.Body.Body }}
{{ .OriginalRequest.Body.Body | prettyFormat | html }}
{{ end }}
```
{{ end }}

{{- if eq .OriginalRequest.Body.ContentType "multipart/form-data" -}}
    <table>
    <tr>
    <th>Key</th>
    <th>Type</th>
    <th>Value</th>
    </tr>
    {{- if isArray .OriginalRequest.Body.Body }}
        {{- range .OriginalRequest.Body.Body }}
        <tr>
        <td><code>{{- .key | html }}</code></td>
        <td><code>{{- if .isFile }} @file {{ else }} {{ getDataType .value }} {{- end }}</code></td>
        <td><code>{{- if isHTMLData .value}} `{{- .value | html}}` {{ else }} {{- .value | prettyFormat | html}} {{- end}}</code></td>
        </tr>
        {{- end }}
    {{- end }}
    </table>
<p></p>
{{- end -}}

{{- if eq .OriginalRequest.Body.ContentType "application/x-www-form-urlencoded" -}}
    <table>
    <tr>
    <th>Key</th>
    <th>Type</th>
    <th>Value</th>
    </tr>
    {{- if isArray .OriginalRequest.Body.Body }}
        {{- range .OriginalRequest.Body.Body }}
        <tr>
        <td><code>{{- .key | html }}</code></td>
        <td><code>{{- getDataType .value}}</code></td>
        <td><code>{{- if isHTMLData .value}} `{{- .value | html}}` {{ else }} {{- .value | prettyFormat | html}} {{- end}}</code></td>
        </tr>
        {{- end }}
    {{- end }}
    </table>
<p></p>
{{- end -}}

{{- if (contains .OriginalRequest.Body.ContentType "xml") -}}
{{- if .OriginalRequest.Body.Body }}
```xml
{{ .OriginalRequest.Body.Body | prettyFormat | html }}
```
{{- end }}
{{- end -}}



<!-- Level 3 Example Response Header -->
<h3>Response</h3>
<hr>

**Header:**
    <table>
        <tr>
            <th>Key</th>
            <th>Value</th>
            <th>Description</th>
        </tr>
        {{- range .Headers -}}
        <tr>
            <td>{{- .Key | html -}}</td>
            <td>{{- .Value | html -}}</td>
            <td>{{- .Description | html -}}</td>
        </tr>
        {{- end -}}
    </table>


<!-- Level 3 Example Response Body -->
**Body:**
```json
{{ .Response | prettyFormat | html }}
```

{{- end -}}
{{ tabEnd | html }}
{{- end -}}
<!-- End Folder/Request Request Variable -->


{{ if .PreRequestScript }}
**Pre Request Script**:

```js
{{ .PreRequestScript | html}}
```
{{ end }}

{{ if .TestScript }}
**Test Script**:

```js
{{ .TestScript | html}}
```
{{ end }}


{{- end -}}
{{- end }} {{- end }}
<!-- End Level 3 Request -->

<!-- Start Level 2 Request -->
{{-  range .Requests }}

---

### [{{ .Method }}] {{.Name}} :id={{ getSlug . }}

**Method**: {{.Method}}

**RequestURL**: `{{ .URL | html }}{{ .Path }}`

{{if .Headers}}
**Headers**: 
    <table>
    <tr>
    <th>Key</th>
    <th>Value</th>
    <th>Description</th>
    </tr>
    {{- range .Headers}}
    <tr>
    <td>{{- .Key | html}}</td>
    <td>`{{- .Value | html}}`</td>
    <td>{{- .Description | html}}</td>
    </tr>
    {{- end -}}
    </table>
{{- end}}
{{- if .Params}}
**Query Params**: 
    <table>
    <tr>
    <th>Key</th>
    <th>Type</th>
    <th>Value</th>
    <th>Description</th>
    </tr>
    {{- range .Params}}
    <tr>
    <td>`{{- .key | html}}`</td>
    <td><code>{{- getDataType .value}}</code></td>
    <td>{{- if .value }}`{{ .value | html }}`{{ else }} {{ end }}</td>
    <td>{{- .description | html}}</td>
    </tr>
    {{-  end}}
    </table>
{{- end -}}

{{ if ne .Auth.AuthType "None"}}
**Authentication Type**: `{{.Auth.AuthType | html}}` 
    
{{ if .Auth.Token}}
**BearerToken**: `{{ .Auth.Token | html}}`

{{ end}}

{{- if and .User .Pass}}

**Username**: `{{  .User}}`
**Password**: `{{  .Pass}}`

{{ end -}}

{{ end }}

{{- if and .RawParams (ne .RawParams "{}")}}
**RawParams:**

```json
{{ .RawParams | html}}
```
{{ end -}}

**ContentType**: `{{ .Body.ContentType}}` 

{{- if getRequestVariables .RequestVariable }}
<p><strong>Request Variable:</strong></p>
<table>
    <tr>
    <th>Type</th>
    <th>Value</th>
    </tr>
    {{- range getRequestVariables .RequestVariable }}
    <tr>
    <td><code>{{ .Key | html}}</code></td>
    <td><code>{{ .Value | html }}</code></td>
    </tr>
    {{- end }}
</table>
{{- end }}

**Request Body**:
{{- if or (eq .Body.ContentType "application/json") ( eq .Body.ContentType "application/json; charset=utf-8") }}
```json
{{- if .Body.Body }}
{{ .Body.Body | prettyFormat | html }}
{{- end }}
```
{{ end -}}

{{- if eq .Body.ContentType "multipart/form-data" -}}
    <table>
    <tr>
    <th>Key</th>
    <th>Type</th>
    <th>Value</th>
    </tr>
    {{- if isArray .Body.Body }}
        {{- range .Body.Body }}
        <tr>
        <td><code>{{- .key | html }}</code></td>
        <td><code>{{- if .isFile }} @file {{ else }} {{ getDataType .value }} {{- end }}</code></td>
        <td><code>{{- if isHTMLData .value}} `{{- .value | html}}` {{ else }} {{- .value | prettyFormat | html}} {{- end}}</code></td>
        </tr>
        {{- end }}
    {{- end }}
    </table>
{{- end -}}

{{- if eq .Body.ContentType "application/x-www-form-urlencoded" -}}
    <table>
    <tr>
    <th>Key</th>
    <th>Type</th>
    <th>Value</th>
    </tr>
    {{- if isArray .Body.Body }}
        {{- range .Body.Body }}
        <tr>
        <td><code>{{- .key | html }}</code></td>
        <td><code>{{- getDataType .value}}</code></td>
        <td><code>{{- if isHTMLData .value}} `{{- .value | html}}` {{ else }} {{- .value | prettyFormat | html}} {{- end}}</code></td>
        </tr>
        {{- end }}
    {{- end }}
    </table>
{{- end -}}

{{- if (contains .Body.ContentType "xml") -}}
{{- if .Body.Body }}
```xml
{{ .Body.Body | prettyFormat | html }}
```
{{- end }}
{{- end -}}

<!-- Start Folder/Request Request Variable -->
{{- if getRequestExamples .RequestVariable .ExampleResponses -}}
{{ tabStart | html }}
**Example Response**

{{- range getRequestExamples .RequestVariable .ExampleResponses -}}
#### **{{ .Status }} {{.Name | html}}**

<!-- Level 2 Example Request Variable -->
<h3>Request</h3>
<hr>

{{- if exists (getRequestVariables .OriginalRequest.RequestVariable) }}
<strong>Variable:</strong>
    <table>
        <tr>
        <th>Type</th>
        <th>Value</th> 
        </tr>
        {{- range getRequestVariables .OriginalRequest.RequestVariable -}}
        <tr>
        <td><code>{{ .Key | html}}</code></td>
        <td><code>{{ .Value | html }}</code></td>
        </tr>
        {{- end -}}
    </table>
{{- end }}

<!-- Level 2 Example Request QueryParams -->
{{- if exists (.OriginalRequest.Params) }}
**QueryParams:**
    <table>
    <tr>
    <th>Key</th>
    <th>Type</th>
    <th>Value</th>
    <th>Description</th>
    </tr>
    {{- range .OriginalRequest.Params -}}
    <tr>
    <td>`{{- .key | html}}`</td>
    <td><code>{{- getDataType .value}}</code></td>
    <td>{{- if .value }}`{{ .value | html }}`{{ else }} {{ end }}</td>
    <td>{{- .description | html}}</td>
    </tr>
    {{-  end -}}
    </table>
{{- end }}
    

<!-- Level 2 Example Request QueryParams -->
{{- if exists (.OriginalRequest.Headers) }}
**Header:**
    <table>
    <tr>
    <th>Key</th>
    <th>Value</th>
    <th>Description</th>
    </tr>
    {{- range .OriginalRequest.Headers -}}
    <tr>
    <td>`{{- .Key | html -}}`</td>
    <td>`{{- .Value | html -}}`</td>
    <td>{{- .Description | html -}}</td>
    </tr>
    {{- end -}}
    </table>
{{- end }}

<!-- Level 2 Example Request Body -->
**Body:**
{{ if or (eq .OriginalRequest.Body.ContentType "application/json") ( eq .OriginalRequest.Body.ContentType "application/json; charset=utf-8") }}
```json
{{ if .OriginalRequest.Body.Body }}
{{ .OriginalRequest.Body.Body | prettyFormat | html }}
{{ end }}
```
{{ end }}

{{- if eq .OriginalRequest.Body.ContentType "multipart/form-data" -}}
    <table>
    <tr>
    <th>Key</th>
    <th>Type</th>
    <th>Value</th>
    </tr>
    {{- if isArray .OriginalRequest.Body.Body }}
        {{- range .OriginalRequest.Body.Body }}
        <tr>
        <td><code>{{- .key | html }}</code></td>
        <td><code>{{- if .isFile }} @file {{ else }} {{ getDataType .value }} {{- end }}</code></td>
        <td><code>{{- if isHTMLData .value}} `{{- .value | html}}` {{ else }} {{- .value | prettyFormat | html}} {{- end}}</code></td>
        </tr>
        {{- end }}
    {{- end }}
    </table>
<p></p>
{{- end -}}

{{- if eq .OriginalRequest.Body.ContentType "application/x-www-form-urlencoded" -}}
    <table>
    <tr>
    <th>Key</th>
    <th>Type</th>
    <th>Value</th>
    </tr>
    {{- if isArray .OriginalRequest.Body.Body }}
        {{- range .OriginalRequest.Body.Body }}
        <tr>
        <td><code>{{- .key | html }}</code></td>
        <td><code>{{- getDataType .value}}</code></td>
        <td><code>{{- if isHTMLData .value}} `{{- .value | html}}` {{ else }} {{- .value | prettyFormat | html}} {{- end}}</code></td>
        </tr>
        {{- end }}
    {{- end }}
    </table>
<p></p>
{{- end -}}

{{- if (contains .OriginalRequest.Body.ContentType "xml") -}}
{{- if .OriginalRequest.Body.Body }}
```xml
{{ .OriginalRequest.Body.Body | prettyFormat | html }}
```
{{- end }}
{{- end -}}



<!-- Level 2 Example Response Header -->
<h3>Response</h3>
<hr>

**Header:**
    <table>
        <tr>
            <th>Key</th>
            <th>Value</th>
            <th>Description</th>
        </tr>
        {{- range .Headers -}}
        <tr>
            <td>{{- .Key | html -}}</td>
            <td>{{- .Value | html -}}</td>
            <td>{{- .Description | html -}}</td>
        </tr>
        {{- end -}}
    </table>


<!-- Level 2 Example Response Body -->
**Body:**
```json
{{ .Response | prettyFormat | html }}
```

{{- end -}}
{{ tabEnd | html }}
{{- end -}}
<!-- End Folder/Request Request Variable -->


{{- if .PreRequestScript }}
**Pre Request Script**:

```js
{{ .PreRequestScript | html}}
```
{{- end }}

{{- if .TestScript }}
**Test Script**:

```js
{{ .TestScript | html}}
```
{{- end }}

{{ end -}}{{ end -}}

<!-- End Level 2 Request Root -->

---
{{end}}{{end}}

<!-- End Request Root -->
