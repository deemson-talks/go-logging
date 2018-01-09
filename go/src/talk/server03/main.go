package main

import (
	"net/http"
	"github.com/go-chi/chi"
	"talk/fruits"
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	r := chi.NewRouter()
	r.Get("/eat-fruit", func(w http.ResponseWriter, r *http.Request) {
		rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

		log.Print(`Choosing fruit to eat`)
		fruit := fruits.ChooseRandom()
		time.Sleep(time.Millisecond * time.Duration(10+rnd.Intn(10)))

		log.Printf(`Chose "%s"`, fruit)
		result, err := fruits.Eat(fruit)
		time.Sleep(time.Millisecond * time.Duration(10+rnd.Intn(10)))

		if err != nil {
			log.Printf(`Error eating "%s": %v`, fruit, err)
			w.Write([]byte(fmt.Sprintf(`failed eating "%s": "%s"`, fruit, err)))
			return
		}

		log.Printf(`Yum "%s", %s`, fruit, result)
		w.Write([]byte(fmt.Sprintf(`ate "%s": "%s"`, fruit, result)))
	})
	http.ListenAndServe(":3333", r)
}
