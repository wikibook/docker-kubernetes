package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gihyodocker/todoapi"
)

func main() {

	// (1) 환경 변수를 저장할 구조
	env, err := todoapi.CreateEnv()
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}

    // (2) MySQL 마스터에 접속하는 데 사용하는 구조체
	masterDB, err := todoapi.CreateDbMap(env.MasterURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s is invalid database", env.MasterURL)
		return
	}

	// (3) MySQL 슬레이브에 접속하는 데 사용하는 구조체
	slaveDB, err := todoapi.CreateDbMap(env.SlaveURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s is invalid database", env.SlaveURL)
		return
	}

	mux := http.NewServeMux()

	// (4) 헬스체크에 사용되는 API 핸들러
	hc := func(w http.ResponseWriter, r *http.Request) {
		log.Println("[GET] /hc")
		w.Write([]byte("OK"))
	}

	// (5) TODO API 핸들러
	todoHandler := todoapi.NewTodoHandler(masterDB, slaveDB)

	// (6) 핸들러를 API 엔드포인트로 등록
	mux.Handle("/todo", todoHandler)
	mux.HandleFunc("/hc", hc)

	// (7) 서버 포트 및 핸들러 설정, 서버 시작
	s := http.Server{
		Addr:    env.Bind,
		Handler: mux,
	}
	log.Printf("Listen HTTP Server")
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
