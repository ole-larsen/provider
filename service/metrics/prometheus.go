package metrics

import (
	"fmt"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	namespace = "provider-service"

	counter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Name:      "endpoint_request_count",
		Help:      "provider service request count.",
	}, []string{"app", "name", "method", "state"})

	histogram = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: namespace,
		Name:      "endpoint_duration_seconds",
		Help:      "time taken to execute endpoint.",
	}, []string{"app", "name", "method", "status"})

	TotalRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "number of get requests.",
		},
		[]string{"path"},
	)
)

type metricResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func newMetricResponseWriter(w http.ResponseWriter) *metricResponseWriter {
	return &metricResponseWriter{w, http.StatusOK}
}

func (w *metricResponseWriter) WriteHeader(code int) {
	w.statusCode = code
	w.ResponseWriter.WriteHeader(code)
}

// SetupHandler enable CORS, handler metrics
func SetupHandler(handler http.Handler) http.Handler {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		lrw := newMetricResponseWriter(w)
		handler.ServeHTTP(lrw, r)
		statusCode := lrw.statusCode
		duration := time.Since(start)
		histogram.WithLabelValues("provider-service", r.URL.String(), r.Method, fmt.Sprintf("%d", statusCode)).Observe(duration.Seconds())
		counter.WithLabelValues("provider-service", r.URL.String(), r.Method, fmt.Sprintf("%d", statusCode)).Inc()
	})
	prometheus.Register(histogram)
	prometheus.Register(counter)
	prometheus.Register(TotalRequests)
	return h
}
