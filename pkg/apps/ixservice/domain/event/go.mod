module github.com/innoxchain/ixstorage/pkg/apps/ixservice/domain/event

replace github.com/innoxchain/ixstorage/pkg/apps/ixservice/domain/enum v0.0.0 => ../enum

replace github.com/innoxchain/ixstorage/config v0.0.0 => ../../../../../config

require (
	github.com/innoxchain/ixstorage/config v0.0.0
	github.com/innoxchain/ixstorage/pkg/apps/ixservice/domain/enum v0.0.0
	github.com/kr/pretty v0.1.0 // indirect
	github.com/kr/pty v1.1.3 // indirect
	github.com/lib/pq v1.0.0
	github.com/pkg/errors v0.8.1
	github.com/satori/go.uuid v1.2.0
	github.com/sirupsen/logrus v1.3.0
	github.com/stretchr/testify v1.3.0
	golang.org/x/crypto v0.0.0-20190222235706-ffb98f73852f // indirect
	golang.org/x/sys v0.0.0-20190222171317-cd391775e71e
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
)
