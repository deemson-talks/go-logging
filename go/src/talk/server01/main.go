package main

import (
	"net/http"
	"github.com/go-chi/chi"
	"talk/fruits"
	"fmt"
	"time"
	"math/rand"
)

func main() {
	r := chi.NewRouter()
	r.Get("/eat-fruit", func(w http.ResponseWriter, r *http.Request) {
		rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

		fruit := fruits.ChooseRandom()
		time.Sleep(time.Millisecond * time.Duration(10+rnd.Intn(10)))

		result, err := fruits.Eat(fruit)
		time.Sleep(time.Millisecond * time.Duration(10+rnd.Intn(10)))

		if err != nil {
			w.Write([]byte(fmt.Sprintf(`failed eating "%s": "%s"`, fruit, err)))
			return
		}
		w.Write([]byte(fmt.Sprintf(`ate "%s": "%s"`, fruit, result)))
	})
	http.ListenAndServe(":3333", r)
}
