package main

import (
	"encoding/json"
	"flag"
	"goRoutines/hollymicro/types"
	"log"
	"net/http"
	"time"

	"github.com/anthdm/hollywood/actor"
	"github.com/anthdm/hollywood/remote"
)

const (
	scrapeInterval = time.Second
	url            = "https://catfact.ninja/fact"
)

type scraper struct {
	url      string
	storePID *actor.PID
	engine   *actor.Engine
}

func newScraper(url string, storePID *actor.PID) actor.Producer {
	return func() actor.Receiver {
		return &scraper{
			url:      url,
			storePID: storePID,
		}
	}
}

func (s *scraper) Receive(c *actor.Context) {
	switch msg := c.Message().(type) {
	case actor.Started:
		s.engine = c.Engine()
		go s.scrapeLoop()
	case actor.Stopped:
	default:
		_ = msg
	}
}

func (s *scraper) scrapeLoop() {
	for {
		resp, err := http.Get(s.url)
		if err != nil {
			panic(err)
		}
		var res Catfact
		if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
			log.Println("err")
			continue
		}

		s.engine.Send(s.storePID, &types.CatFact{
			Fact: res.Fact,
		})

		time.Sleep(scrapeInterval)
	}
}

type Catfact struct {
	Fact string `json:"fact"`
}

func main() {
	listenAddr := flag.String("listenAddr", "127.0.0.1:3000", "todo")
	flag.Parse()

	e := actor.NewEngine()
	r := remote.New(e, remote.Config{ListenAddr: *listenAddr})
	e.WithRemote(r)

	// pid 127.0.0.1/store
	storePID := actor.NewPID("127.0.0.1:4000", "store")

	e.Spawn(newScraper(url, storePID), "scrapper")

	select {}
}
