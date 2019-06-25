package version

// Version is overridden at compile time by `-ldflags "-X github.com/anderskvist/GoHelpers/version.Version=`date -u '+%Y%m%d-%H%M%S'`-`git rev-parse --short HEAD`"`
var Version = "developer-build"
