package todoapi

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	gorp "gopkg.in/gorp.v1"
)

// 핸들러에서 DB에 쿼리를 보낼 수 있도록 DB 접속 정보 구조체를 갖는다
type TodoHandler struct {
	master *gorp.DbMap
	slave  *gorp.DbMap
}

// 핸들러 생성 함수
func NewTodoHandler(master *gorp.DbMap, slave *gorp.DbMap) http.Handler {
	return &TodoHandler{
		master: master,
		slave:  slave,
	}
}

// HTTP 요청을 받아 비즈니스 로직을 수행하고 응답을 돌려준다
func (h TodoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	log.Printf("[%s] RemoteAddr=%s\tUserAgent=%s", r.Method, r.RemoteAddr, r.Header.Get("User-Agent"))
	switch r.Method {
	case "GET":
		h.serveGET(w, r)
		return
	case "POST":
		h.servePOST(w, r)
		return
	case "PUT":
		h.servePUT(w, r)
		return
	default:
		NewErrorResponse(http.StatusMethodNotAllowed, fmt.Sprintf("%s is Unsupported method", r.Method)).Write(w)
		return
	}
}

type errorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func NewErrorResponse(status int, message string) *errorResponse {
	return &errorResponse{
		Status:  status,
		Message: message,
	}
}

func (e *errorResponse) Write(w http.ResponseWriter) {
	data, err := json.Marshal(e)
	if err != nil {
		log.Println("marshal error json is failed")
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(e.Status)
	w.Write(data)
}

func (h TodoHandler) serveGET(w http.ResponseWriter, r *http.Request) {

	status := r.URL.Query().Get("status")
	if status == "" {
		status = "TODO"
	}

	result, err := h.slave.Select(Todo{}, "SELECT * FROM todo WHERE status = ? ORDER BY updated DESC", status)
	if err != nil {
		log.Println(err.Error())
		NewErrorResponse(http.StatusInternalServerError, "Execute Query is failed").Write(w)
		return
	}

	todoItems := make([]*Todo, 0)

	for _, e := range result {
		todo := e.(*Todo)
		todoItems = append(todoItems, todo)
	}

	data, err := json.Marshal(todoItems)
	if err != nil {
		log.Println(err.Error())
		NewErrorResponse(http.StatusInternalServerError, "Marshal JSON is failed").Write(w)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(data)
}

type TodoPostPayload struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (h TodoHandler) servePOST(w http.ResponseWriter, r *http.Request) {

	raw, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err.Error())
		NewErrorResponse(http.StatusInternalServerError, "Read payload is failed").Write(w)
		return
	}

	var payload TodoPostPayload
	if err := json.Unmarshal(raw, &payload); err != nil {
		log.Println(err.Error())
		NewErrorResponse(http.StatusInternalServerError, "Parse payload is failed").Write(w)
		return
	}

	now := time.Now()
	todo := Todo{
		Title:   payload.Title,
		Content: payload.Content,
		Status:  "TODO",
		Created: now,
		Updated: now,
	}

	if err := h.master.Insert(&todo); err != nil {
		log.Println(err.Error())
		NewErrorResponse(http.StatusInternalServerError, "Insert Data is failed").Write(w)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusCreated)
}

type TodoPutPayload struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Status  string `json:"status"`
}

func (h TodoHandler) servePUT(w http.ResponseWriter, r *http.Request) {

	raw, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err.Error())
		NewErrorResponse(http.StatusInternalServerError, "Read payload is failed").Write(w)
		return
	}

	var payload TodoPutPayload
	if err := json.Unmarshal(raw, &payload); err != nil {
		log.Println(err.Error())
		NewErrorResponse(http.StatusInternalServerError, "Parse payload is failed").Write(w)
		return
	}

	var target Todo
	if err := h.slave.SelectOne(&target, "SELECT * FROM todo WHERE id = ?", payload.ID); err != nil {
		if err == sql.ErrNoRows {
			NewErrorResponse(http.StatusNotFound, fmt.Sprintf("id=%d is not found", payload.ID)).Write(w)
		} else {
			log.Println(err.Error())
			NewErrorResponse(http.StatusInternalServerError, "Select todo is failed").Write(w)
		}
		return
	}

	now := time.Now()
	todo := Todo{
		ID:      payload.ID,
		Title:   payload.Title,
		Content: payload.Content,
		Status:  payload.Status,
		Created: target.Created,
		Updated: now,
	}

	if _, err := h.master.Update(&todo); err != nil {
		log.Println(err.Error())
		NewErrorResponse(http.StatusInternalServerError, "Update Data is failed").Write(w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
