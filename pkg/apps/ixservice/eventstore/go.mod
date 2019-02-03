module github.com/innoxchain/ixstorage/pkg/apps/ixservice/eventstore

replace github.com/innoxchain/ixstorage/config v0.0.0 => ../../../../config

replace github.com/innoxchain/ixstorage/pkg/apps/ixservice/domain/event v0.0.0 => ../domain/event

replace github.com/innoxchain/ixstorage/pkg/apps/ixservice/domain/enum v0.0.0 => ../domain/enum

require (
	github.com/innoxchain/ixstorage/config v0.0.0
	github.com/innoxchain/ixstorage/pkg/apps/ixservice/domain/enum v0.0.0
	github.com/innoxchain/ixstorage/pkg/apps/ixservice/domain/event v0.0.0
	github.com/lib/pq v1.0.0
	github.com/pkg/errors v0.8.1
	github.com/satori/go.uuid v1.2.0
	github.com/sirupsen/logrus v1.3.0
	github.com/stretchr/testify v1.3.0
)
