package tts

import (
	"fmt"
	"github.com/mozillazg/go-pinyin"
	"io"
	"log"
	"os"
	"os/exec"
	"path"
)

const VOICE = "voice"

func Generate(text string, writer io.Writer) error {
	args := pinyin.NewArgs()
	args.Style = pinyin.Tone3
	args.Heteronym = true
	pins := pinyin.LazyPinyin(text, args)

	log.Println(pins)

	//files := []string{"ffmpeg"}
	var files []string
	for _, p := range pins {
		fn := path.Join(VOICE, p+".wav")
		files = append(files, "-i", fn)
	}
	files = append(files, "-filter_complex")

	filter := ""
	for i, _ := range pins {
		filter += fmt.Sprintf("[%d:0] ", i)
	}
	filter += fmt.Sprintf("concat=n=%d:v=0:a=1 [out]", len(pins))
	files = append(files, filter)

	//files = append(files, " -map '[out]' test.wav")
	files = append(files, "-map", "[out]")
	files = append(files, "-f", "mp3") //输出MP3
	files = append(files, "pipe:1")    //标准输出
	files = append(files, "-y")

	fmt.Println(files)

	//c := strings.Join(files, " ")
	cmd := exec.Command("./ffmpeg", files...)
	cmd.Stdout = writer
	//cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
