package collector

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	appUp = prometheus.NewDesc(
		prometheus.BuildFQName("", "app", "up"),
		"Was the last query of the API successful.",
		nil, nil)
)

type Exporter struct {
	Username string
	Password string
}

func New(opts ...Option) prometheus.Collector {
	e := &Exporter{}
	for _, opt := range opts {
		opt(e)
	}
	return e
}

type Option func(*Exporter)

func WithCredentials(username, password string) Option {
	return func(e *Exporter) {
		e.Username = username
		e.Password = password
	}
}
func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- appUp
}

func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	ch <- prometheus.MustNewConstMetric(appUp, prometheus.GaugeValue, 1)
}
