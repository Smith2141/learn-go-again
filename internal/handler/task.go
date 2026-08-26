package handler

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"my-go-app/internal/model"
)

// TaskHandler управляет логикой работы с задачами
type TaskHandler struct {
	mu     sync.Mutex // Защищает карту от одновременной записи из разных горутин
	tasks  map[int]model.Task
	nextID int
}

// NewTaskHandler — конструктор для нашего обработчика
func NewTaskHandler() *TaskHandler {
	return &TaskHandler{
		tasks:  make(map[int]model.Task),
		nextID: 1,
	}
}

// CreateTask обрабатывает POST-запрос на создание задачи
func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	var t model.Task
	// Декодируем JSON из тела запроса в структуру Task
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, "Неверный формат JSON", http.StatusBadRequest)
		return
	}

	// Заполняем системные поля под защитой мьютекса
	h.mu.Lock()
	t.ID = h.nextID
	t.CreatedAt = time.Now()
	if t.Status == "" {
		t.Status = "todo"
	}
	h.tasks[t.ID] = t
	h.nextID++
	h.mu.Unlock()

	// Возвращаем созданную задачу клиенту
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(t)
}

// GetTasks обрабатывает GET-запрос для получения всех задач
func (h *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	h.mu.Lock()
	// Превращаем карту в слайс (массив), чтобы отдать в JSON
	taskList := make([]model.Task, 0, len(h.tasks))
	for _, task := range h.tasks {
		taskList = append(taskList, task)
	}
	h.mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(taskList)
}
