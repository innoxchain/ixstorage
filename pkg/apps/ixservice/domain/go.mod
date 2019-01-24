module github.com/innoxchain/ixstorage/pkg/apps/ixservice/domain

replace github.com/innoxchain/ixstorage/pkg/apps/ixservice/domain/event v0.0.0 => ./event

replace github.com/innoxchain/ixstorage/pkg/apps/ixservice/domain/enum v0.0.0 => ./enum

replace github.com/innoxchain/ixstorage/pkg/apps/ixservice/eventstore v0.0.0 => ../eventstore

replace github.com/innoxchain/ixstorage/config v0.0.0 => ../../../../config

require (
	github.com/innoxchain/ixstorage/pkg/apps/ixservice/domain/enum v0.0.0
	github.com/innoxchain/ixstorage/pkg/apps/ixservice/domain/event v0.0.0
	github.com/innoxchain/ixstorage/pkg/apps/ixservice/eventstore v0.0.0
	github.com/satori/go.uuid v1.2.0
	github.com/sirupsen/logrus v1.3.0
	golang.org/x/sys v0.0.0-20190109145017-48ac38b7c8cb // indirect
)
