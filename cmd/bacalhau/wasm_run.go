package bacalhau

import (
	"fmt"

	"github.com/bytecodealliance/wasmtime-go"
	"github.com/filecoin-project/bacalhau/pkg/system"
	"github.com/filecoin-project/bacalhau/pkg/version"
	"github.com/spf13/cobra"
)

func init() { //nolint:gochecknoinits // idiomatic for cobra commands
	wasmCmd.AddCommand(runWasmCommand)
}

var wasmCmd = &cobra.Command{
	Use:   "wasm",
	Short: "Run and prepare WASM jobs on the network",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// Check that the server version is compatible with the client version
		serverVersion, _ := GetAPIClient().Version(cmd.Context()) // Ok if this fails, version validation will skip
		if err := ensureValidVersion(cmd.Context(), version.Get(), serverVersion); err != nil {
			Fatal(fmt.Sprintf("version validation failed: %s", err), 1)
			return err
		}

		return nil
	},
}

var runWasmCommand = &cobra.Command{
	Use:     "run",
	Short:   "Run a WASM job on the network",
	Long:    languageRunLong,
	Example: languageRunExample,
	Args:    cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		cm := system.NewCleanupManager()
		defer cm.Cleanup()

		ctx, rootSpan := system.NewRootSpan(cmd.Context(), system.GetTracer(), "cmd/bacalhau/wasm_run.runWasmCommand")
		defer rootSpan.End()
		cm.RegisterCallback(system.CleanupTraceProvider)

		programPath := args[0]
		OLR.ContextPath = programPath
		OLR.Command = args[1]

		return SubmitLanguageJob(cmd, ctx, "wasm", "2.0", programPath)
	},
}
