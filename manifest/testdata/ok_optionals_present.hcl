title = "Weird Sounds"
album = "Phonogrifter"
authors = ["Brett Buddin"]

samplerate = 48000
bitrate = 256

marker = "beep.wav"

chapter "ZOOM0012_TrLR.WAV" {
    title = "Chapter 1"
}

chapter "ZOOM0013_TrLR.WAV" {
    title = "Chapter 2"
    gain = -2
    start = 3
    stop = 8
}

normalize {
    loudness_target = -24 
    loudness_range = 7
    true_peak = -2
}
