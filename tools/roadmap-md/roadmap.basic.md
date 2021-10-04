{{ $root := . -}}

# {{ .Title }}
{{ .Description }}

## Authors
{{ range .Authors }}
- {{ .Name }} (*{{ .Contact }}*)
{{ end }}

## Important Dates
{{ range .Timeline }}
### **{{ .Date.Format "2006-01-02" }}** | {{ .Title }}
{{ .Description }}
{{ end }}


## Milestones
{{ range $i,$m := .Milestones }}
### **M{{ add $i 1 }}** | {{ $m.Title }}
{{ $m.Description }}

{{ range $di, $d := $m.Deliverables }}
#### **{{ .State }}** | {{ with $d.Reference }}[{{ $d.Title }}]({{ . }}){{ else }}{{ $d.Title }}{{ end }}
{{ .Description }}
{{ end }}
{{ end }}
