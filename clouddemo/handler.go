package function

import (
	"bytes"
	"fmt"
	"html/template"
)

type HomepageTokens struct {
	Dark bool
}

// Handle a serverless request
func Handle(req []byte) string {
	dark := true

	var err error
	tmpl, err := template.ParseFiles("./template/index.html")
	if err != nil {
		return fmt.Sprintf("Internal server error with homepage: %s", err.Error())
	}
	var tpl bytes.Buffer

	err = tmpl.Execute(&tpl, HomepageTokens{
		Dark: dark,
	})
	if err != nil {
		return fmt.Sprintf("Internal server error with homepage template: %s", err.Error())
	}

	return string(tpl.Bytes())

}
