package stats

import (
	"net"
	"os"

	"github.com/rcrowley/go-metrics"
	"github.com/rcrowley/go-metrics/stathat"

	"github.com/70data/dnsmasq-goalng/server"
)

var (
	graphiteServer = os.Getenv("GRAPHITE_SERVER")
	graphitePrefix = os.Getenv("GRAPHITE_PREFIX")
	stathatUser    = os.Getenv("STATHAT_USER")
)

func init() {
	if graphitePrefix == "" {
		graphitePrefix = "go-dnsmasq"
	}
	server.StatsForwardCount = metrics.NewCounter()
	metrics.Register("go-dnsmaq-forward-requests", server.StatsForwardCount)
	server.StatsStubForwardCount = metrics.NewCounter()
	metrics.Register("go-dnsmaq-stub-forward-requests", server.StatsStubForwardCount)
	server.StatsDnssecOkCount = metrics.NewCounter()
	metrics.Register("go-dnsmaq-dnssecok-requests", server.StatsDnssecOkCount)
	server.StatsDnssecCacheMiss = metrics.NewCounter()
	metrics.Register("go-dnsmaq-dnssec-cache-miss", server.StatsDnssecCacheMiss)
	server.StatsLookupCount = metrics.NewCounter()
	metrics.Register("go-dnsmaq-internal-lookups", server.StatsLookupCount)
	server.StatsRequestCount = metrics.NewCounter()
	metrics.Register("go-dnsmaq-requests", server.StatsRequestCount)
	server.StatsNameErrorCount = metrics.NewCounter()
	metrics.Register("go-dnsmaq-nameerror-responses", server.StatsNameErrorCount)
	server.StatsNoDataCount = metrics.NewCounter()
	metrics.Register("go-dnsmaq-nodata-responses", server.StatsNoDataCount)
	server.StatsCacheMiss = metrics.NewCounter()
	metrics.Register("go-dnsmaq-nodata-responses", server.StatsCacheMiss)
	server.StatsCacheHit = metrics.NewCounter()
	metrics.Register("go-dnsmaq-nodata-responses", server.StatsCacheHit)
}

func Collect() {
	if graphiteServer != "" {
		addr, err := net.ResolveTCPAddr("tcp", graphiteServer)
		if err == nil {
			go metrics.Graphite(metrics.DefaultRegistry, 10e9, graphitePrefix, addr)
		}
	}
	if stathatUser != "" {
		go stathat.Stathat(metrics.DefaultRegistry, 10e9, stathatUser)
	}
}