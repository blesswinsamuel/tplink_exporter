package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/blesswinsamuel/tplink_exporter/ipdb"
	"github.com/blesswinsamuel/tplink_exporter/tplink"
)

func GetEnvStr(name, value string) string {
	if os.Getenv(name) != "" {
		return os.Getenv(name)
	}
	return value
}

func main() {
	Address := flag.String(
		"a",
		GetEnvStr("TPLINK_ROUTER_ADDR", "192.168.0.1"),
		"Router's address",
	)
	User := flag.String(
		"u",
		GetEnvStr("TPLINK_ROUTER_USER", "admin"),
		"Router's username",
	)
	Pass := flag.String(
		"w",
		GetEnvStr("TPLINK_ROUTER_PASSWD", "admin"),
		"Router's password",
	)
	Port := flag.Int(
		"p",
		9300,
		"Prometheus port",
	)
	Verbose := flag.Bool(
		"v",
		false,
		"Verbose output",
	)
	Filename := flag.String(
		"f",
		GetEnvStr("TPLINK_ROUTER_IPS", "/etc/hosts"),
		"IP Database - hosts file",
	)
	flag.Parse()

	ips, err := ipdb.Load(*Filename)
	if err != nil {
		log.Println("Unable to load IP database:", err)
	} else {
		log.Printf("%d custom IPs loaded\n", len(ips))
	}

	router := tplink.NewRouter(*Address, *User, *Pass)
	router.Verbose = *Verbose

	c := newRouterCollector(router, ips)
	prometheus.MustRegister(c)

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(*Port), nil))
}
