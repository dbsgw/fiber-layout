package service

import (
	v1 "fiber-layout/validator/form"
	"fmt"
)

type Default struct {
}

func NewDefaultService() *Default {
	return &Default{}
}

func (t *Default) Register(loginForm v1.Register) (string, error) {
	msg := fmt.Sprintf("✋ ---- %s", loginForm.UserName)
	return msg, nil
}

func (t *Default) Login(loginForm v1.Login) (string, error) {
	msg := fmt.Sprintf("✋ ---- %s", loginForm.PassWord)
	return msg, nil
}
