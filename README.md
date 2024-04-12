# Simple Monitoring Stack

This project sets up a simple monitoring stack using Docker Compose, which includes a Go application exposing Prometheus metrics, Prometheus for metrics scraping, and Grafana for metrics visualization.

## Components

1. **Go Application**: A simple Go server that exposes custom metrics to Prometheus.
2. **Prometheus**: Scrapes and stores metrics provided by the Go application.
3. **Grafana**: Visualizes the metrics collected by Prometheus using a pre-configured dashboard.

## Prerequisites

Make sure you have Docker and Docker Compose installed on your machine.

## Setup Instructions

1. **Clone the repository**

2. **Start the services**
   ```
   docker-compose up --build -d
   ```

   This command builds the Docker images and starts the services in the background.

3. **Access the services**:
   - Go Application: [http://localhost:8080/metrics](http://localhost:8080/metrics) to view the metrics exposed by the Go application.
   - Prometheus: [http://localhost:9090](http://localhost:9090) to access the Prometheus web UI.
   - Grafana: [http://localhost:3000](http://localhost:3000) to access the Grafana dashboard. Default login is `admin` for both username and password.

## Grafana Dashboard

To view the dashboard:
- Navigate to Grafana at `http://localhost:3000`.
- Log in with the default credentials (admin/admin).
- Go to "Dashboards".
- You should see the pre-configured dashboard "My Dashboard" listed there. Click on it to view the metrics.

## Stopping the Services

To stop and remove all running services, use the following Docker Compose command:
```
docker-compose down
```

This command stops all the running containers and removes them along with their networks.

## Conclusion

This setup provides a basic example of how to monitor a Go application using Prometheus and Grafana. It is easily extendable and can be adapted to monitor multiple services or applications.

