# Mismatched Deposit Count Error — Recovery Procedure

## Overview

This runbook describes the recovery procedure for the `mismatched deposit count` error in `zkevm-bridge-service`. When this error occurs, the bridge service stops syncing and gets stuck in a loop, making deposits unavailable to claim.

## 1. Error Signature

The bridge service logs the following error pattern (`...` stands for anything):
```json
{"level":"error", "caller":"synchronizer/synchronizer.go:...",
 "msg":"networkID: ..., failed to store new deposit in the bridge tree,
 BlockNumber: , ... err: mismatched deposit count: ..., expected: ..."}

{"level":"warn", "caller":"synchronizer/synchronizer.go:...",
 "msg":"networkID: ..., error syncing blocks: mismatched deposit count: ..., expected: ..."}
```

Key fields to extract from the error log:

| Field | Description |
|---|---|
| `networkID` | The network ID; `0` means L1, `>0` means L2 |
| `BlockNumber` | The block number where the mismatch was detected — used for surgical recovery |
| `DepositCount (X)` | The deposit count reported by the chain for this block |
| `expected (Y)` | The deposit count the bridge service internally expected |
| `version` | Bridge service version — used to assess upgrade necessity |

---

## 2. Recovery Options

There are three possible fixes:

> **Note — L1 RPC switch (applies to any path)**
> If `networkID` is `0`, switching to a different L1 RPC provider may be enough to fix the problem. Edit the config file to set the new endpoint before resyncing:
> ```toml
> [Etherman]
> # Replace with a different, reliable L1 RPC endpoint
> L1URL = "https://<NEW_RPC_ENDPOINT>"
> ```

- **Path A**: The faster solution — it only requires resyncing a few blocks, not all of them.
- **Path B**: The most reliable solution, but also the slowest — it requires a full resync from genesis.


> ⚠️ **Important**
> Upgrading from v0.5.x to v0.6.x always requires a full DB resync (Path B).



## 3. Path A — Surgical DB Rollback (Skip Full Resync)

Use this path when a full resync is not practical due to large DB size or long sync times. This procedure was developed as a targeted fix that removes only the offending blocks from the DB.


> ⚠️ **Prerequisites**
> - If the gap is large (e.g., 80+ deposits), use Path B instead — the database state is too inconsistent for a surgical repair.
> - Always back up the database before executing any `DELETE` statement.

### Step 1 — Identify the NetworkID and choose a rollback point

From the error log, extract the `BlockNumber` and `networkID` fields:
```
# Error log example:
"networkID: 0, failed to store new deposit in the bridge tree, BlockNumber: 24563985, Deposit: {Id:2853084 LeafType:0 OriginalNetwork:0 OriginalAddress:0x0000000000000000000000000000000000000000 Amount:+47662292539520054918 DestinationNetwork:20 DestinationAddress:0xf70da97812CB96acDF810712Aa562db8dfA3dbEF DepositCount:238813 BlockID:3007655 BlockNumber:24563985 NetworkID:0 TxHash:0xb30440eef8281f841da91b187400777b39d45a7e96b44f8019c6a1c73f4a3deb Metadata:[] ReadyForClaim:false} err: mismatched deposit count: 238813, expected: 238812
```
The extracted values are:
-  `networkID: 0`: network **0** (L1)
-  `BlockNumber: 24563985`: the offending block number is **24563985**

The **rollback point** is the **offending block** minus a safety margin of 10,000:

```
# Set a safe rollback point (subtract 10,000 blocks as margin):
# SAFE_BLOCK = 24563985 - 10000 = 24553985
```

This gives the following values for the example:
| Field | Value |
| ------|-------|
| Rollback_Block | 24553985 |
| NetworkID | 0 |

### Step 2 — Stop the bridge service
Stop the service. The exact procedure depends on your infrastructure.

### Step 3 — Back up the bridge database
Make a backup using `pg_dump`:
```bash
pg_dump -h <DB_HOST> -U <DB_USER> -d <DB_NAME> \
  -F c -f bridge_db_backup_before_rollback_$(date +%Y%m%d_%H%M%S).dump
```

### Step 4 — Find the last safe block ID (verification step)

Run this query to confirm the last clean block ID before the rollback point:
```sql
SELECT id FROM sync.block
WHERE network_id = $NetworkID
AND block_num < $Rollback_Block
ORDER BY block_num DESC
LIMIT 1;

-- Example with SAFE_BLOCK = 24553985:
SELECT id FROM sync.block WHERE network_id=0 AND block_num < 24553985
ORDER BY block_num DESC LIMIT 1;
```
Save this value for reference.

### Step 5 — Delete offending block data
```sql
DELETE FROM sync.block
WHERE block_num >= $Rollback_Block
AND network_id = $NetworkID;

-- Example:
DELETE FROM sync.block WHERE block_num >= 24553985 AND network_id = 0;
```

> ⚠️ **Critical**
> - Double-check the `WHERE` clause before executing.
> - Double-check that `network_id = $NetworkID` is correct.
> - Verify the row count returned by the `DELETE`.

### Step 6 — Restart the bridge service

Start the bridge service. It will re-sync only the deleted blocks — a much shorter process than a full resync.

### Step 7 — Monitor and confirm recovery
- Check that there are no `ERROR` or `WARN` log entries.
- Confirm the deposit count error is gone.
- Confirm block sync is progressing past the previously failing block.

---

## 4. Path B — L1 RPC Switch + Full Resync
This is the most reliable recovery method, but it has the drawback of also being the slowest.

### Step 1 — Stop the bridge service
Stop the service. The exact procedure depends on your infrastructure.

### Step 2 — Back up the existing database
Make a backup using `pg_dump`:
```bash
pg_dump -h <DB_HOST> -U <DB_USER> -d <DB_NAME> \
  -F c -f bridge_db_backup_before_resync_$(date +%Y%m%d_%H%M%S).dump
```

### Step 3 — Delete the bridge database
Drop and recreate the bridge database to force a clean resync from genesis.
```bash
psql -h <DB_HOST> -U <DB_USER> -c "DROP DATABASE <DB_NAME>;"
psql -h <DB_HOST> -U <DB_USER> -c "CREATE DATABASE <DB_NAME>;"
```

### Step 4 — Restart the bridge service
Start the bridge service. It will begin a full resync from the genesis block. This may take several hours depending on chain history.

### Step 5 — Monitor the logs
Confirm the service is syncing cleanly without errors.

---

## 5. Known Scenarios & Notes

### Error persists after resync with the same L1 RPC
If the error reappears on a different block after resync, the L1 RPC is consistently returning bad data. Switch to a completely different RPC provider and resync again. This has been confirmed by multiple operators: the error only resolved after switching to a different RPC provider.

### Large gap between actual and expected deposit count
Gaps of 80+ (e.g., expected: 41838, got: 41918) indicate the DB state is significantly corrupted and surgical rollback is insufficient. Use Path B (full resync).

### Error occurs immediately after version upgrade from v0.5.x
Upgrading from v0.5.x to v0.6.x always requires a full DB resync. This is expected behavior, not a bug. Follow Path B.

### Error on fresh resync (clean DB)
If the error appears during a resync of a clean DB, the L1 RPC is the root cause. The bridge is reading inconsistent event data from the start. Switch the L1 RPC before resyncing again.

---

## 6. Escalation

If the error persists after completing the fix, escalate to the team with the following information:

- Full error log output (at least 200 lines surrounding the error)
- Bridge service version (`version` field from the error log)
- L1 RPC / L2 RPC endpoint(s) used (masked if sensitive)
- Whether a full resync was already attempted

---

## 7. Reference

| Resource | Details |
|---|---|
| **Recommended image** | `hermeznetwork/zkevm-bridge-service:v0.6.2` |
| **Min. supported version** | v0.6.2 (v0.5.x is deprecated) |
| **Public runbook for Running L2 Trusted Environment** | [agglayer/runbooks — run-l2-trusted-environment.md](https://github.com/agglayer/runbooks/blob/main/operations/run-l2-trusted-environment.md) |
| **Bridge API health check** | `<BRIDGE_URL>/bridges/<WALLET_ADDRESS>` |
