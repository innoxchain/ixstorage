package version

var (
	//Commit is the git commit hash when the software was built
	Commit string

	//Branch is the branch of the current version
	Branch string

	//BuildTime is the time when the built was completed
	BuildTime string
)