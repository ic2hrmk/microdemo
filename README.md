# MicroDemo

Super-simple micro service based on [go-micro](https://github.com/micro/micro) framework which could be used 
for scratches and development purposes.

## Usage

### MicroDemo Web [FQDN: dev.ic2hrmk.api.microdemo]

#### Build

`make build`

#### Tune
Run flags:
~~~bash
--handler value   Path to receive HTTP requests (default: "/api/call") [$MICRODEMO_WEB_HANDLER]
--version         Shows a version of application and exists
~~~

#### Docker Image

TODO:
* [ ] Add a link to image at DockerHub
* [ ] Add description and usage example
  

## Micro Tune

~~~bash
MICRO_CLIENT                     --client value                     Client for go-micro; rpc
MICRO_CLIENT_REQUEST_TIMEOUT     --client_request_timeout value     Sets the client request timeout. e.g 500ms, 5s, 1m. Default: 5s
MICRO_CLIENT_RETRIES             --client_retries value             Sets the client retries. Default: 1 (default: 1)
MICRO_CLIENT_POOL_SIZE           --client_pool_size value           Sets the client connection pool size. Default: 1 (default: 0)
MICRO_CLIENT_POOL_TTL            --client_pool_ttl value            Sets the client connection pool ttl. e.g 500ms, 5s, 1m. Default: 1m
MICRO_REGISTER_TTL               --register_ttl value               Register TTL in seconds (default: 60)
MICRO_REGISTER_INTERVAL          --register_interval value          Register interval in seconds (default: 30)
MICRO_SERVER                     --server value                     Server for go-micro; rpc
MICRO_SERVER_NAME                --server_name value                Name of the server. go.micro.srv.example
MICRO_SERVER_VERSION             --server_version value             Version of the server. 1.1.0
MICRO_SERVER_ID                  --server_id value                  Id of the server. Auto-generated if not specified
MICRO_SERVER_ADDRESS             --server_address value             Bind address for the server. 127.0.0.1:8080
MICRO_SERVER_ADVERTISE           --server_advertise value           Used instead of the server_address when registering with discovery. 127.0.0.1:8080
MICRO_SERVER_METADATA            --server_metadata value            A list of key-value pairs defining metadata. version=1.0.0
MICRO_BROKER                     --broker value                     Broker for pub/sub. http, nats, rabbitmq
MICRO_BROKER_ADDRESS             --broker_address value             Comma-separated list of broker addresses
MICRO_DEBUG_PROFILE              --profile value                    Debug profiler for cpu and memory stats
MICRO_REGISTRY                   --registry value                   Registry for discovery. etcd, mdns
MICRO_REGISTRY_ADDRESS           --registry_address value           Comma-separated list of registry addresses
MICRO_RUNTIME                    --runtime value                    Runtime for building and running services e.g local, kubernetes (default: "local")
MICRO_SELECTOR                   --selector value                   Selector used to pick nodes for querying
MICRO_STORE                      --store value                      Store used for key-value storage
MICRO_STORE_ADDRESS              --store_address value              Comma-separated list of store addresses
MICRO_STORE_NAMESPACE            --store_namespace value            Namespace for store data
MICRO_TRANSPORT                  --transport value                  Transport mechanism used; http (default: "nats")
MICRO_TRANSPORT_ADDRESS          --transport_address value          Comma-separated list of transport addresses
MICRO_TRACER                     --tracer value                     Tracer for distributed tracing, e.g. memory, jaeger (default: "memory")
MICRO_TRACER_ADDRESS             --tracer_address value             Comma-separated list of tracer addresses
~~~
