{{ range .  }}

# {{.Name}}

{{-  if .Folders}}
{{- range .Folders}}

---

## Folder: {{.Name}}
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

**Authentication Type**:  {{ .Auth.AuthType}}

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

**SubContentType**: `{{ .Body.ContentType}}` 

**Request Body**:
{{- if or (eq .Body.ContentType "application/json") ( eq .Body.ContentType "application/json; charset=utf-8") }}
```json
{{ if .Body.Body }}
    {{ .Body.Body | html }}
{{ else }}
    
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
**Authentication Type**: {{.Auth.AuthType}}  

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
{{ else }}
    
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

**Example Response**:
{{- if .RequestVariable }}
    {{- range .RequestVariable }}
        {{- if .Examples }}
            {{- range .Examples }}
                {{.Status}}
                {{.Name}}
                {{.Response}}
            {{- end }}
        {{- end }}
    {{- end}}
{{- end }}

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