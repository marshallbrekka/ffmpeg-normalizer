package main

import "fmt"
import "os"
import "path/filepath"
import "github.com/marshallbrekka/ffmpeg-normalizer/ffmpeg"

func main() {
	source, _ := filepath.Abs(os.Args[1])
	destination, _ := filepath.Abs(os.Args[2])
	result := ffmpeg.Probe(source)
	streams := ffmpeg.ProbeToStreams(result)
	videoSettings := ffmpeg.Video(streams[0])
	idealAudio := ffmpeg.BestAudioStream(streams)
	audioSettings := ffmpeg.Audio(idealAudio)
	fmt.Println(streams)
	fmt.Println(idealAudio)
	fmt.Println(videoSettings)
	fmt.Println(audioSettings)
	ffmpeg.Convert(source, destination, videoSettings, audioSettings)
}
