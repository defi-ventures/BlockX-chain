<!--
parent:
  order: false
-->

<div align="center">
  <h1> Tokn </h1>
</div>

<div align="center">
  <a href="https://github.com/defi-ventures/ethermint/releases/latest">
    <img alt="Version" src="https://img.shields.io/github/tag/cosmos/ethermint.svg" />
  </a>
  <a href="https://github.com/defi-ventures/ethermint/blob/development/LICENSE">
    <img alt="License: Apache-2.0" src="https://img.shields.io/github/license/cosmos/ethermint.svg" />
  </a>
  <a href="https://pkg.go.dev/github.com/defi-ventures/ethermint?tab=doc">
    <img alt="GoDoc" src="https://godoc.org/github.com/defi-ventures/ethermint?status.svg" />
  </a>
  <a href="https://goreportcard.com/report/github.com/defi-ventures/ethermint">
    <img alt="Go report card" src="https://goreportcard.com/badge/github.com/defi-ventures/ethermint"/>
  </a>
  <a href="https://codecov.io/gh/cosmos/ethermint">
    <img alt="Code Coverage" src="https://codecov.io/gh/cosmos/ethermint/branch/development/graph/badge.svg" />
  </a>
</div>
<div align="center">
  <a href="https://github.com/defi-ventures/ethermint">
    <img alt="Lines Of Code" src="https://tokei.rs/b1/github/cosmos/ethermint" />
  </a>
  <a href="https://discord.gg/AzefAFd">
    <img alt="Discord" src="https://img.shields.io/discord/669268347736686612.svg" />
  </a>
  <a href="https://github.com/defi-ventures/ethermint/actions?query=workflow%3ABuild">
    <img alt="Build Status" src="https://github.com/defi-ventures/ethermint/workflows/Build/badge.svg" />
  </a>
  <a href="https://github.com/defi-ventures/ethermint/actions?query=workflow%3ALint">
    <img alt="Lint Status" src="https://github.com/defi-ventures/ethermint/workflows/Lint/badge.svg" />
  </a>
</div>

Tokn is a scalable, high-throughput Proof-of-Stake blockchain that is fully compatible and
interoperable with Ethereum. It's build using the the [Cosmos SDK](https://github.com/cosmos/cosmos-sdk/) which runs on top of [Tendermint Core](https://github.com/tendermint/tendermint) consensus engine.

> **WARNING:** Tokn is under VERY ACTIVE DEVELOPMENT and should be treated as pre-alpha software. This means it is not meant to be run in production, its APIs are subject to change without warning and should not be relied upon, and it should not be used to hold any value. We will remove this warning when we have a release that is stable, secure, and properly tested.

**Note**: Requires [Go 1.15+](https://golang.org/dl/)

## Quick Start

To learn how the Tokn works from a high-level perspective, go to the [Introduction](./docs/intro/overview.md) section from the documentation.

For more, please refer to the [Tokn Docs](./docs/), which are also hosted on [docs.ethermint.zone](https://docs.ethermint.zone/).

## Tests

Unit tests are invoked via:

```bash
make test
```

To run JSON-RPC tests, execute:

```bash
make test-rpc
```

There is also an included Ethereum mainnet exported blockchain file in `importer/blockchain`
that includes blocks up to height `97638`. To execute and test a full import of
these blocks using the EVM module, execute:

```bash
make test-import
```

You may also provide a custom blockchain export file to test importing more blocks
via the `--blockchain` flag. See `TestImportBlocks` for further documentation.

### Community

The following chat channels and forums are a great spot to ask questions about Tokn:

- [Cosmos Discord](https://discord.gg/W8trcGV)
- [Cosmos Forum](https://forum.cosmos.network)
