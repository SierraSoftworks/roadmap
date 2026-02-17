{{- /*gotype: github.com/SierraSoftworks/roadmap.Roadmap*/ -}}
{{ $root := . -}}
---
{{- range $key, $value := metadata }}
{{ $key }}: {{ $value | json }}
{{- end }}
---

# {{ .Title }}

<div style="display: flex; flex-direction: row; background: rgba(200, 200, 200, 0.15); padding: 1rem; border-radius: 6px; margin-bottom: 2rem;">
<div style="vertical-align: middle; font-size: 0.7rem; line-height: 1.8rem; opacity: 0.6;">Authored by</div>
{{ range .Authors }}
<div style="margin-left: 1rem; padding-left: 1rem; border-left: 1px solid rgba(200, 200, 200, 0.5);">
    <h5 style="font-size: 0.8rem; margin: 0;">{{ .Name }}</h5>
    <p style="font-size: 0.6rem; margin: 0;">{{ .Contact }}</p>
</div>
{{ end }}
</div>


{{ .Description }}

## Important Dates

<div style="border-left: 4px solid gray; border-radius: 0 4px 4px 0; background: rgba(200, 200, 200, 0.15); padding: 1rem 2rem; position: relative; text-align: center; margin: 2rem auto 2rem 7rem;">
{{ range .Timeline }}
<div style="text-align: left; position: relative; padding-bottom: 1rem; margin-bottom: 1rem;">
<div style="position: absolute; left: -10rem; text-align: right; font-size: 0.9rem; font-weight: 700; opacity: 0.7; min-width: 6rem; top: 2px;">{{ .Date.Format "2006-01-02" }}</div>

<h3>{{ .Title }}</h3>
{{ .Description | markdown }}

<div style="position: absolute; box-shadow: 0 0 0 4px gray; left: -2.5rem; background: #444; border-radius: 50%; height: 11px; width: 11px; top: 5px;"></div>
</div>
{{ end }}
</div>

## Objectives
{{ range .Objectives }}
### {{ .Title }}
{{ .Description}}

{{ end }}

## Milestones

<div style="border-left: 4px solid gray; border-radius: 0 4px 4px 0; background: rgba(200, 200, 200, 0.15); padding: 1rem 2rem; position: relative; text-align: center; margin: 2rem auto 2rem 7rem;">
{{ range $i,$m := .Milestones }}
<div style="text-align: left; position: relative; padding-bottom: 1rem; margin-bottom: 1rem;">
<div style="position: absolute; left: -10rem; text-align: right; font-size: 0.9rem; font-weight: 700; opacity: 0.7; min-width: 6rem; top: 2px;">M{{ add $i 1 }}</div>

<h3>{{ $m.Title }}</h3>
{{ $m.Description | markdown }}

{{ with $m.Reference }}<a href="{{ . }}">Read more &rarr;</a>{{ end }}

{{ range $m.Deliverables }}
<div style="position: relative; border-radius: 4px; box-shadow: 2px 2px 10px rgba(0,0,0,0.3); background-color: rgba(0, 0, 0, 0.1); margin: 2rem 0; padding: 10px 20px;">
<div style="position: absolute; top: 0; left: 0; bottom: 0; width: 8px; border-radius: 4px 0 0 4px; background-color: {{ .State | stateColor }}"></div>
<h4 style="margin-top: 0; line-height: 1.8">
<span style="float: right; margin: 0;">{{ .State }}</span>

{{ .Title }}
{{ with .Requirement }}<span style="display: inline; font-size: 90%; padding: 3px 5px; border-radius: 4px; background-color: {{ . | requirementColor }}; color: white; margin: 0 2px;"> {{ . }}</span>{{ end }}
</h4>

{{ .Description | markdown }}

{{ with .Reference }}<a href="{{ . }}">Read more &rarr;</a>{{ end }}
</div>
{{ end }}

<div style="position: absolute; box-shadow: 0 0 0 4px gray; left: -2.5rem; background: #444; border-radius: 50%; height: 11px; width: 11px; top: 5px;"></div>
</div>
{{ end }}
</div>
