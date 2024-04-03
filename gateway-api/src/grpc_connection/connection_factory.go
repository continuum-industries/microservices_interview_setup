package grpc_connection

import (
	"flag"
	"log"
	"time"

	"github.com/c2h5oh/datasize"
	"google.golang.org/grpc"
	"google.golang.org/grpc/backoff"
	"google.golang.org/grpc/keepalive"
)

type Servicer int

const (
	CalculatorAPI Servicer = iota
)

type ProtoConnection struct {
	servicer   Servicer
	connection *grpc.ClientConn
}

var (
	calculatorAddress = flag.String("calculator-api-address", "calculator-api:50051", "Sets the address for the CalculatorAPI")
	liveConnections   = make(map[Servicer]*grpc.ClientConn)
	addresses         = make(map[Servicer]string)
)

func init() {
	addresses[CalculatorAPI] = *calculatorAddress
}

func GetConnection(s Servicer) *grpc.ClientConn {
	if _, ok := liveConnections[s]; !ok {
		liveConnections[s] = setupgRPCConnection(addresses[s])
	}
	return liveConnections[s]
}

func CloseConnections() {
	for _, element := range liveConnections {
		element.Close()
	}
}

func setupgRPCConnection(address string) *grpc.ClientConn {
	backoffConfig := backoff.Config{BaseDelay: time.Second, Multiplier: 2, MaxDelay: 10 * time.Second}
	connOptions := grpc.WithConnectParams(grpc.ConnectParams{Backoff: backoffConfig, MinConnectTimeout: 20 * time.Second})
	proxyOptions := grpc.WithNoProxy()
	msgOptions := grpc.WithMaxMsgSize(int(5 * datasize.MB))
	serviceConfigJSONString := `{
        "methodConfig": [
            {
                "name": [
                    {"service": "geodata.Handler"},
                    {"service": "tiling.AbstractTiling"}
                ],
                "retryPolicy": {
                    "maxAttempts": 25,
                    "initialBackoff": "0.3s",
                    "maxBackoff": "10s",
                    "backoffMultiplier": 3,
                    "retryableStatusCodes": ["UNAVAILABLE", "RESOURCE_EXHAUSTED"]
                }
            }
        ]
    }`
	serviceOptions := grpc.WithDefaultServiceConfig(serviceConfigJSONString)
	keepAliveConfig := keepalive.ClientParameters{
		Time:                10 * time.Second,
		Timeout:             240 * time.Second,
		PermitWithoutStream: true,
	}
	keepAliveOptions := grpc.WithKeepaliveParams(keepAliveConfig)
	c, err := grpc.Dial(
		address,
		grpc.WithInsecure(),
		connOptions,
		proxyOptions,
		msgOptions,
		serviceOptions,
		keepAliveOptions,
	)
	if err != nil {
		log.Fatalf("could not connect to %s: %v", address, err)
	}

	return c
}
