# ai-adk — documentation

  <img src=".github/assets/togo-mark.svg" alt="togo" height="64" />

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

Environment variables read by this plugin (extracted from the source):

| Env var | Notes |
|---|---|
| `ADK_BASE_URL` | _see provider docs_ |
| `G` | _see provider docs_ |

## Usage

See the package API in the source.

## Links

- Marketplace: https://to-go.dev/marketplace
- Source: https://github.com/togo-framework/ai-adk
- README: ../README.md
