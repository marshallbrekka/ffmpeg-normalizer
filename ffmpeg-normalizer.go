package main

import "fmt"
import "flag"
import "path/filepath"
import "github.com/marshallbrekka/ffmpeg-normalizer/ffmpeg"

func main() {
	sourcePtr := flag.String("i", "fail", "The file to convert")
	destinationPtr := flag.String("o", "fail", "The destination file")
	audioStreamIndexPtr := flag.Int("audio-index", 0, "The audio stream index to use")
	flag.Parse()
	source, _ := filepath.Abs(*sourcePtr)
	destination, _ := filepath.Abs(*destinationPtr)
	fmt.Print("checking source\n")
	fmt.Print(source)
	fmt.Print("\nchecking source END")
	result := ffmpeg.Probe(source)
	streams := ffmpeg.ProbeToStreams(result)
	videoSettings := ffmpeg.Video(streams[0])
	idealAudio := ffmpeg.BestAudioStream(streams, *audioStreamIndexPtr)
	audioSettings := ffmpeg.Audio(idealAudio)
	fmt.Println(streams)
	fmt.Println(idealAudio)
	fmt.Println(videoSettings)
	fmt.Println(audioSettings)
	ffmpeg.Convert(source, destination, videoSettings, audioSettings)
}
