package models

//Template Data holds data sent from handlers.go to templates folder
type TemplateData struct {
	StringData map[string]string
	IntData    map[string]int
	FloatData  map[string]float64
	Data       map[string]interface{}
	CSRFToken  string
	Flash      string
	Error      string
}
