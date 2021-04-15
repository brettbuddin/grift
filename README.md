# grift

`grift` composites WAV files into a chaptered MP3 file.

The program is used to build episodes of my podcast: [Phonogrifter](https://phonogrifter.buddin.org).

## Installation

```
go get github.com/brettbuddin/grift
```

## Usage

Create a directory for your podcast episode and create an `episode.hcl` file
inside of it that looks like this:

```hcl
title = "Weird Sounds"
album = "Phonogrifter"
authors = ["Brett Buddin"]

// (Optional) Sample Rate for the output. Default is 44100.
samplerate = 48000

// (Optional) Bit Rate (constant) for the output. Default is 128.
bitrate = 256

// (Optional) Separator between parts of the episode. This file will be played
// each time a chapter begins. Path can be relative or absolute.
marker = "beep.wav"

// Chapter of the episode. The label is a path to the WAV file. Path can be
// relative or absolute.
chapter "ZOOM0012_TrLR.WAV" {
    // Title of the chapter.
    title = "Chapter 1"
}

chapter "ZOOM0013_TrLR.WAV" {
    title = "Chapter 2"

    // (Optional) Gain to be applied to this chapter in dB.
    gain = -2
}

// (Optional) Normalization options. Search the internet for recommendations for
// these values. No normalization will be applied if this block is not
// specified. Normalization is applied using a two passes before MP3 compression
// is applied.
normalize {
    // Integrated loudness target (in LUFS)
    loudness_target = -24 

    // Loudness range target (in dB)
    loudness_range = 7

    // Maximum True Peak (in dB)
    true_peak = -2
}
```

Move (`cd`) into this episode's directory and run `grift` to render the episode.
By default, a file called `output.mp3` will be produced in the current
directory.
