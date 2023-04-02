package build

var Commit string

const version = "0.1.0"

func GetVersion() string {
	basicVersion := "v" + version

	if Commit == "" {
		return basicVersion
	}

	return basicVersion + "-" + Commit
}
