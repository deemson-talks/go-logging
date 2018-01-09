package main

import (
	"net/http"
	"github.com/go-chi/chi"
	"talk/fruits"
	"fmt"
	"log"
	"math/rand"
	"time"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Get("/eat-fruit", func(w http.ResponseWriter, r *http.Request) {
		rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

		reqID := middleware.GetReqID(r.Context())
		log.Printf(`%s: Choosing fruit to eat`, reqID)
		fruit := fruits.ChooseRandom()
		time.Sleep(time.Millisecond * time.Duration(10+rnd.Intn(10)))

		log.Printf(`%s: Chose "%s"`, reqID, fruit)
		result, err := fruits.Eat(fruit)
		time.Sleep(time.Millisecond * time.Duration(10+rnd.Intn(10)))

		if err != nil {
			log.Printf(`%s: Error eating "%s": %v`, reqID, fruit, err)
			w.Write([]byte(fmt.Sprintf(`failed eating "%s": "%s"`, fruit, err)))
			return
		}

		log.Printf(`%s: Yum "%s", %s`, reqID, fruit, result)
		w.Write([]byte(fmt.Sprintf(`ate "%s": "%s"`, fruit, result)))
	})
	err := http.ListenAndServe(":3333", r)
	if err != nil {
		panic(err)
	}
}
