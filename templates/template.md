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
        </tr>
        <tr>
        <td>{{ .Property.Auth.AddTo | html }}</td>
        <td>Authorization</td>
        <td>`{{ .Property.Auth.AuthType | html}}{{" "}}{{- .Property.Auth.Token | html}}`</td>
        </tr>
        </table>
    {{- end -}}
    {{- if eq .Property.Auth.AuthType "api-key" -}}
        <table>
        <tr>
        <th>AddTo</th>
        <th>Key</th>
        <th>Value</th>
        </tr>
        <tr>
        <td>{{ .Property.Auth.AddTo | html }}</td>
        <td>`{{ .Property.Auth.Key | html}}`</td>
        <td>`{{- .Property.Auth.Value | html}}`</td>
        </tr>
        </table>
    {{- end -}}
    {{- if eq .Property.Auth.AuthType "basic" -}}
        <table>
        <tr>
        <th>AddTo</th>
        <th>Username</th>
        <th>Password</th>
        </tr>
        <tr>
        <td>{{ .Property.Auth.AddTo | html }}</td>
        <td>`{{ .Property.Auth.Username | html}}`</td>
        <td>`{{- .Property.Auth.Password | html}}`</td>
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
    </tr>
    {{- range .Property.Headers -}}
    <tr>
    <td>{{- .Key | html -}}</td>
    <td>`{{- .Value | html -}}`</td>
    </tr>
    {{- end -}}
    </table>
{{- end -}}
<!-- END Root Folder Property -->


<!-- Range Folder And Request -->
{{-  if .Folders}}
{{- range .Folders}}

---

## Folder: {{.Name}}
<!-- Start Sub Folder Property -->
{{- if ne .Property.Auth.AuthType "" -}}
**Authentication**: `{{ .Property.Auth.AuthType | html }}`
    {{- if eq .Property.Auth.AuthType "bearer" -}}
        <table>
        <tr>
        <th>AddTo</th>
        <th>Key</th>
        <th>Value</th>
        </tr>
        <tr>
        <td>{{ .Property.Auth.AddTo | html }}</td>
        <td>Authorization</td>
        <td>`{{ .Property.Auth.AuthType | html}}{{" "}}{{- .Property.Auth.Token | html}}`</td>
        </tr>
        </table>
    {{- end -}}
    {{- if eq .Property.Auth.AuthType "api-key" -}}
        <table>
        <tr>
        <th>AddTo</th>
        <th>Key</th>
        <th>Value</th>
        </tr>
        <tr>
        <td>{{ .Property.Auth.AddTo | html }}</td>
        <td>`{{ .Property.Auth.Key | html}}`</td>
        <td>`{{- .Property.Auth.Value | html}}`</td>
        </tr>
        </table>
    {{- end -}}
    {{- if eq .Property.Auth.AuthType "basic" -}}
        <table>
        <tr>
        <th>AddTo</th>
        <th>Username</th>
        <th>Password</th>
        </tr>
        <tr>
        <td>{{ .Property.Auth.AddTo | html }}</td>
        <td>`{{ .Property.Auth.Username | html}}`</td>
        <td>`{{- .Property.Auth.Password | html}}`</td>
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
    </tr>
    {{- range .Property.Headers -}}
    <tr>
    <td>{{- .Key | html -}}</td>
    <td>`{{- .Value | html -}}`</td>
    </tr>
    {{- end -}}
    </table>
{{- end -}}
<!-- End Sub Folder Property -->


{{-  range .Requests }}


---

### [{{ .Method }}] {{.Name}}

**Method**: {{.Method}}

**RequestURL**: `{{ .URL | html }}{{ .Path }}`

{{if .Headers}}
**Headers**: 
    <table>
    <tr>
    <th>Key</th>
    <th>Value</th>
    </tr>
    {{- range .Headers}}
    <tr>
    <td>{{- .Key | html}}</td>
    <td>`{{- .Value | html}}`</td>
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
    </tr>
    {{- range .Params}}
    <tr>
    <td>`{{- .key | html}}`</td>
    <td><code>{{- getDataType .value}}</code></td>
    <td>{{- if .value }}`{{ .value | html }}`{{ else }} {{ end }}</td>
    </tr>
    {{-  end}}
    </table>
{{- end -}}

{{- if ne .Auth.AuthType "None" }}

**Authentication Type**:  `{{ .Auth.AuthType}}`

{{- if .Auth.Token}}

**BearerToken**:  `{{  .Auth.Token | html -}} `

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

**Request Body**:
{{- if or (eq .Body.ContentType "application/json") ( eq .Body.ContentType "application/json; charset=utf-8") }}
```json
{{ if .Body.Body }}
    {{ .Body.Body | html }}
{{ end }}
```
{{- end }}

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
        <td><code>{{- .value | html}}</code></td>
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
        <td><code>{{- .value | html}}</code></td>
        </tr>
        {{- end }}
    {{- end }}
    </table>
{{- end -}}

<!-- Start Folder/Request Request Variable -->
{{- if getRequestExamples .RequestVariable -}}
{{ tabStart | html }}
**Example Response**

{{- range getRequestExamples .RequestVariable -}}
#### **{{ .Status }} {{.Name | html}}**
**Response Header** 


**Response Body**
```json
{{ .Response | html }}
```

{{- end -}}
{{ tabEnd | html }}
{{- end -}}
<!-- End Folder/Request Request Variable -->


<!-- 
**Pre Request Script**:

```js
{{ .PreRequestScript | html}}
```

**Test Script**:

```js
{{ .TestScript | html}}
``` -->

{{ end -}}{{ end -}}{{ else}}
{{- range .Requests}}

---

### [{{ .Method }}] {{ .Name}}

**Method**: {{ .Method}}

**RequestURL**:  `{{ .URL | html}}{{ .Path}}`

{{ if .Headers}}
**Headers**: 
    <table>
    <tr>
    <th>Key</th>
    <th>Value</th>
    </tr>
    {{- range .Headers}}
    <tr>
    <td>{{- .Key | html}}</td>
    <td>`{{- .Value | html }}`</td>
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
    </tr>
    {{- range .Params}}
    <tr>
    <td>`{{- .key | html}}`</td>
    <td><code>{{- getDataType .value}}</code></td>
    <td>`{{- if .value }}{{ .value | html }}`{{ else }} {{ end }}</td>
    </tr>
    {{-  end}}
    </table>
{{- end -}}
{{ if ne .Auth.AuthType "None"}}
**Authentication Type**: `{{.Auth.AuthType}}`

{{ if .Auth.Token}}**BearerToken**: `{{ .Auth.Token | html}}`{{ end}}

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
{{ if .Body.Body }}
    {{ .Body.Body | html }}
{{ end }}
```
{{- end }}

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
        <td><code>{{- .value | html}}</code></td>
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
        <td><code>{{- .value | html}}</code></td>
        </tr>
        {{- end }}
    {{- end }}
    </table>
{{- end -}}

<!-- Start Folder/Request Request Variable -->
{{- if getRequestExamples .RequestVariable -}}
{{ tabStart | html }}
**Example Response**

{{- range getRequestExamples .RequestVariable -}}
#### **{{ .Status }} {{.Name | html}}**
**Response Header** 


**Response Body**
```json
{{ .Response | html }}
```

{{- end -}}
{{ tabEnd | html }}
{{- end -}}
<!-- End Folder/Request Request Variable -->

<!-- 
**Pre Request Script**: 

```js
{{ .PreRequestScript | html}}
```

**Test Script**: 

```js
{{ .TestScript | html}}
``` -->

{{ end }}{{ end}}{{ end}}

---

Generated by [HoppScotch CLI](https://github.com/BlackMocca/hopp-cli) ❤️