package model

type Block struct {
	Index     int
	Timestamp string
	BPM       int
	Hash      string
	PrevHash  string
}

type AddRequest struct {
	BPM int `json:"bpm"`
}
