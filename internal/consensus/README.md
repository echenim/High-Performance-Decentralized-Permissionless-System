# Consensus Package

## Overview

The `consensus` package implements the mechanisms required to achieve agreement among nodes in a decentralized permissionless system. This ensures that all nodes in the network share a consistent view of the system's state, even in the presence of malicious actors or unreliable communication.

---

## Features

1. **Consensus Algorithms**:
   - **Proof-of-Stake (PoS)**:
     - Nodes propose and validate blocks by staking assets, reducing energy consumption compared to Proof-of-Work (PoW).
   - **Delegated Proof-of-Stake (DPoS)**:
     - A subset of nodes (validators) is elected by stakeholders to propose and validate blocks, improving performance.

2. **Block Proposal and Voting**:
   - Validators propose blocks containing transactions.
   - Other validators vote to achieve consensus on the block to append to the blockchain.

3. **Finality**:
   - Implements Practical Byzantine Fault Tolerance (PBFT) to ensure blocks are finalized and cannot be reverted.

4. **Security Measures**:
   - Protection against double-spending and Sybil attacks.
   - Randomized validator selection to enhance fairness and prevent collusion.

---

## Structure

The package is organized as follows:

```plaintext
internal/consensus/
├── consensus.go          # Core logic for consensus algorithms
├── pbft.go               # Practical Byzantine Fault Tolerance implementation
├── pos.go                # Proof-of-Stake consensus logic
├── dpos.go               # Delegated Proof-of-Stake consensus logic
├── README.md             # Documentation for the consensus package
```