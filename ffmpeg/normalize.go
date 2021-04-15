package ffmpeg

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/brettbuddin/grift/manifest"
)

func normalize(ctx context.Context, env Environment, wavFile string, normalize manifest.Normalize) (string, error) {
	analysis, err := normalizeAnalyze(ctx, env, wavFile, normalize)
	if err != nil {
		return "", err
	}
	wavOut, err := normalizeApply(ctx, env, wavFile, normalize, analysis)
	if err != nil {
		return "", err
	}
	return wavOut, nil
}

type analysis struct {
	InputI       Float64 `json:"input_i"`
	InputTP      Float64 `json:"input_tp"`
	InputLRA     Float64 `json:"input_lra"`
	InputThresh  Float64 `json:"input_thresh"`
	TargetOffset Float64 `json:"target_offset"`
}

var parsedLoudNormPattern = regexp.MustCompile("\\[Parsed_loudnorm_\\d+ @ \\w+\\]\\s*")

func normalizeAnalyze(ctx context.Context, env Environment, wavFile string, normalize manifest.Normalize) (analysis, error) {
	buf := bytes.NewBuffer(nil)
	args := []string{
		"-hide_banner",
		"-y",
		"-i",
		wavFile,
		"-af",
		fmt.Sprintf(
			"loudnorm=I=%f:LRA=%f:TP=%f:print_format=json",
			normalize.LoudnessTarget,
			normalize.LoudnessRange,
			normalize.TruePeak,
		),
		"-f",
		"null",
		"-",
	}
	fmt.Fprintf(env.Writer, "%s %s\n", env.Commands.FFMpeg, strings.Join(args, " "))
	cmd := exec.CommandContext(ctx, env.Commands.FFMpeg, args...)
	cmd.Stderr = io.MultiWriter(buf, env.ErrWriter)
	if err := cmd.Run(); err != nil {
		return analysis{}, err
	}
	fmt.Fprintln(env.Writer)

	parts := parsedLoudNormPattern.Split(buf.String(), -1)
	if len(parts) < 1 {
		return analysis{}, fmt.Errorf("no json content found in analysis output")
	}

	var res analysis
	if err := json.NewDecoder(strings.NewReader(parts[len(parts)-1])).Decode(&res); err != nil {
		return analysis{}, err
	}
	return res, nil
}

func normalizeApply(ctx context.Context, env Environment, wavFile string, normalize manifest.Normalize, a analysis) (string, error) {
	output := filepath.Join(env.WorkspaceDir, "normalized.wav")
	err := ffmpeg(ctx, env,
		"-hide_banner",
		"-y",
		"-i",
		wavFile,
		"-af",
		fmt.Sprintf(
			"loudnorm=I=%.3f:LRA=%.3f:TP=%.3f:measured_I=%.3f:measured_LRA=%.3f:measured_TP=%.3f:measured_thresh=%.3f:offset=%.3f:linear=true:print_format=summary",
			normalize.LoudnessTarget,
			normalize.LoudnessRange,
			normalize.TruePeak,
			a.InputI,
			a.InputLRA,
			a.InputTP,
			a.InputThresh,
			a.TargetOffset,
		),
		output,
	)
	if err != nil {
		return "", err
	}
	return output, nil
}
