package main

import (
	"crypto/rand"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"github.com/oxtoacart/bpool"
	"github.com/tucnak/telebot"
)

var (
	// buildVersion will automatically fill using Makefile
	buildVersion = "-"

	port = "9113"

	templates *template.Template

	store *securecookie.SecureCookie
)

func init() {
	log.Println("buildVersion", buildVersion)
	runtime.GOMAXPROCS(runtime.NumCPU())

	store = securecookie.New(generateToken(32), generateToken(16))
}

func main() {
	initTemplate()

	initTelegram()
	go listenToMessages()

	go autoUpdate()

	router := mux.NewRouter().StrictSlash(true)
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("template/static"))))
	router.Methods("GET").Path("/").Name("IndexGetHandler").Handler(http.HandlerFunc(indexGetHandler))
	router.Methods("POST").Path("/api/submit").Name("IndexPostHandler").Handler(http.HandlerFunc(indexPostHandler))

	srv := &http.Server{
		Addr:         "127.0.0.1:" + port,
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
	log.Println("listening to", port)
	log.Fatal(srv.ListenAndServe())
}

func indexGetHandler(w http.ResponseWriter, r *http.Request) {
	var lastTime int64 = 0
	if cookie, err := r.Cookie("rcjrescue"); err == nil {
		value := make(map[string]int64)
		if err = store.Decode("rcjrescue", cookie.Value, &value); err == nil {
			lastTime = value["time"]
		}
	}

	var Content struct {
		Enable bool
	}

	Content.Enable = lastTime == 0

	w.Write(render("index", Content))
}

func indexPostHandler(w http.ResponseWriter, r *http.Request) {
	var lastTime int64 = 0
	if cookie, err := r.Cookie("rcjrescue"); err == nil {
		value := make(map[string]int64)
		if err = store.Decode("rcjrescue", cookie.Value, &value); err == nil {
			lastTime = value["time"]
		}
	}
	if lastTime == 0 {
		value := map[string]int64{
			"time": time.Now().Unix(),
		}
		if encoded, err := store.Encode("rcjrescue", value); err == nil {
			cookie := &http.Cookie{
				Name:     "rcjrescue",
				Value:    encoded,
				Path:     "/",
				HttpOnly: true,
				MaxAge:   3600 * 30 * 4,
			}
			http.SetCookie(w, cookie)
		}
	}

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	input := Input{}
	err = json.Unmarshal(bytes, &input)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if lastTime == 0 {
		s := Survey{}
		s.Input = input
		s.Time = time.Now().Unix()
		// because this server will be use proxy with caddy
		s.IP = r.Header.Get("X-Forwarded-For")
		s.Save()

		bot.SendMessage(telebot.User{ID: adminID}, string(bytes), nil)
	}
}

func initTemplate() {
	go func() {
		for {
			templates = template.New("")
			err := filepath.Walk("./template", func(path string, info os.FileInfo, err error) error {
				if strings.Contains(path, ".html") {
					_, err = templates.ParseFiles(path)
				}
				return err
			})
			if err != nil {
				log.Println(err)
				return
			}
			time.Sleep(1 * time.Second)
		}
	}()
}

func render(name string, data interface{}) []byte {
	var bufpool *bpool.BufferPool
	bufpool = bpool.NewBufferPool(128)

	m := make(map[string]interface{})

	m["buildVersion"] = buildVersion
	m["data"] = data

	buf := bufpool.Get()
	defer bufpool.Put(buf)
	err := templates.ExecuteTemplate(buf, name, m)
	if err != nil {
		log.Println(err)
		return nil
	}
	return buf.Bytes()
}

func generateToken(length int) []byte {
	token := make([]byte, length)
	_, err := rand.Read(token)
	if err != nil {
		log.Fatal(err)
	}
	return token
}
