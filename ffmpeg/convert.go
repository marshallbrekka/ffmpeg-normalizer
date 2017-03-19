package ffmpeg

import "os/exec"
import "fmt"

var video_flags = map[string]string{
	"codec":       "-c:v",
	"bit_rate":    "-b:v",
	"dimmensions": "-s",
	"map_stream":  "-map",
}
var audio_flags = map[string]string{
	"codec":       "-c:a",
	"bit_rate":    "-b:a",
	"channels":    "-ac",
	"sample_rate": "-ar",
	"map_stream":  "-map",
}

func settingsToFlags(video map[string]string, audio map[string]string) []string {
	args := make([]string, len(video)*2+len(audio)*2)
	index := 0
	for k, v := range video {
		args[index] = video_flags[k]
		args[index+1] = v
		index += 2
	}
	for k, v := range audio {
		args[index] = audio_flags[k]
		args[index+1] = v
		index += 2
	}
	return args
}

func Convert(source string, destination string, video map[string]string, audio map[string]string) string {
	fmt.Println("Video:")
	fmt.Println(video)
	fmt.Println("Audio:")
	fmt.Println(audio)
	flags := settingsToFlags(video, audio)
	args := make([]string, 0)
	args = append(args, "-i", source, "-threads", "0", "-n")
	for _, v := range flags {
		args = append(args, v)
	}
	args = append(args, destination)

	fmt.Println(args)

	result, error := exec.Command("ffmpeg", args...).Output()
	if error != nil {
		panic(error)
	}
	return string(result)
}
