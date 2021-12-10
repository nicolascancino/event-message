package dto

// Out struct is the structure to be logged
type Out struct {
	ID          string
	Attributes  map[string]string
	Data        map[string]interface{}
	PublishTime string
}
