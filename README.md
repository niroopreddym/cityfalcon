
# Accounts Falcon

REST API built in GO Mux handling long running operations over event driven architecture using RabbitMQ and go-routines.

Monitoring the HTTP Calls using Prometheus and Graphana, The metrics for preometheus are collected by embedding a middleware layer between HTTPHandler Func and the request.

Caching frequently used get calls using REDIS.

POSTGRES is used as the SQL Database.


## Long-Running Task Architecture

![App Screenshot](https://raw.githubusercontent.com/niroopreddym/images/38c634ca50b4eb0f619db361266bf746f853a80b/longrunningtaskevented.drawio.svg)

### API Reference

#### Create a new bank

```http
POST URI:    localhost:9295/api/bank
request-payload:
    {
    "bankname":"testbank",
    "ifsccode":"TEST00008017",
    "branchname":"stonehousepet"
    }
response-payload:
    {
        "BankUUID": "dcb50ac8-3420-49ca-9980-7986a3b6d5b8"
    }
```

#### Create a Bank Account

```http
POST URI:    localhost:9295/api/account
request-payload:
    {
        "accountholdername":"Yams",
        "bankid":1,
        "firstname":"Yamuna",
        "lastname":"Karuka",
        "balance": 9256.89
    }
response-payload:
    {
        "AccountUUID": "dcb50ac8-3420-49ca-9980-7986a3b6d5b8"
    }
```

#### Get Bank Balance - ( simulated long-running process )

```http
GET URI:    localhost:9295/api/account/{id}
response-payload:
    {
        "trackingURL": "/account/getaccountdetails/asyncresponse/{id}"
    }
```

| Parameter | Type     | Description                                    |
| :-------- | :------- | :--------------------------------------------- |
| `id`      | `string` | **Required**. UUID of account balance to fetch |

## Booting the project using Docker-Compose

### understanding bridge-network for metric collections to prometheus in Docker-Compose

```Note
The Prometheus and Graphana containers have to be bridge connected so assign a common network to both of them in the compose file. Like below

prometheus:
    image: niroop/prom
    build: 
      context: ./prometheus
      dockerfile: Dockerfile
    ports:
      - "9090:9090"
    command:
      - '--config.file=/etc/prometheus/prometheus.yml'
      - '--storage.tsdb.path=/prometheus'
      - '--web.console.libraries=/usr/share/prometheus/console_libraries'
      - '--web.console.templates=/usr/share/prometheus/consoles'
    volumes:
      - ./prometheus/:/etc/prometheus/
      - prometheus_data:/prometheus
    networks:
      - gateway
  graphana:
    image: grafana/grafana:latest
    container_name: graphana
    ports:
      - "3000:3000"
    networks:
      - gateway
    volumes:
      - grafana-storage:/var/lib/grafana
networks:
  gateway: {}
  
```

### Docker-Compose Boot Commands

navigate to the infrastrcuture folder on root level

```docker
 docker-compose build
 docker-compose up -d
```

## Monitoring solution using Prometheus and Graphana

Used the atomic metrics collection middleware inside the gorilla-mux router map.

![App Screenshot](https://raw.githubusercontent.com/niroopreddym/images/master/graphana.jpg)
