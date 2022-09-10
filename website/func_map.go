package website

import "html/template"

var FuncMap = template.FuncMap{
	"getConfig": getConfig,
	"smallImg":  smallImg,
}

func getConfig(s string) string {
	return s
}
func smallImg(s string) string {
	return s
}
