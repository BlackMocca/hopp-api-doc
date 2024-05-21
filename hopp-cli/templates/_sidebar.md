{{ range .  }}
<!-- Root -->
- [**{{.Name}}**](/?id={{getSlug . }})

  <!-- Sub Level 2 Folder -->
  {{if .Folders }}
  {{range .Folders}}

  - **{{.Name}}**
    <!-- Sub Level 3 Folder -->
    {{- if .Folders }}
    {{- range .Folders }}

    - **{{.Name}}**
      <!-- Sub Level 3 Request-->
      {{- if .Requests -}}
      {{- range .Requests }}

        - [[{{ .Method }}] {{.Name}}](/?id={{getSlug . }})

      {{ end -}}
      {{- end -}}
    
    {{ end -}}
    {{ end -}}


    <!-- Sub Level 2 Request-->
    {{- if .Requests }}
    {{ range .Requests}}

    - [[{{ .Method }}] {{.Name}}](/?id={{getSlug . }})

    {{ end }}
    {{- end }}
    

  {{ end }}
  {{ end }}

  <!-- Root Request-->
  {{if .Requests }}
  {{range .Requests}}

  - [[{{ .Method }}] {{.Name}}](/?id={{getSlug . }})

  {{ end }}
  {{ end }}

{{ end }}