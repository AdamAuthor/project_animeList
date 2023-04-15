package http

import (
	"animeList/internal/content"
	"animeList/internal/models"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Server struct {
	ctx        context.Context
	idleConsCh chan struct{}

	// не самая лучшая практика. Обычно делается на уровне 3 слоев:
	// бизнес логика, HTTP хэндлеры, база данных (пока не знаю как это сделать)
	Address string
	content content.Content
}

// NewServer is the function for creating new server
// Здесь мы создаём свой сервер
func NewServer(ctx context.Context, address string, content content.Content) *Server {
	return &Server{
		ctx:        ctx,
		idleConsCh: make(chan struct{}),
		content:    content,
		Address:    address,
	}
}

// basicHandler был создан для инкапсуляции логики настройки мультиплексера
// К тому же, вместо использования мультиплексера, используется роутер
func (s *Server) basicHandler() chi.Router {
	r := chi.NewRouter()

	// Create
	r.Post("/content", func(w http.ResponseWriter, r *http.Request) {
		nContent := new(models.Anime)
		if err := json.NewDecoder(r.Body).Decode(nContent); err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			fmt.Fprintf(w, "Unknown error: %v", err)
			return
		}

		err := s.content.Create(r.Context(), nContent)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "DB err: %v", err)
		}
		w.WriteHeader(http.StatusCreated)
	})

	// Read All
	r.Get("/content", func(w http.ResponseWriter, r *http.Request) {
		queryValues := r.URL.Query()
		filter := &models.ContentFilter{}
		if searchQuery := queryValues.Get("query"); searchQuery != "" {
			filter.Query = &searchQuery
		}

		nContent, err := s.content.All(r.Context(), filter)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "DB err: %v", err)
		}

		render.JSON(w, r, nContent)
	})

	// Read by id
	r.Get("/content/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Unknown error: %v", err)
			return
		}

		nContent, err := s.content.ByID(r.Context(), id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Unknown error: %v", err)
			return
		}

		render.JSON(w, r, nContent)
	})

	// Update
	r.Put("/content", func(w http.ResponseWriter, r *http.Request) {
		nContent := new(models.Anime)
		if err := json.NewDecoder(r.Body).Decode(nContent); err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			fmt.Fprintf(w, "Unknown error: %v", err)
			return
		}

		err := validation.ValidateStruct(
			nContent,
			validation.Field(&nContent.Title, validation.Required),
			validation.Field(&nContent.Genre, validation.Required),
			validation.Field(&nContent.ImageURL, validation.Required),
			validation.Field(&nContent.Author, validation.Required),
			validation.Field(&nContent.ReleaseYear, validation.Required))
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			fmt.Fprintf(w, "Unknown error: %v", err)
			return
		}

		err = s.content.Update(r.Context(), nContent)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Unknown error: %v", err)
			return
		}
	})

	// Delete
	r.Delete("/content/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Unknown error: %v", err)
			return
		}
		_ = s.content.Delete(r.Context(), id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Unknown error: %v", err)
			return
		}
	})

	return r
}

// Run is the function for running server
func (s *Server) Run() error {
	srv := &http.Server{
		Addr:         s.Address,
		Handler:      s.basicHandler(),
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 30,
	}
	go s.ListenCtxForGT(srv)

	log.Println("[HTTP] Server running on", s.Address)
	return srv.ListenAndServe()
}

// ListenCtxForGT is the function for Graceful Shutdown
// При запуске сервера мы также запускаем горутину, которая дожидается своего часа
// и как только контекст будет завершён, происходит Shutdown
func (s *Server) ListenCtxForGT(srv *http.Server) {
	<-s.ctx.Done() // Blocked until the application context is canceled

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Println("[HTTP] Got err while shutting down:", err)
	}

	log.Println("[HTTP] Processed all idle connections")
	close(s.idleConsCh)
	// как только закрывается канал, функция WaitForGT завершается
	// и наш сервер полностью завершает работу
}

// WaitForGT is the function for waiting until ListenCtxForGT it will work
// С помощью канала функция позволяет дождаться исполнения Graceful Shutdown
func (s *Server) WaitForGT() {
	<-s.idleConsCh // блок до записи или закрытия канала
}
