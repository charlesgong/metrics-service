# metrics-service

# Assignment description
collector includes 
1. Metrics Collection: collector.go
2. Data Storage: pusher.go (still wondering about this part, best practise might expose metrics and wait Prometheus to pull, if we want to save to Prometheus we need to push)
3. API Endpoint: main.go 
~~4. alert~~ alert is in [here](github.com/charlesgong/alert-service)
5. Visualization: import docs/grafana_sample.json to a Grafana, docs/grafana-screenshot.png is the sample grafana dashboard's screenshot at Sep 15 2024 10:34 a.m. UTC+13
6. Documentation: contains most of the instruction to setup the project



# Go Metrics

Collecting server metrics is essential for maintaining the health, performance, and security of IT infrastructure, as well as for optimizing resource usage and ensuring compliance with industry standards and regulations.
### Promethues

Metric storage

### Grafana

Visualization, please import sample-grafana.json after startup

## Getting Start
In our Go project, we just need to import the dependent library.

```go mod download```

Attach suitable metrics to relevant functions to evaluate the performance of the scope. Check example code [here](./main.go).

## Step One

Start the server to collect the metrics.

```go run main.go```

and check the metrics by using follow URL,

```http:\\localhost:9101\metrics```

you can find the metrics begin as **go_metrics_***

## Step Two

Server metrics need to be collected and visualized to analyze the data. Prometheus and Grafana need to be started and configured along with Docker.

### Docker Setup

Install Docker with Docker Compose. Check out [here](https://docs.docker.com/engine/install/)

Prometheus and Grafana Docker configuration can be found in docker-compose.yml. First of all, we need to create the network mode,

```bash 
docker network create metrics
```

Then initiate the Docker Compose by using the following command,

```bash
./metrics-run.sh
```


Now you can check the promethues dashboard by using follow link,

```http://localhost:9090/graph```

Server target is already set via the [promethues.yml](./collector/prometheus/prometheus.yml)

#### Grafana

Let's check the grafana dashboard through,

```http://localhost:3000/```

default username: admin, default password: admin

Configure the grafana database configuration by adding database,








