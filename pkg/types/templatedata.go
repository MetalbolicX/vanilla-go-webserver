package types

// TemplateData hols data sent from handlers to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float64
	Data      map[string]any
	CSRToken  string
	Flash     string
	Warning   string
	Error     string
}
