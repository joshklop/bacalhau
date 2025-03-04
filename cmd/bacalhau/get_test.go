//go:build integration

package bacalhau

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/filecoin-project/bacalhau/pkg/model"

	"github.com/filecoin-project/bacalhau/pkg/docker"
	"github.com/filecoin-project/bacalhau/pkg/system"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestGetSuite(t *testing.T) {
	suite.Run(t, new(GetSuite))
}

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type GetSuite struct {
	BaseSuite
}

// Before each test
func (s *GetSuite) SetupTest() {
	docker.MustHaveDocker(s.T())
	s.BaseSuite.SetupTest()
}

func testResultsFolderStructure(t *testing.T, baseFolder, hostID string) {
	files := []string{}
	err := filepath.Walk(baseFolder, func(path string, _ os.FileInfo, _ error) error {
		usePath := strings.Replace(path, baseFolder, "", 1)
		if usePath != "" {
			files = append(files, usePath)
		}
		return nil
	})
	require.NoError(t, err, "Error walking results directory")

	shortID := system.GetShortID(hostID)

	// if you change the docker run command in any way we need to change this
	resultsCID := "QmR92HM96X3seZEaRWXfRJDDFHKcprqbmQsEQ9uhbrA7MQ"

	expected := []string{
		"/" + model.DownloadVolumesFolderName,
		"/" + model.DownloadVolumesFolderName + "/data",
		"/" + model.DownloadVolumesFolderName + "/data/apples",
		"/" + model.DownloadVolumesFolderName + "/data/apples/file.txt",
		"/" + model.DownloadVolumesFolderName + "/data/file.txt",
		"/" + model.DownloadVolumesFolderName + "/outputs",
		"/" + model.DownloadVolumesFolderName + "/" + model.DownloadFilenameStderr,
		"/" + model.DownloadVolumesFolderName + "/" + model.DownloadFilenameStdout,
		"/" + model.DownloadShardsFolderName,
		"/" + model.DownloadShardsFolderName + "/0_node_" + shortID,
		"/" + model.DownloadShardsFolderName + "/0_node_" + shortID + "/data",
		"/" + model.DownloadShardsFolderName + "/0_node_" + shortID + "/data/apples",
		"/" + model.DownloadShardsFolderName + "/0_node_" + shortID + "/data/apples/file.txt",
		"/" + model.DownloadShardsFolderName + "/0_node_" + shortID + "/data/file.txt",
		"/" + model.DownloadShardsFolderName + "/0_node_" + shortID + "/exitCode",
		"/" + model.DownloadShardsFolderName + "/0_node_" + shortID + "/outputs",
		"/" + model.DownloadShardsFolderName + "/0_node_" + shortID + "/stderr",
		"/" + model.DownloadShardsFolderName + "/0_node_" + shortID + "/stdout",
		"/" + model.DownloadCIDsFolderName,
		"/" + model.DownloadCIDsFolderName + "/" + resultsCID,
		"/" + model.DownloadCIDsFolderName + "/" + resultsCID + "/data",
		"/" + model.DownloadCIDsFolderName + "/" + resultsCID + "/data/apples",
		"/" + model.DownloadCIDsFolderName + "/" + resultsCID + "/data/apples/file.txt",
		"/" + model.DownloadCIDsFolderName + "/" + resultsCID + "/data/file.txt",
		"/" + model.DownloadCIDsFolderName + "/" + resultsCID + "/exitCode",
		"/" + model.DownloadCIDsFolderName + "/" + resultsCID + "/outputs",
		"/" + model.DownloadCIDsFolderName + "/" + resultsCID + "/stderr",
		"/" + model.DownloadCIDsFolderName + "/" + resultsCID + "/stdout",
	}

	require.Equal(t, strings.Join(expected, "\n"), strings.Join(files, "\n"), "The discovered results output structure was not correct")
}

func testDownloadOutput(t *testing.T, cmdOutput, jobID, outputDir string) {
	require.True(t, strings.Contains(
		cmdOutput,
		fmt.Sprintf("Results for job '%s'", jobID),
	), "Job ID not found in output")
	require.True(t, strings.Contains(
		cmdOutput,
		fmt.Sprintf("%s", outputDir),
	), "Download location not found in output")

}

func setupTempWorkingDir(t *testing.T) (string, func()) {
	// switch wd to a temp dir so we are not writing folders to the current directory
	// (the point of this test is to see what happens when we DONT pass --output-dir)
	tempDir := t.TempDir()
	originalWd, err := os.Getwd()
	require.NoError(t, err)
	err = os.Chdir(tempDir)
	require.NoError(t, err)
	// Below looks redundant but it is necessary on Darwin which does some
	// freaky symlinking of root dirs: https://unix.stackexchange.com/a/63565
	newTempDir, err := os.Getwd()
	require.NoError(t, err)
	return newTempDir, func() {
		os.Chdir(originalWd)
	}
}

func (s *GetSuite) getDockerRunArgs(extraArgs []string) []string {
	swarmAddresses, err := s.node.IPFSClient.SwarmAddresses(context.Background())
	require.NoError(s.T(), err)
	args := []string{
		"docker", "run",
		"--api-host", s.host,
		"--api-port", s.port,
		"--ipfs-swarm-addrs", strings.Join(swarmAddresses, ","),
		"-o", "data:/data",
		"--wait",
	}
	args = append(args, extraArgs...)
	args = append(args,
		"ubuntu:latest",
		"--",
		"bash", "-c",
		"echo hello > /data/file.txt && echo hello && mkdir /data/apples && echo oranges > /data/apples/file.txt",
	)
	return args
}

// this tests that when we do docker run with no --output-dir
// it makes it's own folder to put the results in and does not splat results
// all over the current directory
func (s *GetSuite) TestDockerRunWriteToJobFolderAutoDownload() {
	tempDir, cleanup := setupTempWorkingDir(s.T())
	defer cleanup()

	args := s.getDockerRunArgs([]string{
		"--wait",
		"--download",
	})
	_, runOutput, err := ExecuteTestCobraCommand(s.T(), args...)
	require.NoError(s.T(), err, "Error submitting job")
	jobID := system.FindJobIDInTestOutput(runOutput)
	hostID := s.node.Host.ID().String()
	outputFolder := filepath.Join(tempDir, getDefaultJobFolder(jobID))
	testDownloadOutput(s.T(), runOutput, jobID, tempDir)
	testResultsFolderStructure(s.T(), outputFolder, hostID)

}

// this tests that when we do docker run with an --output-dir
// the results layout adheres to the expected folder layout
func (s *GetSuite) TestDockerRunWriteToJobFolderNamedDownload() {
	tempDir, err := os.MkdirTemp("", "docker-run-download-test")
	require.NoError(s.T(), err)

	args := s.getDockerRunArgs([]string{
		"--wait",
		"--download",
		"--output-dir", tempDir,
	})
	_, runOutput, err := ExecuteTestCobraCommand(s.T(), args...)
	require.NoError(s.T(), err, "Error submitting job")
	jobID := system.FindJobIDInTestOutput(runOutput)
	hostID := s.node.Host.ID().String()
	testDownloadOutput(s.T(), runOutput, jobID, tempDir)
	testResultsFolderStructure(s.T(), tempDir, hostID)
}

// this tests that when we do get with no --output-dir
// it makes it's own folder to put the results in and does not splat results
// all over the current directory
func (s *GetSuite) TestGetWriteToJobFolderAutoDownload() {
	swarmAddresses, err := s.node.IPFSClient.SwarmAddresses(context.Background())
	require.NoError(s.T(), err)
	tempDir, cleanup := setupTempWorkingDir(s.T())
	defer cleanup()

	args := s.getDockerRunArgs([]string{
		"--wait",
	})
	_, out, err := ExecuteTestCobraCommand(s.T(), args...)
	require.NoError(s.T(), err, "Error submitting job")
	jobID := system.FindJobIDInTestOutput(out)
	hostID := s.node.Host.ID().String()

	_, getOutput, err := ExecuteTestCobraCommand(s.T(), "get",
		"--api-host", s.node.APIServer.Address,
		"--api-port", fmt.Sprintf("%d", s.node.APIServer.Port),
		"--ipfs-swarm-addrs", strings.Join(swarmAddresses, ","),
		jobID,
	)
	require.NoError(s.T(), err, "Error getting results")

	testDownloadOutput(s.T(), getOutput, jobID, filepath.Join(tempDir, getDefaultJobFolder(jobID)))
	testResultsFolderStructure(s.T(), filepath.Join(tempDir, getDefaultJobFolder(jobID)), hostID)
}

// this tests that when we do get with an --output-dir
// the results layout adheres to the expected folder layout
func (s *GetSuite) TestGetWriteToJobFolderNamedDownload() {
	ctx := context.Background()
	swarmAddresses, err := s.node.IPFSClient.SwarmAddresses(ctx)
	require.NoError(s.T(), err)

	tempDir, err := os.MkdirTemp("", "docker-run-download-test")
	require.NoError(s.T(), err)

	args := s.getDockerRunArgs([]string{
		"--wait",
	})
	_, out, err := ExecuteTestCobraCommand(s.T(), args...)

	require.NoError(s.T(), err, "Error submitting job")
	jobID := system.FindJobIDInTestOutput(out)
	hostID := s.node.Host.ID().String()

	_, getOutput, err := ExecuteTestCobraCommand(s.T(), "get",
		"--api-host", s.node.APIServer.Address,
		"--api-port", fmt.Sprintf("%d", s.node.APIServer.Port),
		"--ipfs-swarm-addrs", strings.Join(swarmAddresses, ","),
		"--output-dir", tempDir,
		jobID,
	)
	require.NoError(s.T(), err, "Error getting results")
	testDownloadOutput(s.T(), getOutput, jobID, tempDir)
	testResultsFolderStructure(s.T(), tempDir, hostID)
}
