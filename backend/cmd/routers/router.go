// D:/GoProjects/UsmCallCenter/main/app/routers/router.go
package routers

import (
	"html/template"
)

var tpl = template.Must(template.ParseFiles("./backend/cmd/static/income-calls.html"))
