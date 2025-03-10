package operations

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/0xPolygonHermez/zkevm-bridge-service/utils"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/health/grpc_health_v1"
)

const (
	defaultInterval = 1 * time.Second
	defaultDeadline = 60 * time.Second
	// DefaultInterval is a time interval
	DefaultInterval = 2 * time.Millisecond
	// DefaultDeadline is a time interval
	DefaultDeadline = 2 * time.Minute
)

var (
	// ErrTimeoutReached is thrown when the timeout is reached and
	// because the condition is not matched
	ErrTimeoutReached = fmt.Errorf("timeout has been reached")
)

func poll(interval, deadline time.Duration, condition ConditionFunc) error {
	return Poll(interval, deadline, condition)
}

// WaitRestHealthy waits for a rest endpoint to be ready
func WaitRestHealthy(address string) error {
	return poll(defaultInterval, defaultDeadline, func() (bool, error) {
		return restHealthyCondition(address)
	})
}

func restHealthyCondition(address string) (bool, error) {
	resp, err := http.Get(address + "/healthz")

	return resp.StatusCode == http.StatusOK, err
}

// WaitGRPCHealthy waits for a gRPC endpoint to be responding according to the
// health standard in package grpc.health.v1
func WaitGRPCHealthy(address string) error {
	return Poll(DefaultInterval, DefaultDeadline, func() (bool, error) {
		return grpcHealthyCondition(address)
	})
}

func (m *Manager) networkUpCondition() (bool, error) {
	return NodeUpCondition(m.cfg.L1NetworkURL)
}

func proverUpCondition() (bool, error) {
	return true, nil
	// return ops.ProverUpCondition()
}

func (m *Manager) zkevmNodeUpCondition() (done bool, err error) {
	return NodeUpCondition(m.cfg.L2NetworkURL)
}

func bridgeUpCondition() (done bool, err error) {
	res, err := http.Get("http://localhost:8080/healthz")
	if err != nil {
		// we allow connection errors to wait for the container up
		return false, nil
	}
	if res.Body != nil {
		defer func() {
			err = res.Body.Close()
		}()
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return false, err
	}
	r := struct {
		Status string
	}{}
	err = json.Unmarshal(body, &r)
	if err != nil {
		return false, err
	}
	done = r.Status == "SERVING"
	return done, nil
}

// WaitTxToBeMined waits until a tx is mined or forged.
func WaitTxToBeMined(ctx context.Context, client *ethclient.Client, tx *types.Transaction, timeout time.Duration) error {
	return utils.WaitTxToBeMined(ctx, client, tx, timeout)
}

// ConditionFunc is a generic function
type ConditionFunc func() (done bool, err error)

// Poll retries the given condition with the given interval until it succeeds
// or the given deadline expires.
func Poll(interval, deadline time.Duration, condition ConditionFunc) error {
	timeout := time.After(deadline)
	tick := time.NewTicker(interval)

	for {
		select {
		case <-timeout:
			return ErrTimeoutReached
		case <-tick.C:
			ok, err := condition()
			if err != nil {
				return err
			}
			if ok {
				return nil
			}
		}
	}
}

func grpcHealthyCondition(address string) (bool, error) {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	conn, err := grpc.NewClient(address, opts...)
	if err != nil {
		// we allow connection errors to wait for the container up
		return false, nil
	}
	defer func() {
		err = conn.Close()
	}()

	healthClient := grpc_health_v1.NewHealthClient(conn)
	state, err := healthClient.Check(context.Background(), &grpc_health_v1.HealthCheckRequest{})
	if err != nil {
		// we allow connection errors to wait for the container up
		return false, nil
	}

	done := state.Status == grpc_health_v1.HealthCheckResponse_SERVING

	return done, nil
}

// NodeUpCondition check if the container is up and running
func NodeUpCondition(target string) (bool, error) {
	var jsonStr = []byte(`{"jsonrpc":"2.0","method":"eth_syncing","params":[],"id":1}`)
	req, err := http.NewRequest(
		"POST", target,
		bytes.NewBuffer(jsonStr))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		// we allow connection errors to wait for the container up
		return false, nil
	}

	if res.Body != nil {
		defer func() {
			err = res.Body.Close()
		}()
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return false, err
	}

	r := struct {
		Result bool
	}{
		Result: true,
	}
	err = json.Unmarshal(body, &r)
	if err != nil {
		return false, err
	}

	done := !r.Result

	return done, nil
}
