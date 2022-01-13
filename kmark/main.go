package main

import (
	"bufio"
	_ "embed"
	"os"
	"sync"
	"text/template"

	"bitbucket.org/ai69/amoy"
	"github.com/1set/gut/yhash"
	"github.com/1set/gut/yos"
	"github.com/1set/gut/ystring"
	"github.com/fsnotify/fsnotify"
	gfm "github.com/shurcooL/github_flavored_markdown"
	flag "github.com/spf13/pflag"
	"go.uber.org/zap"
)

var (
	//go:embed resource/index.html
	templateContent string
	templateName    = "basic"
)

func main() {
	args := flag.Args()
	if len(args) <= 0 {
		log.Fatal("no arguments")
	}

	// check input file
	inputFile := args[0]
	if !yos.ExistFile(inputFile) {
		log.Fatalw("input file doesn't not exist", "filename", inputFile)
	}

	// output file
	if ystring.IsBlank(outputFile) {
		outputFile = ystring.NotBlankOrDefault(amoy.SubstrBeforeLast(inputFile, "."), "render") + ".html"
	}
	log.Infow("will render markdown as html", "input_file", inputFile, "output_file", outputFile, "live_render", liveRender)

	// rendering
	if liveRender {
		// set rendering func and watcher
		var renderLock sync.Mutex
		render := func() {
			renderLock.Lock()
			defer renderLock.Unlock()

			if err := renderMarkdown(inputFile, outputFile); err != nil {
				log.Warnw("fail to render markdown", zap.Error(err))
			}
		}
		fileWatcher(inputFile, render)

		// block until exit
		render()
		c := make(chan struct{})
		<-c
	} else {
		if err := renderMarkdown(inputFile, outputFile); err != nil {
			log.Fatalw("fail to render markdown", zap.Error(err))
		}
	}
}

func renderMarkdown(inputFile, outputFile string) error {
	// read content
	content, err := amoy.ReadFileString(inputFile)
	if err != nil {
		log.Errorw("fail to read content file", "filename", inputFile, zap.Error(err))
		return err
	}

	// render markdown
	renderContent := string(gfm.Markdown([]byte(content)))

	// load html template
	tmpl, err := template.New(templateName).Parse(templateContent)
	if err != nil {
		log.Errorw("fail to parse template", zap.Error(err))
		return err
	}
	data := struct {
		Title   string
		Content string
	}{
		Title:   inputFile,
		Content: renderContent,
	}

	// render as html file
	fout, err := os.Create(outputFile)
	if err != nil {
		log.Errorw("fail to create file", "filename", outputFile, zap.Error(err))
		return err
	}
	defer fout.Close()

	fw := bufio.NewWriter(fout)
	defer fw.Flush()
	if err = tmpl.Execute(fw, data); err != nil {
		log.Errorw("fail to execute template", zap.Error(err))
		return err
	}

	// done
	log.Infow("rendered markdown as html", "input_file", inputFile, "output_file", outputFile)
	return nil
}

func fileWatcher(fileName string, action func()) error {
	hash, err := yhash.FileMD5(fileName)
	if err != nil {
		return err
	}
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	go func() {
		defer watcher.Close()

		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					if newHash, _ := yhash.FileMD5(fileName); newHash != hash {
						log.Debugw("watched file modified", "event", event.Name, "md5_old", hash, "md5_new", newHash)
						hash = newHash
						action()
					}
				}
			}
		}
	}()
	return watcher.Add(fileName)
}
