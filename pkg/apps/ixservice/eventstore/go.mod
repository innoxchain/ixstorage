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
	github.com/sirupsen/logrus v1.3.0
	github.com/stretchr/testify v1.3.0 // indirect
	golang.org/x/crypto v0.0.0-20190103213133-ff983b9c42bc // indirect
	golang.org/x/sys v0.0.0-20190108104531-7fbe1cd0fcc2 // indirect
)
