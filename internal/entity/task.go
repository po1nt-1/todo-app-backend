package entity

// Task -.
type Task struct {
	ID        uint   `json:"id"`
	Details   string `json:"details"`
	Completed bool   `json:"completed"`
}
