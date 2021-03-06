package commands_test

import (
	"bytes"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"
	"github.com/spf13/cobra"

	"github.com/buildpacks/pack/internal/commands"
	"github.com/buildpacks/pack/internal/commands/fakes"
	"github.com/buildpacks/pack/internal/commands/testmocks"
	"github.com/buildpacks/pack/internal/config"
	ilogging "github.com/buildpacks/pack/internal/logging"
	"github.com/buildpacks/pack/logging"
	h "github.com/buildpacks/pack/testhelpers"
)

func TestBuildpackCommand(t *testing.T) {
	spec.Run(t, "BuildpackCommand", testBuildpackCommand, spec.Parallel(), spec.Report(report.Terminal{}))
}

func testBuildpackCommand(t *testing.T, when spec.G, it spec.S) {
	var (
		cmd        *cobra.Command
		logger     logging.Logger
		outBuf     bytes.Buffer
		mockClient *testmocks.MockPackClient
	)

	it.Before(func() {
		logger = ilogging.NewLogWithWriters(&outBuf, &outBuf)
		mockController := gomock.NewController(t)
		mockClient = testmocks.NewMockPackClient(mockController)
		cmd = commands.NewBuildpackCommand(logger, config.Config{}, mockClient, fakes.NewFakePackageConfigReader())
		cmd.SetOut(logging.GetWriterForLevel(logger, logging.InfoLevel))
	})

	when("buildpack", func() {
		it("prints help text", func() {
			cmd.SetArgs([]string{})
			h.AssertNil(t, cmd.Execute())
			output := outBuf.String()
			h.AssertContains(t, output, "Interact with buildpacks")
			h.AssertContains(t, output, "Usage:")
			h.AssertContains(t, output, "package")
			for _, command := range []string{"register", "yank", "pull"} {
				h.AssertNotContains(t, output, command)
			}
		})

		it("only shows experimental commands if in the config", func() {
			cmd = commands.NewBuildpackCommand(logger, config.Config{Experimental: true}, mockClient, fakes.NewFakePackageConfigReader())
			cmd.SetOut(logging.GetWriterForLevel(logger, logging.InfoLevel))
			cmd.SetArgs([]string{})
			h.AssertNil(t, cmd.Execute())
			output := outBuf.String()
			h.AssertContains(t, output, "Interact with buildpacks")
			h.AssertContains(t, output, "Usage:")
			for _, command := range []string{"package", "register", "yank", "pull"} {
				h.AssertContains(t, output, command)
			}
		})
	})
}
