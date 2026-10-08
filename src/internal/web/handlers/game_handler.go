package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"

	"project03/internal/domain/models"
	"project03/internal/domain/service"
	"project03/internal/web/mappers"
	webmodels "project03/internal/web/models"
)

type GameHandler struct {
	service service.GameService
}

func NewGameHandler(service service.GameService) *GameHandler {
	return &GameHandler{service: service}
}

func (h *GameHandler) CreateGame(w http.ResponseWriter, r *http.Request) {

	playerID, err := uuid.Parse(r.Header.Get("X-User-ID"))
	if err != nil {
		sendError(w, "invalid or missing user id", http.StatusUnauthorized)
		return
	}

	var req struct {
		Mode string `json:"mode"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendError(w, "invalid request body", http.StatusBadRequest)
		return
	}

	var mode models.GameMode

	switch req.Mode {
	case "player":
		mode = models.GameWithPlayer
	case "computer":
		mode = models.GameWithComputer
	default:
		sendError(w, "invalid game mode", http.StatusBadRequest)
		return
	}

	game, err := h.service.CreateGame(
		r.Context(),
		mode,
		playerID,
	)
	if err != nil {
		sendError(w, "failed to create game", http.StatusInternalServerError)
		return
	}

	response := mappers.ToNewGameResponse(game)
	sendJSON(w, response, http.StatusCreated)
}
func (h *GameHandler) MakeMove(w http.ResponseWriter, r *http.Request) {
	gameID := r.PathValue("id")

	var req webmodels.GameRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendError(w, "invalid request body", http.StatusBadRequest)
		return
	}

	// Временно используем UUID из заголовка.
	// Позже его будет устанавливать UserAuthenticator
	// после успешной авторизации.
	playerID, err := uuid.Parse(r.Header.Get("X-User-ID"))
	if err != nil {
		sendError(w, "invalid or missing user id", http.StatusUnauthorized)
		return
	}

	err = h.service.MakeMove(
		r.Context(),
		gameID,
		playerID,
		req.Row,
		req.Col,
	)
	if err != nil {
		sendError(w, err.Error(), http.StatusBadRequest)
		return
	}

	game, err := h.service.GetGame(r.Context(), gameID)
	if err != nil {
		sendError(w, "failed to get game", http.StatusInternalServerError)
		return
	}

	response := mappers.ToWebResponse(game, "")
	sendJSON(w, response, http.StatusOK)
}

func (h *GameHandler) GetGame(w http.ResponseWriter, r *http.Request) {
	gameID := r.PathValue("id")

	game, err := h.service.GetGame(r.Context(), gameID)
	if err != nil {
		sendError(w, "game not found", http.StatusNotFound)
		return
	}

	response := mappers.ToWebResponse(game, "")
	sendJSON(w, response, http.StatusOK)
}

func sendJSON(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		return
	}
}

func sendError(w http.ResponseWriter, message string, status int) {
	sendJSON(w, map[string]string{
		"error": message,
	}, status)
}
