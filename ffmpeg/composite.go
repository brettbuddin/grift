package ffmpeg

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/brettbuddin/grift/ffmpeg/md"
	"github.com/brettbuddin/grift/manifest"
)

// Composite creates a chaptered MP3 for a given manifest.
func Composite(ctx context.Context, env Environment, episode manifest.Episode) (string, error) {
	wavOut, wavSegments, err := joinChapters(ctx, env, episode)
	if err != nil {
		return "", fmt.Errorf("join wav files: %s", err)
	}

	if episode.Normalize != nil {
		wavOut, err = normalize(ctx, env, wavOut, *episode.Normalize)
		if err != nil {
			return "", fmt.Errorf("normalize: %s", err)
		}
	}

	tmpOut, err := encodeMP3(ctx, env, episode, wavOut, wavSegments...)
	if err != nil {
		return "", fmt.Errorf("encode mp3 file: %s", err)
	}
	return tmpOut, nil
}

func joinChapters(ctx context.Context, env Environment, episode manifest.Episode) (string, []wavSegment, error) {
	if len(episode.Chapters) == 0 {
		return "", nil, fmt.Errorf("chapters required")
	}

	chapterSegments, err := prepareChapters(ctx, env, episode.Chapters...)
	if err != nil {
		return "", nil, fmt.Errorf("prepare chapters: %s", err)
	}

	var segments []wavSegment
	if episode.Marker != "" {
		marker, err := prepareMarker(ctx, env, episode.Marker)
		if err != nil {
			return "", nil, fmt.Errorf("prepare marker: %s", err)
		}
		for _, chs := range chapterSegments {
			segments = append(segments, marker, chs)
		}
	} else {
		for _, chs := range chapterSegments {
			segments = append(segments, chs)
		}
	}

	output, err := joinWAVSegments(ctx, env, segments...)
	if err != nil {
		return "", nil, fmt.Errorf("join wav files: %s", err)
	}
	return output, segments, nil
}

func encodeMP3(ctx context.Context, env Environment, episode manifest.Episode, wavFile string, segments ...wavSegment) (string, error) {
	output := filepath.Join(env.WorkspaceDir, "output.mp3")
	err := ffmpeg(ctx, env,
		"-hide_banner",
		"-y",
		"-i",
		wavFile,
		"-codec:a",
		"libmp3lame",
		"-ar",
		fmt.Sprintf("%d", episode.Samplerate),
		"-b:a",
		fmt.Sprintf("%dk", episode.Bitrate),
		output,
	)
	if err != nil {
		return "", err
	}

	var (
		mdChapters []md.Chapter
		offset     Timecode
	)
	for _, s := range segments {
		nextOffset := offset + s.details.Format.Duration
		if s.chapterNumber < 0 {
			offset = nextOffset
			continue
		}
		mdChapters = append(mdChapters, md.Chapter{
			Filename: s.filename,
			Title:    episode.Chapters[s.chapterNumber].Title,
			Start:    uint32(offset),
			End:      uint32(nextOffset),
		})
		offset = nextOffset
	}

	track := md.Track{
		Title:  episode.Title,
		Album:  episode.Album,
		Artist: strings.Join(episode.Authors, ", "),
	}
	if err := md.WriteChapters(output, track, mdChapters); err != nil {
		return "", fmt.Errorf("write chapter information: %s", err)
	}

	return output, nil
}

type wavSegment struct {
	chapterNumber int
	filename      string
	details       ProbeResult
}

func prepareChapters(ctx context.Context, env Environment, chapters ...manifest.Chapter) ([]wavSegment, error) {
	var segments []wavSegment
	for i, ch := range chapters {
		prepared := filepath.Join(env.WorkspaceDir, fmt.Sprintf("%d.wav", i))
		err := ffmpeg(ctx, env,
			"-hide_banner",
			"-y",
			"-i",
			ch.Filename,
			"-acodec",
			"pcm_s24le",
			"-ar",
			"48000",
			"-filter:a",
			fmt.Sprintf("volume=%.2fdB", ch.Gain),
			prepared,
		)
		if err != nil {
			return nil, err
		}
		info, err := Inspect(env, prepared)
		if err != nil {
			return nil, err
		}
		segments = append(segments, wavSegment{
			chapterNumber: i,
			filename:      prepared,
			details:       info,
		})
	}
	return segments, nil
}

func prepareMarker(ctx context.Context, env Environment, file string) (wavSegment, error) {
	prepared := filepath.Join(env.WorkspaceDir, "marker.wav")
	err := ffmpeg(ctx, env,
		"-hide_banner",
		"-y",
		"-i",
		file,
		"-acodec",
		"pcm_s24le",
		"-ar",
		"48000",
		prepared,
	)
	if err != nil {
		return wavSegment{}, err
	}
	info, err := Inspect(env, prepared)
	if err != nil {
		return wavSegment{}, err
	}
	return wavSegment{
		chapterNumber: -1,
		filename:      prepared,
		details:       info,
	}, nil
}

func joinWAVSegments(ctx context.Context, env Environment, segments ...wavSegment) (string, error) {
	chapterlistFile := filepath.Join(env.WorkspaceDir, "chapterlist.txt")
	if err := writeChapterList(chapterlistFile, segments...); err != nil {
		return "", fmt.Errorf("write chapterlist.txt: %s", err)
	}

	combined := filepath.Join(env.WorkspaceDir, "combined.wav")
	err := ffmpeg(ctx, env,
		"-hide_banner",
		"-y",
		"-f",
		"concat",
		"-safe",
		"0",
		"-i",
		chapterlistFile,
		"-c",
		"copy",
		combined,
	)
	if err != nil {
		return "", err
	}
	return combined, nil
}

func writeChapterList(outputFile string, segments ...wavSegment) error {
	f, err := os.OpenFile(outputFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	for _, ch := range segments {
		rel, err := filepath.Abs(ch.filename)
		if err != nil {
			return err
		}
		fmt.Fprintf(f, "file '%s'\n", rel)
	}
	return nil
}

func ffmpeg(ctx context.Context, env Environment, args ...string) error {
	fmt.Fprintf(env.Writer, "%s %s\n", env.Commands.FFMpeg, strings.Join(args, " "))
	cmd := exec.CommandContext(ctx, env.Commands.FFMpeg, args...)
	cmd.Stdout = env.Writer
	cmd.Stderr = env.ErrWriter
	if err := cmd.Run(); err != nil {
		return err
	}
	fmt.Fprintln(env.Writer)
	return nil
}
