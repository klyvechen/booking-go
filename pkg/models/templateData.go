package models

// TemplateData holds data to
type TemplateData struct {
	StringMap map[string]string
	IntMAp    map[string]string
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Waring    string
	Error     string
}
