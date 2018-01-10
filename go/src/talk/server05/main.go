package main

import (
	"net/http"
	"github.com/go-chi/chi"
	"talk/fruits"
	"fmt"
	"math/rand"
	"time"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog"
	"os"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Get("/eat-fruit", func(w http.ResponseWriter, r *http.Request) {
		rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

		reqID := middleware.GetReqID(r.Context())
		log := log.With().Str("reqID", reqID).Logger()
		log.Info().Msg("Choosing")
		fruit := fruits.ChooseRandom()
		time.Sleep(time.Millisecond * time.Duration(10+rnd.Intn(10)))

		log = log.With().Str("fruit", fruit).Logger()
		log.Info().Msg("Chosen")
		result, err := fruits.Eat(fruit)
		time.Sleep(time.Millisecond * time.Duration(10+rnd.Intn(10)))

		if err != nil {
			log.Error().Err(err).Msg("Failed")
			w.Write([]byte(fmt.Sprintf(`failed eating "%s": "%s"`, fruit, err)))
			return
		}

		log.Info().Str("result", result).Msg("Yum")
		w.Write([]byte(fmt.Sprintf(`ate "%s": "%s"`, fruit, result)))
	})
	err := http.ListenAndServe(":3333", r)
	if err != nil {
		panic(err)
	}
}
