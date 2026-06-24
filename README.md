# ai-adk — Google Agent Development Kit (ADK) integration for togo

Bridges togo to the **Google ADK** agent runtime. Run your Google ADK agents as a sidecar exposing `POST /run`; togo calls them over HTTP. Set `ADK_BASE_URL`.

```bash
togo install togo-framework/ai-adk
```

```go
svc, _ := adk.FromKernel(k)
res, _ := svc.Run(ctx, adk.RunRequest{Agent: "researcher", Input: "..."})
```

Mount `Handler(k)` under `/api/ai/adk`. MIT
