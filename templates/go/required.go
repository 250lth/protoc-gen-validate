package golang

const requiredTpl = `
	{{ if .Rules.GetRequired }}
		if {{ accessor . }} == nil {
			err := {{ err . "value is required" }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ end }}
`

const testTpl = `{{ $f := .Field }}{{ $r := .Rules }}
	{{ template "required" . }}

	!!!!!! 1111111111111
	{{ $f }}
`
