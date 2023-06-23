package entity

type Task struct {
	ID        uint   `json:"id"`
	Details   string `json:"details"`
	Completed bool   `json:"completed"`
}
