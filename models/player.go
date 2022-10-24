package models

type Player struct {
	ClientID int `json:"client_id"`
	Username string `json:"user_name"`
	Position Position `json:"position"`
	Animation Animation `json:"animation"`

}

type Animation struct {
	Idle string `json:"idle"`
	Sleeping string `json:"sleeping"`
	Running string `json:"running"`
	Angry string `json:"angry"`
}

type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
}
