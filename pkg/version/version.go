package version

var (
	version  = "dev"
	revision string
)

func Version() string {
	return version
}

func Revision() string {
	return revision
}
