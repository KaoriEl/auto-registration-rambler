package structures

type AccInfo struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Gender   string `json:"gender"`
	City     string `json:"city"`
	ClientId int64  `json:"client_id"`
	Service  string `json:"service"`
	Status   int64  `json:"false"`
	Password string
	Email    string
}

type ReadyAccInfo struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Gender   string `json:"gender"`
	City     string `json:"city"`
	ClientId int64  `json:"client_id"`
	Service  string `json:"service"`
	Status   int64  `json:"false"`
	Password string `json:"password"`
	Email    string
}
