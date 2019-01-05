module github.com/innoxchain/ixstorage/pkg/apps/ixservice/domain

replace github.com/innoxchain/ixstorage/pkg/apps/ixservice/domain/event v0.0.0 => ./event

replace github.com/innoxchain/ixstorage/pkg/apps/ixservice/domain/enum v0.0.0 => ./enum

require (
	github.com/innoxchain/ixstorage/pkg/apps/ixservice/domain/enum v0.0.0
	github.com/innoxchain/ixstorage/pkg/apps/ixservice/domain/event v0.0.0
	github.com/kr/pretty v0.1.0 // indirect
	github.com/satori/go.uuid v1.2.0
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
)
