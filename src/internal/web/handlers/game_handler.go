package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"project03/internal/domain/models"
	"project03/internal/domain/service"
	"project03/internal/web/mappers"
	webmodels "project03/internal/web/models"
)

type GameHandler struct {
	service service.GameService
}

func NewGameHandler(service service.GameService) *GameHandler {
	return &GameHandler{
		service: service,
	}
}

// newBoard := mappers.ToDomainBoard(req.Board)

func (h *GameHandler) CreateGame(w http.ResponseWriter, r *http.Request) {
	game := h.service.NewGame()
	response := mappers.ToNewGameResponse(game)
	sendJSON(w, response, http.StatusCreated)
}

func (h *GameHandler) MakeMove(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 3 {
		sendError(w, "invalid URL", http.StatusBadRequest)
		return
	}
	gameID := pathParts[2]

	var req webmodels.GameRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendError(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if req.ID != gameID {
		sendError(w, "game ID mismatch", http.StatusBadRequest)
		return
	}

	currentGame, err := h.service.GetGame(gameID)
	if err != nil {
		sendError(w, "game not found", http.StatusNotFound)
		return
	}

	if finished, _, _ := currentGame.IsGameOver(); finished {
		sendError(w, "game is already over", http.StatusConflict)
		return
	}

	newBoard := mappers.ToDomainBoard(req.Board)

	if err := h.service.ValidateBoard(currentGame, newBoard); err != nil {
		sendError(w, err.Error(), http.StatusBadRequest)
		return
	}

	gameAfterPlayerMove := models.Game{
		ID:    currentGame.ID,
		Board: newBoard,
		Turn:  2,
	}

	if finished, _, result := gameAfterPlayerMove.IsGameOver(); finished {
		h.service.SaveGame(gameAfterPlayerMove)
		response := mappers.ToWebResponse(gameAfterPlayerMove, result)
		sendJSON(w, response, http.StatusOK)
		return
	}

	computerBoard, err := h.service.GetNextMove(gameAfterPlayerMove)
	if err != nil {
		sendError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	gameAfterComputer := models.Game{
		ID:    currentGame.ID,
		Board: computerBoard,
		Turn:  1,
	}

	_, _, result := gameAfterComputer.IsGameOver()

	if err := h.service.SaveGame(gameAfterComputer); err != nil {
		sendError(w, "failed to save game", http.StatusInternalServerError)
		return
	}

	response := mappers.ToWebResponse(gameAfterComputer, result)
	sendJSON(w, response, http.StatusOK)
}

func (h *GameHandler) GetGame(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 3 {
		sendError(w, "invalid URL", http.StatusBadRequest)
		return
	}
	gameID := pathParts[2]

	game, err := h.service.GetGame(gameID)
	if err != nil {
		sendError(w, "game not found", http.StatusNotFound)
		return
	}

	finished, _, result := game.IsGameOver()
	if finished {
		response := mappers.ToWebResponse(game, result)
		sendJSON(w, response, http.StatusOK)
		return
	}

	response := mappers.ToWebResponse(game, "")
	sendJSON(w, response, http.StatusOK)
}

func sendJSON(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func sendError(w http.ResponseWriter, message string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}
