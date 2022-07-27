package linkchecker

import (
	"net/http"
)

func GetStatus(url string) (int, error)  {
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	return resp.StatusCode, nil	
}