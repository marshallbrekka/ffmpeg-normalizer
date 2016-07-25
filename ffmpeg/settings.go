package ffmpeg

import "sort"
import "strconv"

var videoPassover = map[string]bool{
	"MKV":  true,
	"h264": true,
	"mp4":  true,
}

var  audioPassover = map[string]bool{
	"aac": true,
}

const videoCodec = "libx264"
const audioCodec = "aac"
const copyCodec = "copy"

func Video(stream Stream) map[string]string {
	if videoPassover[stream.codec] {
		return map[string]string{
			"codec":      copyCodec,
			"map_stream": "0:" + strconv.Itoa(stream.index),
		}
	} else {
		return map[string]string{
			"codec":       videoCodec,
			"dimmensions": strconv.Itoa(stream.width) + "x" + strconv.Itoa(stream.height),
			"map_stream":  "0:" + strconv.Itoa(stream.index),
			"bit_rate":    strconv.Itoa(stream.bit_rate),
		}
	}
}

func Audio(stream Stream) map[string]string {
	if audioPassover[stream.codec] {
		return map[string]string{
			"codec":      copyCodec,
			"map_stream": "0:" + strconv.Itoa(stream.index),
		}
	} else {
		return map[string]string{
			"codec":       audioCodec,
			"map_stream":  "0:" + strconv.Itoa(stream.index),
			"channels":    strconv.Itoa(stream.channels),
			"bit_rate":    strconv.Itoa(stream.bit_rate),
			"sample_rate": strconv.Itoa(stream.sample_rate),
		}
	}
}

func BestAudioStream(streams Streams) Stream {
	audioStreams := make(Streams, 0)
	for _, stream := range streams {
		if stream.codec_type == "audio" {
			if stream.language == "" ||
				stream.language == "und" ||
				stream.language == "eng" {
				audioStreams = append(audioStreams, stream)
			}
		}
	}
	sort.Sort(audioStreams)
	return audioStreams[0]
}
