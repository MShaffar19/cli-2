if [ -f ~/.bashrc ]; then source ~/.bashrc; fi

if [ -z "$PROMPT_COMMAND" ]; then
  export PS1="[{{.Owner}}/{{.Name}}] $PS1"
fi

{{- range $K, $V := .Env}}
{{- if eq $K "PATH"}}
export {{$K}}="{{$V}}:$PATH"
{{- else}}
export {{$K}}="{{$V}}"
{{- end}}
{{- end}}

{{range $K, $CMD := .Scripts}}
alias {{$K}}='state run {{$CMD}}'
{{end}}

cd "{{.WD}}"

{{.UserScripts}}
