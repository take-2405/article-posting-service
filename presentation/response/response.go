package response

import (
	"encoding/json"
	"log"
	"net/http"
)

// HTTPError エラー用
type HTTPError struct {
	Code    int
	Message string
}

// respondError レスポンスとして返すエラーを生成する
func RespondError(w http.ResponseWriter, code int, err error) {
	log.Printf("err: %v", err)
	if err != nil {
		he := HTTPError{
			Code:    code,
			Message: err.Error(),
		}
		RespondJSON(w, code, he)
	}
	//if e, ok := err.(*HTTPError); ok {
	//	respondJSON(w, e.Code, e)
	//} else if err != nil {
	//	he := HTTPError{
	//		Code:    code,
	//		Message: err.Error(),
	//	}
	//	respondJSON(w, code, he)
	//}
}

// respondJSON レスポンスとして返すjsonを生成して、writerに書き込む
func RespondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.MarshalIndent(payload, "", "    ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}
