module github.com/innoxchain/ixstorage/pkg/apps/ixservice/examples/eventsourcing

replace github.com/innoxchain/ixstorage/pkg/apps/ixservice/event v0.0.0 => ../../event

replace github.com/innoxchain/ixstorage/config v0.0.0 => ../../../../../config

require (
	github.com/innoxchain/ixstorage/pkg/apps/ixservice/event v0.0.0
	github.com/konsorten/go-windows-terminal-sequences v1.0.2 // indirect
	github.com/satori/go.uuid v1.2.0
	github.com/sirupsen/logrus v1.3.0
	golang.org/x/crypto v0.0.0-20190228161510-8dd112bcdc25 // indirect
	golang.org/x/sys v0.0.0-20190228124157-a34e9553db1e // indirect
)
