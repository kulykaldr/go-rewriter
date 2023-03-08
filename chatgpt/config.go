package chatgpt

type Config struct {
	Login     string `json:"login"`
	Password  string `json:"password"`
	IsLogin   bool   `json:"is_login,omitempty"`
	Headless  bool   `json:"headless,omitempty"`
	Debug     bool   `json:"debug,omitempty"`
	Timeout   int64  `json:"timeout,omitempty"`
	UserAgent string `json:"user_agent,omitempty"`
}
