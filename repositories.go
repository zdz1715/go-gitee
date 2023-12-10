package gitee

type Repo struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Fork        bool      `json:"fork"`
	FullName    string    `json:"full_name"`
	HtmlURL     string    `json:"html_url"`
	HumanName   string    `json:"human_name"`
	Internal    bool      `json:"internal"`
	Name        string    `json:"name"`
	Namespace   Namespace `json:"namespace"`
	Owner       User      `json:"owner"`
	Path        string    `json:"path"`
	Private     bool      `json:"private"`
	Public      bool      `json:"public"`
	SshURL      string    `json:"ssh_url"`
	URL         string    `json:"url"`
}

type Namespace struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Path    string `json:"path"`
	Type    string `json:"type"`
	HtmlURL string `json:"html_url"`
}
