package recommendation

type Package struct {
	Name string `json:"name"`
}

type Environment struct {
	Channels []string  `json:"channels"`
	Packages []Package `json:"packages"`
}

type StartupConfiguration struct {
	Environment string   `json:"environment"`
	Command     []string `json:"command"`
}

type UserConfiguration struct {
	Name  string `json:"name"`
	Group string `json:"group"`
	UID   int64  `json:"uid"`
	GID   int64  `json:"gid"`
}

type Startup struct {
	DefaultEnvironment    string                 `json:"defaultEnvironment"`
	StartupConfigurations []StartupConfiguration `json:"startupConfigurations"`
}

type ImageConfiguration struct {
	Environments map[string]Environment `json:"environments"`
	Startup      Startup                `json:"startup"`
	User         *UserConfiguration     `json:"user"`
}
