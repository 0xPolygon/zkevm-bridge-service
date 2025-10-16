# zkEVM Bridge Service

[![License](https://img.shields.io/github/license/0xPolygonHermez/zkevm-bridge-service)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/0xPolygonHermez/zkevm-bridge-service)](https://goreportcard.com/report/github.com/0xPolygonHermez/zkevm-bridge-service)
[![GitHub release](https://img.shields.io/github/release/0xPolygonHermez/zkevm-bridge-service.svg)](https://github.com/0xPolygonHermez/zkevm-bridge-service/releases)

A robust, high-performance backend service written in Go that enables seamless asset bridging between Ethereum L1 and L2 networks connected to the Agglayer. This service provides the core infrastructure for cross-chain transfers by generating and managing Merkle proofs, monitoring transactions, and handling claim operations.

## 🌟 Features

- **Cross-Chain Asset and Message Transfers**: Bridge ERC-20 tokens, ETH and messages between L1 and L2 networks
- **Merkle Proof Generation**: Automated generation of cryptographic proofs for secure cross-chain claims
- **Transaction Monitoring**: Real-time monitoring and management of bridge transactions
- **Claim Management**: Automated and manual claim processing with retry mechanisms
- **Multi-Network Support**: Support for hundreds of L2 networks simultaneously
- **RESTful API**: HTTP/gRPC API for integration with frontend applications
- **Database Persistence**: PostgreSQL/SQLite support for reliable data storage
- **Metrics & Monitoring**: Prometheus metrics integration for operational visibility
- **Auto-Claim Service**: Automated claiming for authorized addresses
- **Grouping Claims**: Transaction batching for gas optimization

## 🏗️ Architecture

<p align="center">
  <img src="./docs/architecture.drawio.png" alt="zkEVM Bridge Service Architecture"/>
</p>

The service consists of several key components:

- **Synchronizer**: Monitors L1/L2 events and maintains state consistency
- **Claim Transaction Manager**: Handles claim transaction lifecycle and monitoring
- **Bridge Controller**: Manages bridge operations and Merkle tree operations
- **API Server**: Provides REST/gRPC endpoints for client interactions
- **Database Layer**: Persistent storage for deposits, claims, and metadata

## 🚀 Quick Start

### Prerequisites

- Go 1.24+ 
- Docker & Docker Compose
- PostgreSQL (or SQLite for development)
- Access to Ethereum L1 and L2 nodes (Erigon or OP-Geth)

### Installation

```bash
# Clone the repository
git clone https://github.com/0xPolygon/zkevm-bridge-service.git
cd zkevm-bridge-service

# Build the service
make build

# Or build the Docker image
make build-docker
```

### Configuration

1. Copy the example configuration:
```bash
cp config/config.local.toml config/config.toml
```

2. Update the configuration with your network parameters, database details and urls:
```toml
[SyncDB]
  Database = "postgres"
  [SyncDB.PgStorage]
    User = "test_user"
    Password = "test_password"
    Name = "test_db"
    Host = "zkevm-bridge-db"
    Port = "5432"
    MaxConns = 20

[BridgeServer]
  [BridgeServer.DB]
    Database = "postgres"
    [BridgeServer.DB.PgStorage]
    User = "test_user"
    Password = "test_password"
    Name = "test_db"
    Host = "zkevm-bridge-db"
    Port = "5432"
    MaxConns = 20

[Etherman]
  L1URL = "http://zkevm-mock-l1-network:8545"
  L2URLs = ["http://zkevm-node:8123"]

[NetworkConfig]
  GenBlockNumber = 0
  PolygonBridgeAddress = "0xFe12ABaa190Ef0c8638Ee0ba9F828BF41368Ca0E"
  PolygonZkEVMGlobalExitRootAddress = "0x8A791620dd6260079BF849Dc5567aDC3F2FdC318"
  PolygonRollupManagerAddress = "0xB7f8BC63BbcaD18155201308C8f3540b07f84F5e"
  L2ClaimCompressorAddress = "0x2279B7A0a67DB372996a5FaB50D91eAA73d2eBe6"
  L2PolygonBridgeAddresses = ["0xFe12ABaa190Ef0c8638Ee0ba9F828BF41368Ca0E"]
  RequireSovereignChainSmcs = [false]
  L2PolygonZkEVMGlobalExitRootAddresses = ["0xa40d5f56745a118d0906a34e69aec8c0db1cb8fa"]
```

### Running the Service

#### Using Docker Compose (Recommended)
```bash
# Start all services including databases
make run
```

#### Manual Execution
```bash
# Start the database
make run-db-bridge

# Start the bridge service
./bin/zkevm-bridge run --cfg config/config.toml
```

## 📖 Usage

### REST API

The service exposes several REST endpoints:

```bash
# Get bridge information
GET /api/v1/bridges/{address}

# Get claim merkle proof
GET /merkle-proof?deposit_cnt=1&net_id=0

# Health check
GET /api

## Note: Check the proto file to get all the details (proto/src/proto/bridge/v1/query.proto)
```

### CLI Commands

```bash
# Show version
./bin/zkevm-bridge version

# Run the service
./bin/zkevm-bridge run --cfg /app/config.toml
```

## 🧪 Testing

### Unit Tests
```bash
make test
```

### Integration Tests
```bash
# Run integration tests
make test-full
```

### E2E Tests for Real Networks
```bash
# Build E2E test Docker image
make build-docker-e2e-real_network

# Run tests against real networks
docker run --volume "./config/:/config/" \
  --env BRIDGE_TEST_CONFIG_FILE=/config/test.toml \
  bridge-e2e-realnetwork-erc20
```

## 🔧 Development

### Project Structure
```
├── autoclaimservice/    # Standalone automated claim service
├── bridgectrl/          # Bridge controller logic
├── claimtxman/          # Autoclaim transaction management
├── cmd/                 # CLI application entry point
├── config/              # Configuration management
├── db/                  # Database migrations and schemas
├── etherman/            # Ethereum client interactions
├── server/              # HTTP/gRPC API server
├── synchronizer/        # L1/L2 event synchronization
├── test/                # Test suites and utilities
└── utils/               # Common utilities
```

### Building from Source
```bash
# Install dependencies
go mod download

# Generate code (protobuf, mocks, etc.)
make generate-code-from-proto generate-mocks

# Build binary
make build

# Build docker image
make build-docker

# Run linting
make lint

# Run Benchmark
make bench bench-full

# Run local environment
make run

# Stop local environment
make stop

## Note: Check the make file to see all the options
```


## 📊 Monitoring & Metrics

The service exposes Prometheus metrics on `/metrics` endpoint:

- Bridge transaction counters
- Processing latency histograms
- Error rates and types
- Database connection metrics
- Node synchronization status

Example metrics:
```
bridge_deposits_total{network_id="1"}
bridge_claims_total{status="confirmed"}
bridge_processing_duration_seconds
```

## 🔐 Security

- All cross-chain transfers are secured by cryptographic Merkle proofs
- Database credentials should be stored securely (use environment variables)
- API endpoints support rate limiting and authentication
- Regular security audits are performed on smart contracts

## 📚 Documentation

- [Local Development Setup](docs/running_local.md)
- [E2E Testing Guide](docs/e2e-realnetwork-test.md)
- [API Reference](docs/api-reference.md)
- [Configuration Guide](docs/configuration.md)

## 🆘 Support

- **Documentation**: Check our [docs](docs/) directory
- **Issues**: Report bugs via [GitHub Issues](https://github.com/0xPolygonHermez/zkevm-bridge-service/issues)
- **Discord**: Join the [Polygon zkEVM Discord](https://discord.gg/0xPolygonRnD) for community support
- **Forum**: Visit the [Polygon Community Forum](https://forum.polygon.technology/)


---

<p align="center">
  Made with ❤️ by the <a href="https://polygon.technology/">Polygon</a> team
</p>
