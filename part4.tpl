{{template "base.tpl" .}}

{{define "content"}}
<h1>exercise_part4</h1>
{{if .Message}}
  <ul>
  {{range $message := .Messages_line}}
    <li> {{ $message.title }} {{ $message.message }} ({{ $message.created_at }})</li>
  {{end}}
  </ul>
  {{end}}
</form>
{{end}}
