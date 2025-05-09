package recommendation

import (
	"encoding/json"
)

func environmentHasJupyter(environment Environment) bool {
	for _, pkg := range environment.Packages {
		if pkg.Name == "jupyter" {
			return true
		}
	}
	return false
}

func Match(serializedImageConfig []byte) (bool, error) {
	var configuration ImageConfiguration

	if err := json.Unmarshal(serializedImageConfig, &configuration); err != nil {
		return false, err
	}

	// Check if ANY environment has a package named "jupyter"
	for _, env := range configuration.Environments {
		if environmentHasJupyter(env) {
			return true, nil
		}
	}

	// we got this far, so none of the environments have jupyter
	return false, nil
}

func Recommend(serializedImageConfig []byte) ([]byte, error) {
	var configuration ImageConfiguration

	if err := json.Unmarshal(serializedImageConfig, &configuration); err != nil {
		return nil, err
	}

	// Add an unconstrained version of python to each environment
	for envName, env := range configuration.Environments {
		if !environmentHasJupyter(env) {
			continue
		}

		command := []string{"jupyter", "lab", "--ip=0.0.0.0", "--no-browser"}
		if configuration.User != nil && configuration.User.UID == 0 {
			command = append(command, "--allow-root")
		}

		configuration.Startup.StartupConfigurations = append(configuration.Startup.StartupConfigurations, StartupConfiguration{
			Environment: envName,
			Command:     command,
		})

		// save the environment back to the config
		configuration.Environments[envName] = env

	}

	// Convert back to JSON
	updatedImageConfiguration, err := json.Marshal(configuration)
	if err != nil {
		return nil, err
	}

	return updatedImageConfiguration, nil
}
