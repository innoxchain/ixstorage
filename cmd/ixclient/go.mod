module github.com/innoxchain/ixstorage/cmd/ixclient

replace github.com/innoxchain/ixstorage/build v0.0.3 => ../../build

require (
	github.com/innoxchain/ixstorage/build v0.0.3
	github.com/innoxchain/ixstorage/pkg/apps/ixclient v0.0.1
	github.com/spf13/cobra v0.0.3
	github.com/spf13/pflag v1.0.3 // indirect
)
