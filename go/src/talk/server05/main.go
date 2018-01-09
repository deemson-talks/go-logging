package main

import (
	"net/http"
	"github.com/go-chi/chi"
	"talk/fruits"
	"fmt"
	"math/rand"
	"time"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/zerolog"
	"os"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Get("/eat-fruit", func(w http.ResponseWriter, r *http.Request) {
		rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

		reqID := middleware.GetReqID(r.Context())
		log := log.With().Str("reqID", reqID).Logger()
		log.Print("Choosing fruit to eat")
		fruit := fruits.ChooseRandom()
		time.Sleep(time.Millisecond * time.Duration(10+rnd.Intn(10)))

		log = log.With().Str("fruit", fruit).Logger()
		log.Print("chosen")
		result, err := fruits.Eat(fruit)
		time.Sleep(time.Millisecond * time.Duration(10+rnd.Intn(10)))

		if err != nil {
			log.Printf(`Error eating: %v`, err)
			w.Write([]byte(fmt.Sprintf(`failed eating "%s": "%s"`, fruit, err)))
			return
		}

		log.Printf(`Yum %s`, result)
		w.Write([]byte(fmt.Sprintf(`ate "%s": "%s"`, fruit, result)))
	})
	err := http.ListenAndServe(":3333", r)
	if err != nil {
		panic(err)
	}
}
