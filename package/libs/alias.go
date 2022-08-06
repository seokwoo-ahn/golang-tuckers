package libs

import (
	"fmt"
	htemplate "html/template"
	ttemplate "text/template"
)

func Alias() {
	a, err1 := ttemplate.New("foo").Parse(`{{define "T"}}Hello`)
	b, err2 := htemplate.New("foo").Parse(`{{define "T"}}Hello`)
	fmt.Println(a, err1)
	fmt.Println(b, err2)
}
