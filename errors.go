package bosskg

import "fmt"

type APIError struct {
	FunCode FunCode
	Code    string
	Message string
	Raw     string
}

func (e APIError) Error() string {
	if e.FunCode != "" {
		return fmt.Sprintf("bosskg api error: funCode=%s code=%s msg=%s", e.FunCode, e.Code, e.Message)
	}
	return fmt.Sprintf("bosskg api error: code=%s msg=%s", e.Code, e.Message)
}
