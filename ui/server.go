package ui

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/dpordomingo/learning-exercises/ant/actors"
	"github.com/dpordomingo/learning-exercises/ant/generators"
	"github.com/dpordomingo/learning-exercises/ant/geo"
)

//RunServer starts a server that will report the world state
func RunServer(m *geo.Map, target *geo.Point, rover actors.Rover) {
	http.HandleFunc("/", getIntroHandler(m))
	http.HandleFunc("/favicon.ico", dummyHandler)
	http.HandleFunc("/events", getEventSourceHandler(m))
	http.ListenAndServe(":8888", nil)
}

func dummyHandler(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusNoContent)
}

func getIntroHandler(m *geo.Map) http.HandlerFunc {
	templateScope := map[string]int32{
		"Width":  10 * m.W(),
		"Height": 10 * m.H(),
	}
	return func(res http.ResponseWriter, req *http.Request) {
		header := res.Header()
		introTpl, err := readTemplate("ui/html.tpl")
		if err != nil {
			http.Error(res, "Error loading template", http.StatusInternalServerError)
			return
		}
		if err := introTpl.Execute(res, templateScope); err != nil {
			http.Error(res, "Error rendering template", http.StatusInternalServerError)
			return
		}

		header.Set("Content-Type", "text/html;charset=utf-8")
	}
}

const CONTENT string = "" +
	"id: %d\n" +
	"retry: 500\n" +
	"data: %d %d\n\n"

func getEventSourceHandler(m *geo.Map) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		fmt.Println("Conection opened\n")
		defer fmt.Println("Conection closed\n")

		header := res.Header()
		header.Set("Content-Type", "text/event-stream")
		header.Set("Cache-Control", "no-cache")
		header.Set("Connection", "keep-alive")

		flusher, ok := res.(http.Flusher)
		if !ok {
			http.Error(res, "Streaming unsupported!", http.StatusInternalServerError)
			return
		}

		closingChannel := res.(http.CloseNotifier).CloseNotify()
		notifyChannel := make(chan bool)

		go func() {
			notifyChannel <- <-closingChannel
		}()

		go func() {
			for {
				notifyChannel <- false
				time.Sleep(time.Millisecond * 500)
			}
		}()

		for !<-notifyChannel {
			point := generators.GetRandomPoint(10*m.W(), 10*m.H())
			content := fmt.Sprintf(CONTENT, time.Now().UTC().UnixNano(), point.X, point.Y)
			res.Write([]byte(content))
			flusher.Flush()
			fmt.Println(content)
		}
	}
}

func readTemplate(file string) (*template.Template, error) {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	return template.New("intro").Parse(string(bytes))
}
