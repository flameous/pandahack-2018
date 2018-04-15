package types

type Task struct {
	Id            int             `json:"id"`
	Title         string          `json:"title"`
	PersonalTasks []*PersonalTask `json:"personal_tasks"`
}

type PersonalTask struct {
	Id        int    `json:"id"`
	User      *User  `json:"user"`
	Completed bool   `json:"completed"`
	Question  string `json:"question"`
	Answer    string `json:"answer"`
}
