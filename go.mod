module github.com/innoxchain/ixstorage

replace github.com/innoxchain/ixstorage/cmd/ixclient v0.0.3 => ./cmd/ixclient

replace github.com/innoxchain/ixstorage/build v0.0.3 => ./build

require (
	github.com/innoxchain/ixstorage/cmd/ixclient v0.0.3 // indirect
	github.com/innoxchain/ixstorage/pkg/apps/ixservice/eventstore v0.0.0-20190105211345-1314a4ed27e9 // indirect
	github.com/lib/pq v1.0.0 // indirect
	github.com/sirupsen/logrus v1.3.0
)
