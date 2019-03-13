package todoapi

import (
	"errors"
	"os"
)

// 환경 변수를 저장할 구조체
type Env struct {
	Bind      string
	MasterURL string
	SlaveURL  string
}

func CreateEnv() (*Env, error) {

	env := Env{}

	bind := os.Getenv("TODO_BIND") // API 서버 포트 설정
	if bind == "" {
		env.Bind = ":8080"
	}
	env.Bind = bind

	masterURL := os.Getenv("TODO_MASTER_URL") // MySQL 마스터 접속정보
	if masterURL == "" {
		return nil, errors.New("TODO_MASTER_URL is not specified")
	}
	env.MasterURL = masterURL

	slaveURL := os.Getenv("TODO_SLAVE_URL") // MySQL 슬레이브 접속정보
	if slaveURL == "" {
		return nil, errors.New("TODO_SLAVE_URL is not specified")
	}
	env.SlaveURL = slaveURL

	return &env, nil
}
