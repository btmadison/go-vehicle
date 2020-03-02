package crud

// Vehicle type identified by unique vin number
type Vehicle struct {
	Vin        string `json:"vin"`
	Make       string `json:"make"`
	Model      string `json:"model"`
	Year       int    `json:"year"`
	Dealership string `json:"dealership"`
}
