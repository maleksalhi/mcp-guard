package main

import (
	"flag"
	"log/slog"
	"os"
)

func main() {
	policyPath := flag.String("policy", "policy.yaml", "path to policy file")
	flag.Parse()

	logger := slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	slog.SetDefault(logger)

	args := flag.Args()
	if len(args) == 0 {
		slog.Error("missing target MCP server command", "usage", "mcp-guard --policy policy.yaml -- <cmd> [args...]")
		os.Exit(1)
	}

	slog.Info("mcp-guard initialized", "policy", *policyPath, "target_cmd", args[0])
}
