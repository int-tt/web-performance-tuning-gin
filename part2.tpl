{{template "base.tpl" .}}
{{define "post_url"}}
/chapter1/write
{{end}}
{{define "username"}}

<div style="background:#ffeeee; padding:10px; border:1px solid #ff0000; border-radius:20px;"> {{ .User }} </div>
{{end}}
{{define "messages"}}
<a href="/"><span class="badge"> {{ .Message }} </span></a>
{{end}}
{{define "follow"}}

<a href="/"><span class="badge">{{ .Follow }}</span></a>
{{end}}
{{define "follower"}}
<a href="/"><span class="badge">{{ .Follower }}</span></a>
{{end}}
{{define "content"}}
{{if .Messages_line}}
  <ul>
  {{range $message := .Messages_line}}
    <li> {{ $message.message }} ({{ $message.created_at }})</li>
  {{end}}
  </ul>
{{end}}
{{end}}
