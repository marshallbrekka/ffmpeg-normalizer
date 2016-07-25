package ffmpeg

import "os/exec"
import "encoding/json"
import "fmt"
import "strconv"

type Stream struct {
	index       int
	codec       string
	codec_type  string
	width       int
	height      int
	bit_rate    int
	sample_rate int
	language    string
	channels    int
}

type Streams []Stream

func (slice Streams) Len() int {
	return len(slice)
}

func (slice Streams) Less(i, j int) bool {
	return slice[i].channels < slice[j].channels
}

func (slice Streams) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func keyInt(jsonStream map[string]interface{}, key string) int {
	value := jsonStream[key]
	if value != nil {
		return int(value.(float64))
	} else {
		return 0
	}
}

func keyString(jsonStream map[string]interface{}, key string) string {
	value := jsonStream[key]
	if value != nil {
		return value.(string)
	}
	return ""
}

func keyStringToInt(jsonStream map[string]interface{}, key string) int {
	value := jsonStream[key]
	if value != nil {
		result, err := strconv.ParseUint(value.(string), 10, 32)
		if err != nil {
			panic(err)
		}
		return int(result)
	}
	return 0
}

func jsonStreamToStruct(jsonStream map[string]interface{}) Stream {
	fmt.Println(jsonStream)
	var tags map[string]interface{}
	if jsonStream["tags"] == nil {
		tags = map[string]interface{} {}

	} else {
		tags = jsonStream["tags"].(map[string]interface{})
	}

	return Stream{
		index:       keyInt(jsonStream, "index"),
		codec:       keyString(jsonStream, "codec_name"),
		codec_type:  keyString(jsonStream, "codec_type"),
		width:       keyInt(jsonStream, "width"),
		height:      keyInt(jsonStream, "height"),
		bit_rate:    keyStringToInt(jsonStream, "bit_rate"),
		sample_rate: keyStringToInt(jsonStream, "sample_rate"),
		channels:    keyInt(jsonStream, "channels"),
		language:    keyString(tags, "language")}
}

func populateVideoBitRate(probe map[string]interface{}, streams Streams) Streams {
	format := probe["format"].(map[string]interface{})
	total := keyStringToInt(format, "bit_rate")
	for _, v := range streams {
		if v.codec_type == "audio" {
			total -= v.bit_rate
		}
	}
	streams[0].bit_rate = total
	return streams
}

func Probe(filePath string) map[string]interface{} {
	result, error := exec.Command("ffprobe",
		"-i",
		filePath,
		"-print_format",
		"json=compact=1",
		"-show_streams",
		"-show_format").Output()
	if error != nil {
		panic(error)
	} else {
		var data map[string]interface{}
		if err := json.Unmarshal(result, &data); err != nil {
			fmt.Println(err)
			panic(err)
		}
		return data
	}
}

func ProbeToStreams(probe map[string]interface{}) Streams {
	streamsJson := probe["streams"].([]interface{})

	streams := make(Streams, len(streamsJson))
	for i, v := range streamsJson {
		streams[i] = jsonStreamToStruct(v.(map[string]interface{}))
	}

	return populateVideoBitRate(probe, streams)
}
