module github.com/innoxchain/ixstorage

replace github.com/innoxchain/ixstorage/cmd/ixclient v0.0.3 => ./cmd/ixclient

replace github.com/innoxchain/ixstorage/build v0.0.3 => ./build

require (
	github.com/innoxchain/ixstorage/cmd/ixclient v0.0.3 // indirect
	github.com/innoxchain/ixstorage/pkg/apps/ixservice/domain/enum v0.0.0-20190105211345-1314a4ed27e9
	github.com/innoxchain/ixstorage/pkg/apps/ixservice/eventstore v0.0.0-20190105211345-1314a4ed27e9 // indirect
	github.com/lib/pq v1.0.0 // indirect
	github.com/sirupsen/logrus v1.3.0
	github.com/stretchr/testify v1.3.0 // indirect
	golang.org/x/crypto v0.0.0-20190131182504-b8fe1690c613 // indirect
	golang.org/x/sys v0.0.0-20190201152629-afcc84fd7533 // indirect
)
