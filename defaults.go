package micro

import (
	"github.com/chnkenc/go-micro/client"
	"github.com/chnkenc/go-micro/debug/trace"
	"github.com/chnkenc/go-micro/server"
	"github.com/chnkenc/go-micro/store"

	// set defaults
	gcli "github.com/chnkenc/go-micro/client/grpc"
	memTrace "github.com/chnkenc/go-micro/debug/trace/memory"
	gsrv "github.com/chnkenc/go-micro/server/grpc"
	memoryStore "github.com/chnkenc/go-micro/store/memory"
)

func init() {
	// default client
	client.DefaultClient = gcli.NewClient()
	// default server
	server.DefaultServer = gsrv.NewServer()
	// default store
	store.DefaultStore = memoryStore.NewStore()
	// set default trace
	trace.DefaultTracer = memTrace.NewTracer()
}
