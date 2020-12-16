package service

import "go_playground/Go-000/Week04/myApp/internal/app/myApp/biz"

type Service struct {
	biz.Biz
}

func NewService(biz biz.Biz) Service {
	return Service{biz}
}
