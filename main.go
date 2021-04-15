// grift is an compositor for creating chaptered MP3 files.
package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"os/signal"
	"path/filepath"
	"strings"

	"github.com/brettbuddin/grift/ffmpeg"
	"github.com/brettbuddin/grift/manifest"
)

func main() {
	var (
		manifestFile string
		outFile      string
	)
	set := flag.NewFlagSet("grift", flag.ExitOnError)
	set.StringVar(&outFile, "out", "output.mp3", "output file (.mp3 only)")
	set.StringVar(&manifestFile, "manifest", "episode.hcl", "episode manifest file (.hcl)")
	if err := set.Parse(os.Args[1:]); err != nil {
		exitError(err, 2)
	}

	if !strings.HasSuffix(outFile, ".mp3") {
		exitError("-out must be an .mp3 file", 2)
	}
	if outFile == "" {
		exitError("-out is required", 2)
	}
	if manifestFile == "" {
		exitError("-manifest is required", 2)
	}

	env, err := ffmpeg.NewEnvironment()
	if err != nil {
		exitError(err, 2)
	}
	if err := run(env, outFile, manifestFile); err != nil {
		exitError(err, 1)
	}
}

func exitError(err interface{}, code int) {
	fmt.Println(err)
	os.Exit(code)
}

func run(env ffmpeg.Environment, outFile, manifestFile string) error {
	defer env.Cleanup()

	episode, err := manifest.Parse(manifestFile)
	if err != nil {
		return fmt.Errorf("parse manifest: %s", err)
	}
	if err := validate(env, episode); err != nil {
		return fmt.Errorf("validate: %s", err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	tmpOut, err := ffmpeg.Composite(ctx, env, episode)
	if err != nil {
		return fmt.Errorf("composite: %s", err)
	}
	if err := copyFile(outFile, tmpOut); err != nil {
		return fmt.Errorf("copy output file: %s", err)
	}
	return nil
}

func validate(env ffmpeg.Environment, episode manifest.Episode) error {
	if episode.Title == "" {
		return fmt.Errorf("episode title is required")
	}
	if len(episode.Authors) == 0 {
		return fmt.Errorf("episode authors are required")
	}
	for _, a := range episode.Authors {
		if a == "" {
			return fmt.Errorf("episode author names cannot be empty")
		}
	}
	if episode.Album == "" {
		return fmt.Errorf("episode album is required")
	}
	if episode.Marker != "" {
		result, err := ffmpeg.Inspect(env, episode.Marker)
		if err != nil {
			return fmt.Errorf("inspect %q: %s", err)
		}
		if err := validateFile(result); err != nil {
			return fmt.Errorf("validate marker: %s", err)
		}
	}
	for _, ch := range episode.Chapters {
		result, err := ffmpeg.Inspect(env, ch.Filename)
		if err != nil {
			return fmt.Errorf("inspect %q: %s", err)
		}
		if err := validateFile(result); err != nil {
			return fmt.Errorf("validate %q: %s", ch.Title, err)
		}
	}
	return nil
}

func validateFile(ch ffmpeg.ProbeResult) error {
	if ch.Format.FormatName != "wav" {
		return fmt.Errorf("chapter is not a WAV file")
	}
	if ch.Format.NumStreams != 1 {
		return fmt.Errorf("only 1 stream per file is supported")
	}
	return nil
}

func copyFile(dest, src string) error {
	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	absDst, err := filepath.Abs(dest)
	if err != nil {
		return err
	}

	destination, err := os.Create(absDst)
	if err != nil {
		return err
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	return err
}
