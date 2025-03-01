package example

type T struct {
	ResourceType string `json:"resourceType"`
	Id           string `json:"id"`
	Text         struct {
		Status string `json:"status"`
		Div    string `json:"div"`
	} `json:"text"`
	Identifier []struct {
		Value string `json:"value"`
	} `json:"identifier"`
	Status      string   `json:"status"`
	Name        string   `json:"name"`
	Alias       []string `json:"alias"`
	Description string   `json:"description"`
	Mode        string   `json:"mode"`
	Contact     []struct {
		Telecom []struct {
			System string `json:"system"`
			Value  string `json:"value"`
			Use    string `json:"use,omitempty"`
		} `json:"telecom"`
	} `json:"contact"`
	Address struct {
		Use        string   `json:"use"`
		Line       []string `json:"line"`
		City       string   `json:"city"`
		PostalCode string   `json:"postalCode"`
		Country    string   `json:"country"`
	} `json:"address"`
	Form struct {
		Coding []struct {
			System  string `json:"system"`
			Code    string `json:"code"`
			Display string `json:"display"`
		} `json:"coding"`
	} `json:"form"`
	Position struct {
		Longitude float64 `json:"longitude"`
		Latitude  float64 `json:"latitude"`
		Altitude  int     `json:"altitude"`
	} `json:"position"`
	ManagingOrganization struct {
		Reference string `json:"reference"`
	} `json:"managingOrganization"`
	Characteristic []struct {
		Coding []struct {
			System  string `json:"system"`
			Code    string `json:"code"`
			Display string `json:"display"`
		} `json:"coding"`
	} `json:"characteristic"`
	Endpoint []struct {
		Reference string `json:"reference"`
	} `json:"endpoint"`
}
