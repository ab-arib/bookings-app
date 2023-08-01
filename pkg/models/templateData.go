package models

// template data holds data sent from handlers to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int32
	FloatMap  map[string]float32
	Data      map[string]interface{} //interface use for non specific data type
	CSFRToken string
	Flash     string
	Warning   string
	Error     string
}
