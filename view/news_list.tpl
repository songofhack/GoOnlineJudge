{{define "content"}}
<h2 class="page-header">News List</h2>
{{with .News}}

	{{range .}}
		{{if ShowStatus .Status}}

			<p class="news">
				<span class="flag"></span>
				<span class="date">{{.Create}}</span>		
				<br><a href="/news/detail?nid={{.Nid}}">{{.Title}}</a>
			</p>
		{{end}}
	{{end}}
{{end}}
{{end}}
