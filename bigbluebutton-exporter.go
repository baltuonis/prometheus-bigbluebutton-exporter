package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"text/tabwriter"

	"github.com/baltuonis/prometheus-bigbluebutton-exporter/bbb"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/version"

	"gopkg.in/alecthomas/kingpin.v2"
)

const (
	Namespace         = "bbb"
	LabelMeeting      = "meeting"
	LabelParticipants = "participants"
)

var (
	showVersion = kingpin.Flag("version", "Print version information").Bool()
	debug       = kingpin.Flag("debug", "Enable debug features").Bool()
	listenAddr  = kingpin.Flag("web.listen-address", "The address to listen on for HTTP requests.").Default(":9688").String()
	bbbAPI      = kingpin.Flag("bbb.api", "An url that points to BigBlueButton API e.g. https://yoursite.com/bigbluebutton/api/").String()
	bbbSecret   = kingpin.Flag("bbb.secret", "BigBlueButton secret").String()
)

var (
	// BuildTime represents the time of the build
	BuildTime = "N/A"
	// Version represents the Build SHA-1 of the binary
	Version = "N/A"

	// labels are the static labels that come with every metric
	labels = []string{LabelMeeting}

	eventOpts = prometheus.GaugeOpts{
		Name:      "meetings",
		Namespace: Namespace,
		Help:      "Gauge for BigBlueButton meetings",
	}
)

type bbbExporter struct {
	client bbb.BBBClient
}

func (e *bbbExporter) Collect(ch chan<- prometheus.Metric) {
	gv := prometheus.NewGaugeVec(eventOpts, labels)
	e.scrape(gv)
	gv.Collect(ch)
}

func (e *bbbExporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- prometheus.NewDesc(
		prometheus.BuildFQName(eventOpts.Namespace, eventOpts.Subsystem, eventOpts.Name),
		eventOpts.Help,
		labels,
		nil,
	)
}

func (e *bbbExporter) scrape(gv *prometheus.GaugeVec) {
	var meetingsInfo = e.client.GetMeetings()
	if meetingsInfo == nil {
		log.Println("scarpe: Failed to receive meeting data")
	} else {
		var meetings = meetingsInfo.Meetings.Meetings

		for _, e := range meetings {
			gv.WithLabelValues(
				e.MeetingName,
			).Set(float64(e.ParticipantCount))
		}
	}
}

func init() {
	prometheus.MustRegister(version.NewCollector("bbb_exporter"))
}

func main() {

	registerSignals()
	kingpin.Parse()

	if *showVersion {
		tw := tabwriter.NewWriter(os.Stdout, 2, 1, 2, ' ', 0)
		fmt.Fprintf(tw, "Build Time:   %s\n", BuildTime)
		fmt.Fprintf(tw, "Build SHA-1:  %s\n", Version)
		fmt.Fprintf(tw, "Go Version:   %s\n", runtime.Version())
		tw.Flush()
		os.Exit(0)
	}

	if len(*bbbAPI) < 1 || len(*bbbSecret) < 1 {
		if *debug {
			fmt.Printf("Error: bbbAPI & bbbSecret are required")
			log.Printf("BaseURL: '%s'\n", *bbbAPI)
			log.Printf("Secret: '%s'\n", *bbbSecret)
		}
		fmt.Printf("Error: bbbAPI & bbbSecret are required")
		os.Exit(0)
	}

	log.Printf("Starting `bigbluebutton-exporter`: Build Time: '%s' Build SHA-1: '%s'\n", BuildTime, Version)
	log.Printf("BBB Endpoint: '%s'\n", *bbbAPI)

	mux := http.NewServeMux()

	log.Printf("Trying to connect...")
	bbbClient := &bbb.BBBClient{BaseURL: *bbbAPI, Secret: *bbbSecret, Debug: *debug}
	var meetingsInfo = bbbClient.GetMeetings()

	if meetingsInfo == nil {
		log.Printf("Couldn't connect to the BBB server. Check your BaseURL or Secret.")
		os.Exit(0)
	}
	log.Printf("Successfully connected to the BBB server.")

	exporter := &bbbExporter{client: *bbbClient}
	prometheus.MustRegister(exporter)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
             <head><title>Big Blue Button exporter</title></head>
             <body>
             <h1>Prometheus Big Blue Button exporter</h1>
             <p><a href='/metrics'>Metrics</a></p>
             </body>
             </html>`))
	})

	mux.Handle("/metrics", promhttp.Handler())

	if *debug {
		mux.HandleFunc("/debug", func(w http.ResponseWriter, r *http.Request) {
			var info = bbbClient.GetMeetings()
			b, err := json.Marshal(info)
			if err != nil {
				fmt.Printf("Error: %s", err)
				return
			}
			w.Write([]byte(b))
		})
	}

	log.Println("Listening on", *listenAddr)
	log.Fatal(http.ListenAndServe(*listenAddr, mux))
}

func registerSignals() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		log.Print("Received SIGTERM, exiting...")
		os.Exit(1)
	}()
}
