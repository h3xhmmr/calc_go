package application

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"yandex_project/pkg/calculation"
)

type Application struct {
}

func New() *Application {
	return &Application{}
}

type Request struct {
	Expression string `json:"expression"`
}

type Response struct {
	Result string `json:"result"`
}

func CalcHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, `{"error":"Wrong Method"}`, http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, `{"error":"Invalid Body"}`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var request Request
	err = json.Unmarshal(body, &request)
	if err != nil || request.Expression == "" {
		http.Error(w, `{"error":"Invalid Body"}`, http.StatusBadRequest)
		return
	}

	result, err := calculation.Calc(request.Expression)
	if err != nil {
		if err == calculation.ErrInvalidExpression {
			http.Error(w, fmt.Sprintf(`{"error":"%s"}`, calculation.ErrInvalidExpression), http.StatusUnprocessableEntity)
		} else if err == calculation.ErrDivisionByZero {
			http.Error(w, fmt.Sprintf(`{"error":"%s"}`, calculation.ErrDivisionByZero), http.StatusUnprocessableEntity)
		}
		return
	}
	var resp Response
	resp.Result = fmt.Sprintf("%f", result)
	json_resp, err := json.Marshal(resp)
	if err != nil {
		log.Printf("Error while marshaling: %v", err)
		http.Error(w, `{"error":"Unknown error occurred"}`, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(json_resp)
	if err != nil {
		log.Printf("Error writing response: %v", err)
	}
}

func (a *Application) Run() {
	http.HandleFunc("/api/v1/calculate", CalcHandler)
	http.ListenAndServe(":8080", nil)
}

//весь проект мегасырой, тк не хватило времени нормально его доделать, так что, если поставите 0 баллов - не обижусь)
