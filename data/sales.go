
package data

type Sale struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	State  string `json:"type"`
	Window Window `json:"window"`
}
type Window struct {
	Type   string `json:"type"`
	Height string `json:"height"`
	Width  string `json:"width"`
}
type Sales []Sale

var SalesStore = Sales{
	Sale {
		ID: "1", 
		Name: "Dhaka", 
		State: "Uttara", 
		Window: Window{
			Type: "Door", 
			Height: "20", 
			Width: "10",
		},
	},
}