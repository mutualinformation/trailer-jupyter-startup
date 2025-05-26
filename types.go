package recommendation

type VersionSpecifier string

const (
	UnconstrainedVersion VersionSpecifier = "unconstrained"
)

type Version struct {
	Specifier  VersionSpecifier `json:"specifier,omitempty"`
	Constraint string           `json:"constraint"`
}

type Package struct {
	Name        string  `json:"name"`
	Version     Version `json:"version,omitempty"`
	Description string  `json:"description,omitempty"`
	Channel     string  `json:"channel"`
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
	User         *UserConfiguration     `json:"user,omitempty,omitzero"`
}
