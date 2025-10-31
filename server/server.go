package server

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/0xPolygon/zkevm-bridge-service/bridgectrl/pb"
	"github.com/0xPolygon/zkevm-bridge-service/log"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/protobuf/encoding/protojson"
)

// RunServer runs gRPC server and HTTP gateway
func RunServer(ctx context.Context, cfg Config, bridgeService pb.BridgeServiceServer) error {
	if len(cfg.GRPCPort) == 0 {
		return fmt.Errorf("invalid TCP port for gRPC server: '%s'", cfg.GRPCPort)
	}

	if len(cfg.HTTPPort) == 0 {
		return fmt.Errorf("invalid TCP port for HTTP gateway: '%s'", cfg.HTTPPort)
	}

	go func() {
		_ = runRestServer(ctx, cfg.GRPCPort, cfg.HTTPPort)
	}()

	go func() {
		_ = runGRPCServer(ctx, bridgeService, cfg.GRPCPort)
	}()

	return nil
}

// HealthChecker will provide an implementation of the HealthCheck interface.
type healthChecker struct{}

// NewHealthChecker returns a health checker according to standard package
// grpc.health.v1.
func newHealthChecker() *healthChecker {
	return &healthChecker{}
}

// HealthCheck interface implementation.

// Check returns the current status of the server for unary gRPC health requests,
// for now if the server is up and able to respond we will always return SERVING.
func (s *healthChecker) Check(ctx context.Context, req *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	return &grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	}, nil
}

// Watch returns the current status of the server for stream gRPC health requests,
// for now if the server is up and able to respond we will always return SERVING.
func (s *healthChecker) Watch(req *grpc_health_v1.HealthCheckRequest, server grpc_health_v1.Health_WatchServer) error {
	return server.Send(&grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	})
}

func runGRPCServer(ctx context.Context, bridgeServer pb.BridgeServiceServer, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	pb.RegisterBridgeServiceServer(server, bridgeServer)

	healthService := newHealthChecker()
	grpc_health_v1.RegisterHealthServer(server, healthService)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			server.GracefulStop()
			<-ctx.Done()
		}
	}()

	log.Info("gRPC Server is serving at ", port)
	return server.Serve(listen)
}

func preflightHandler(w http.ResponseWriter, r *http.Request) {
	headers := []string{"Content-Type", "Accept"}
	w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
	methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE"}
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
}

// allowCORS allows Cross Origin Resource Sharing from any origin.
// Don't do this without consideration in production systems.
func allowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
				preflightHandler(w, r)
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}

func runRestServer(ctx context.Context, grpcPort, httpPort string) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	endpoint := "localhost:" + grpcPort
	conn, err := grpc.NewClient(endpoint, opts...)
	if err != nil {
		return err
	}

	muxHealthOpt := runtime.WithHealthzEndpoint(grpc_health_v1.NewHealthClient(conn))
	muxJSONOpt := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames:   true,
			EmitUnpopulated: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})
	mux := runtime.NewServeMux(muxJSONOpt, muxHealthOpt)

	if err := pb.RegisterBridgeServiceHandler(ctx, mux, conn); err != nil {
		return err
	}

	// Whitelist handler - only allow valid API routes defined in proto
	gatewayHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		
		// Define valid API routes based on proto file + grpc-gateway endpoints
		validRoutes := map[string]bool{
			"/api":                    true,
			"/sync":                   true,
			"/merkle-proof":           true,
			"/merkle-proof-by-ger":    true,
			"/bridge":                 true,
			"/tokenwrapped":           true,
			"/pending-bridges":        true,
			"/v2/merkle-proof":        true,
			"/healthz":                true, // Automatically added by runtime.WithHealthzEndpoint
		}
		
		// Check if path matches valid routes or starts with valid prefixes
		isValid := validRoutes[path] || 
			strings.HasPrefix(path, "/bridges/") ||
			strings.HasPrefix(path, "/claims/")
		
		if !isValid {
			// Return 404 for invalid routes without involving grpc-gateway
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Content-Type", "application/json")
			
			if r.Method == "OPTIONS" {
				preflightHandler(w, r)
				w.WriteHeader(http.StatusOK)
				return
			}
			
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte(`{"error":"Not Found","path":"` + path + `"}`))
			return
		}
		
		// Valid routes go to grpc-gateway with CORS
		allowCORS(mux).ServeHTTP(w, r)
	})

	srv := &http.Server{
		ReadTimeout: 1 * time.Second, //nolint:mnd
		Addr:        ":" + httpPort,
		Handler:     gatewayHandler,
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			_ = srv.Shutdown(ctx)
			<-ctx.Done()
		}

		_, cancel := context.WithTimeout(ctx, 5*time.Second) //nolint:mnd
		defer cancel()

		_ = srv.Shutdown(ctx)
	}()

	log.Info("Restful Server is serving at ", httpPort)
	return srv.ListenAndServe()
}
