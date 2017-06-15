package vakthund

type Status struct {
	Sha        string
	State      string
	Branches   []Branch
	Repository Repo
}

type Branch struct {
	Name   string
	Commit map[string]string
}

type Repo struct {
	FullName string `json:"full_name"`
}
