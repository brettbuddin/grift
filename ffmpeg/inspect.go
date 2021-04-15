package ffmpeg

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

// Inspect runs ffprobe on a file
func Inspect(ctx context.Context, env Environment, filename string) (ProbeResult, error) {
	buf := bytes.NewBuffer(nil)
	args := []string{
		"-hide_banner",
		"-print_format",
		"json",
		"-show_format",
		"-show_streams",
		filename,
	}
	fmt.Fprintf(env.Writer, "%s %s\n", env.Commands.FFProbe, strings.Join(args, " "))
	cmd := exec.CommandContext(ctx, env.Commands.FFProbe, args...)
	cmd.Stdout = buf
	cmd.Stderr = env.ErrWriter
	if err := cmd.Run(); err != nil {
		return ProbeResult{}, err
	}
	fmt.Fprintln(env.Writer)

	var res ProbeResult
	if err := json.NewDecoder(buf).Decode(&res); err != nil {
		return ProbeResult{}, err
	}
	return res, nil
}

// ProbeResult is the result of an ffprobe inspection
type ProbeResult struct {
	Format  ProbeFormat   `json:"format"`
	Streams []ProbeStream `json:"streams"`
}

// ProbeStream is an individal stream of channels inside an ffprobe result.
type ProbeStream struct {
	Index      int      `json:"index"`
	SampleRate Uint32   `json:"sample_rate"`
	Duration   Timecode `json:"duration"`
	Depth      Uint32   `json:"bits_per_raw_sample"`
}

// Timecode represents a duration in milliseconds used in ID3v2 tags. ffprobe
// results encodes these values as strings so this type properly unmarshals
// those values.
type Timecode uint32

func (c *Timecode) UnmarshalJSON(b []byte) error {
	var str string
	if err := json.Unmarshal(b, &str); err != nil {
		return err
	}
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return err
	}

	*c = Timecode(f * 1000)

	return nil
}

// ProbeFormat is format information presented in an ffprobe result
type ProbeFormat struct {
	NumStreams int      `json:"nb_streams"`
	FormatName string   `json:"format_name"`
	Duration   Timecode `json:"duration"`
}

// Uint32 is an unsigned integer that's encoded as a string in ffprobe results.
// This type properly unmarshals those strings into the appropriate uint32 type.
type Uint32 uint32

func (s *Uint32) UnmarshalJSON(b []byte) error {
	var str string
	if err := json.Unmarshal(b, &str); err != nil {
		return err
	}
	ui, err := strconv.ParseUint(str, 10, 32)
	if err != nil {
		return err
	}

	*s = Uint32(ui)

	return nil
}

type Float64 float64

func (f *Float64) UnmarshalJSON(b []byte) error {
	var str string
	if err := json.Unmarshal(b, &str); err != nil {
		return err
	}
	pf, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return err
	}

	*f = Float64(pf)

	return nil
}
