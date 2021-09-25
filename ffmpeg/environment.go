package ffmpeg

import (
	"io"
	"io/ioutil"
	"os"
	"os/exec"
)

// Environment contains dependency paths (executables, temp directory, etc.)
type Environment struct {
	Commands     CommandPaths
	WorkspaceDir string
	Writer       io.Writer
	ErrWriter    io.Writer
}

// NewEnvironment returns a new Paths
func NewEnvironment() (Environment, error) {
	cmds, err := locateCommands()
	if err != nil {
		return Environment{}, err
	}
	tmp, err := ioutil.TempDir("", "grift-*")
	if err != nil {
		return Environment{}, err
	}
	return Environment{
		Commands:     cmds,
		WorkspaceDir: tmp,
		Writer:       os.Stdout,
		ErrWriter:    os.Stderr,
	}, nil
}

// Cleanup removes any temporary files used to produce the output
func (e Environment) Cleanup() error {
	return os.RemoveAll(e.WorkspaceDir)
}

// CommandPaths contains executable paths to ffmpeg dependencies
type CommandPaths struct {
	FFMpeg, FFProbe string
}

// locateCommands locates the executables required by this tool
func locateCommands() (CommandPaths, error) {
	// Check for ffmpeg
	ffmpegPath, err := exec.LookPath("ffmpeg")
	if err != nil {
		return CommandPaths{}, err
	}

	ffprobePath, err := exec.LookPath("ffprobe")
	if err != nil {
		return CommandPaths{}, err
	}

	return CommandPaths{
		FFMpeg:  ffmpegPath,
		FFProbe: ffprobePath,
	}, nil
}
