package prom

import (
	"github.com/prometheus/client_golang/prometheus"
)

// Prom present prom metrics
//
// You can add metrics like below
// SomeGaugeMetric = New().WithState("gauge_metric_name", "gauge metric name", []string{
//     labels...
// })
// and set state via `SomeGaugeMetric.State("gauge_metric_label_name", gauge_value, extra_labels...)`
type Prom struct {
	timer   *prometheus.HistogramVec
	counter *prometheus.CounterVec
	state   *prometheus.GaugeVec
}

// New creates a Prom instance.
func New() *Prom {
	return &Prom{}
}

// WithTimer sets timer.
func (p *Prom) WithTimer(name, desc string, labels []string) *Prom {
	if p == nil || p.timer != nil {
		return p
	}
	p.timer = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    name,
			Help:    desc,
			Buckets: prometheus.LinearBuckets(0, 10, 1),
		}, labels)
	prometheus.MustRegister(p.timer)
	return p
}

// WithCounter sets counter.
func (p *Prom) WithCounter(name, desc string, labels []string) *Prom {
	if p == nil || p.counter != nil {
		return p
	}
	p.counter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: name,
			Help: desc,
		}, labels)
	prometheus.MustRegister(p.counter)
	return p
}

// WithState sets state.
func (p *Prom) WithState(name, desc string, labels []string) *Prom {
	if p == nil || p.state != nil {
		return p
	}
	p.state = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: name,
			Help: desc,
		}, labels)
	prometheus.MustRegister(p.state)
	return p
}

// ResetState reset state.
func (p *Prom) ResetState() *Prom {
	if p == nil || p.state == nil {
		return p
	}
	p.state.Reset()
	return p
}

// ResetCounter reset counter.
func (p *Prom) ResetCounter() *Prom {
	if p == nil || p.counter == nil {
		return p
	}
	p.counter.Reset()
	return p
}

// Timing log timing information (in milliseconds) without sampling
func (p *Prom) Timing(name string, time int64, extra ...string) {
	if p.timer != nil {
		label := append([]string{name}, extra...)
		p.timer.WithLabelValues(label...).Observe(float64(time))
	}
}

// Incr increments one stat counter without sampling
func (p *Prom) Incr(name string, extra ...string) {
	if p.counter != nil {
		label := append([]string{name}, extra...)
		p.counter.WithLabelValues(label...).Inc()
	}
}

// State set state
func (p *Prom) State(name string, v int64, extra ...string) {
	if p.state != nil {
		label := append([]string{name}, extra...)
		p.state.WithLabelValues(label...).Set(float64(v))
	}
}

// Add add count v must > 0
func (p *Prom) Add(name string, v int64, extra ...string) {
	if p.counter != nil {
		label := append([]string{name}, extra...)
		p.counter.WithLabelValues(label...).Add(float64(v))
	}
}
