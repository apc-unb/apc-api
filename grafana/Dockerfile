FROM grafana/grafana:latest

COPY dashboards.yml     /etc/grafana/provisioning/dashboards/
COPY datasources.yml    /etc/grafana/provisioning/datasources/
COPY dashboards/*.json  /var/lib/grafana/dashboards/