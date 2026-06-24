// Package adk bridges togo to the Google ADK agent runtime. Run your Google ADK
// agents (typically Python) as a sidecar service exposing POST /run; togo calls
// them over HTTP. Set ADK_BASE_URL. Blank-import to register.
package adk

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/togo-framework/togo"
)

// RunRequest invokes an agent with an input + optional session.
type RunRequest struct {
	Agent   string `json:"agent,omitempty"`
	Input   string `json:"input"`
	Session string `json:"session,omitempty"`
}

// RunResponse is the agent's result.
type RunResponse struct {
	Output string          `json:"output"`
	Raw    json.RawMessage `json:"raw,omitempty"`
}

// Service bridges to the Google ADK runtime.
type Service struct {
	base   string
	client *http.Client
}

func init() {
	togo.RegisterProviderFunc("ai-adk", togo.PriorityService, func(k *togo.Kernel) error {
		base := os.Getenv("ADK_BASE_URL")
		if base == "" {
			return nil // not configured — skip
		}
		k.Set("ai-adk", &Service{base: base, client: &http.Client{Timeout: 120 * time.Second}})
		return nil
	})
}

// FromKernel returns the bridge service bound to the kernel.
func FromKernel(k *togo.Kernel) (*Service, bool) {
	v, ok := k.Get("ai-adk")
	if !ok {
		return nil, false
	}
	s, ok := v.(*Service)
	return s, ok
}

// Run invokes an agent on the Google ADK runtime.
func (s *Service) Run(ctx context.Context, req RunRequest) (RunResponse, error) {
	if s.base == "" {
		return RunResponse{}, errors.New("ai-adk: ADK_BASE_URL not set")
	}
	buf, _ := json.Marshal(req)
	r, err := http.NewRequestWithContext(ctx, http.MethodPost, s.base+"/run", bytes.NewReader(buf))
	if err != nil {
		return RunResponse{}, err
	}
	r.Header.Set("Content-Type", "application/json")
	resp, err := s.client.Do(r)
	if err != nil {
		return RunResponse{}, err
	}
	defer resp.Body.Close()
	raw, _ := io.ReadAll(resp.Body)
	if resp.StatusCode >= 300 {
		return RunResponse{}, fmt.Errorf("ai-adk: %s: %s", resp.Status, string(raw))
	}
	var out RunResponse
	if err := json.Unmarshal(raw, &out); err != nil {
		out = RunResponse{Output: string(raw)}
	}
	return out, nil
}

// Handler proxies POST /run to the agent runtime. Mount under /api/ai/adk.
func Handler(k *togo.Kernel) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /run", func(w http.ResponseWriter, r *http.Request) {
		svc, ok := FromKernel(k)
		if !ok {
			http.Error(w, "ai-adk not configured (set ADK_BASE_URL)", http.StatusInternalServerError)
			return
		}
		var req RunRequest
		_ = json.NewDecoder(r.Body).Decode(&req)
		res, err := svc.Run(r.Context(), req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(res)
	})
	return mux
}
