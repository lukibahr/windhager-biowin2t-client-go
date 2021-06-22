package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

type windhagerCollector struct {
	biowinTimeToCoreMaintenance *prometheus.Desc
	biowinTimeToMaintenance     *prometheus.Desc
}

//You must create a constructor for you collector that
//initializes every descriptor and returns a pointer to the collector
func newWindhagerCollector() *windhagerCollector {
	return &windhagerCollector{
		biowinTimeToMaintenance: prometheus.NewDesc("biowin_maintenance_time",
			"Time in hours to next maintenance",
			nil, nil,
		),
		biowinTimeToCoreMaintenance: prometheus.NewDesc("biowin_core_maintenance_time",
			"Time in hours to next core maintenance",
			nil, nil,
		),
	}
}

//Describe implements the describe function of a collector
func (collector *windhagerCollector) Describe(ch chan<- *prometheus.Desc) {

	ch <- collector.biowinTimeToMaintenance
	ch <- collector.biowinTimeToCoreMaintenance
}

//Collect implements required collect function for all promehteus collectors
func (collector *windhagerCollector) Collect(ch chan<- prometheus.Metric) {

	//Implement logic here to determine proper metric value to return to prometheus
	//for each descriptor or call other functions that do so.
	var metricValue float64
	if 1 == 1 {
		metricValue = 1
	}

	//Write latest value for each metric in the prometheus metric channel.
	//Note that you can pass CounterValue, GaugeValue, or UntypedValue types here.
	ch <- prometheus.MustNewConstMetric(collector.biowinTimeToMaintenance, prometheus.CounterValue, metricValue)
	ch <- prometheus.MustNewConstMetric(collector.biowinTimeToCoreMaintenance, prometheus.CounterValue, metricValue)

}
