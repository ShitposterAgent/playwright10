package types

type Account struct {
	Platform    string `json:"platform"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	CookiesPath string `json:"cookiesPath"`
}

type AIConfig struct {
	Provider string `json:"provider"`
	Model    string `json:"model"`
	Endpoint string `json:"endpoint"`
}

type Task struct {
	Platform string `json:"platform"`
	Action   string `json:"action"`
	Content  string `json:"content,omitempty"`
	Target   string `json:"target,omitempty"`
	Schedule string `json:"schedule,omitempty"`
}

type Settings struct {
	Headless  bool   `json:"headless"`
	Proxy     string `json:"proxy"`
	UserAgent string `json:"userAgent"`
}

type Config struct {
	Accounts []Account `json:"accounts"`
	AI       AIConfig  `json:"ai"`
	Tasks    []Task    `json:"tasks"`
	Settings Settings  `json:"settings"`
}
