package version

// Version is overridden at compile time by `-ldflags "-X version.Version $VERSION"`
var Version = "developer-build"
