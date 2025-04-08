# XLayer Bridge

XLayer is built on top of upstream repository with several custom features tailored
specifically for X Layer.


## Run with XLayer Erigon

Instead of running with [zkevm-node](https://github.com/0xPolygonHermez/zkevm-node) (which has been archived), we will
run locally with cdk-erigon to act as our L2.

First clone the repo.

```bash
git clone https://github.com/okx/xlayer-erigon.git
```
Then follow the instructions [here](https://github.com/okx/xlayer-erigon/tree/dev/test) to run xlayer-erigon.

Once xlayer-erigon is running, you can `make run-single` to run XLayer bridge as a single service. Alternatively,
you can also do:

```bash
# The local configs have already been tailored to match the configs in xlayer-erigon.
# You must run xlayer-erigon successfully first before performing these commands.
make run-api
make run-push 
make run-task
```

to run each XLayer sub-service independently.

## Testing

1. bridge-ui

You can clone [bridge-ui](https://github.com/0xPolygonHermez/zkevm-bridge-ui/tree/develop), and run the
bridge-ui locally to test (following the README.md).

You will need to import the relevant token before bridging any asset (using Metamask etc.).

```bash
# You can import this token to test the bridge locally.
L2 WETH Token: 0x17a2a2e444a7f3446877d1b71eaa2b2ae7533baf
```

2. e2e

SSH into a Linux machine, and run:

```
make test-full # Default e2e test
make test-edge # tests edge cases
```
