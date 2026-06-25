# ai-adk — documentation

Google Agent Development Kit (ADK) integration for togo

## Overview

Package adk bridges togo to the Google ADK agent runtime. Run your Google ADK
agents (typically Python) as a sidecar service exposing POST /run; togo calls
them over HTTP. Set ADK_BASE_URL. Blank-import to register.

## Install

```bash
togo install togo-framework/ai-adk
```

A capability plugin — it self-registers on boot; no driver selector needed.

## Configuration

Environment variables read by this plugin (extracted from the source — see the gateway/provider docs for each value):

| Env var |
|---|
| `ADK_BASE_URL"` |

## Usage

See the package API in the source.

## Links

- Marketplace: https://to-go.dev/marketplace
- Source: https://github.com/togo-framework/ai-adk
- Full README: ../README.md
