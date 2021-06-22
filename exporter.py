"""Application exporter"""

import os
import time
from typing import Counter
from prometheus_client import start_http_server, Gauge, Enum
import requests
from requests.auth import HTTPDigestAuth


metrics = {
  "biowin_core_maintenance_time": "/1/60/0/98/9",
  "biowin_maintenance_time": "/1/60/0/98/8",
  "biowin_run_time": "/1/60/0/98/4",
  "biowin_burner_start_count": "/1/60/0/98/3",
  #"Temperatur Abgas": "/1/60/0/98/1",
  #"Aktuelle Kesselleistung": "/1/60/0/98/0": ,
  #"Kesseltemperatur Istwert": "/1/60/0/100/1",
  #"Brennkammertemperatur": "/1/60/0/100/2",
  #"Aktuelle Betriebsphase": "/1/60/0/100/3",
  "biowin_spiral_conveyor_amount": "/1/60/0/100/9"
}


class AppMetrics:
    """
    Representation of Prometheus metrics and loop to fetch and transform
    application metrics into Prometheus metrics.
    """

    def __init__(self, mes_endpoint, polling_interval_seconds=5):
        self.mes_endpoint = mes_endpoint
        self.polling_interval_seconds = polling_interval_seconds

        # Prometheus metrics to collect
        self.biowin_core_maintenance_time = Gauge("biowin_core_maintenance_time", "Time in hours to next maintenance")
        self.biowin_maintenance_time = Gauge("biowin_maintenance_time", "Time in hours to next maintenance")
        self.biowin_run_time = Gauge("biowin_run_time", "Time in hours the pellet heater has run")
        self.biowin_burner_start_count = Counter("biowin_burner_start_count")
        self.biowin_spiral_conveyor_amount = Gauge("biowin_spiral_conveyor_amount", "The amount of pellets in the spiral conveyor")       
        #self.health = Enum("app_health", "Health", states=["healthy", "unhealthy"])

    def run_metrics_loop(self):
        """Metrics fetching loop"""

        while True:
            self.fetch()
            time.sleep(self.polling_interval_seconds)

    def fetch(self):
        """
        Get metrics from application and refresh Prometheus metrics with
        new values.
        """

        for key, value in metrics.items():
            print(key)
            resp = requests.get(url=f"{self.mes_endpoint}{value}", auth=HTTPDigestAuth('Service', 'Pmg|@03T1M{+'))
            status_data = resp.json()


        # Fetch raw status data from the application
        resp = requests.get(url=f"{self.mes_endpoint}/1/60/0/98/9", auth=HTTPDigestAuth('Service', 'Pmg|@03T1M{+'))
        status_data = resp.json()

        # Update Prometheus metrics with application metrics
        self.biowin_maintenance_time.set(status_data["value"])
        #self.b_core_maintenance_time.set(status_data["value"])
        #self.health.state(status_data["health"])

def main():
    """Main entry point"""

    polling_interval_seconds = int(os.getenv("POLLING_INTERVAL_SECONDS", "5"))
    mes_endpoint = str(os.getenv("BIOWIN_MES_ENDPOINT", "http://192.168.2.121/api/1.0/lookup"))
    exporter_port = int(os.getenv("EXPORTER_PORT", "9877"))

    app_metrics = AppMetrics(
        mes_endpoint=mes_endpoint,
        polling_interval_seconds=polling_interval_seconds
    )
    start_http_server(exporter_port)
    app_metrics.run_metrics_loop()

if __name__ == "__main__":
    main()