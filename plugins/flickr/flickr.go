package flickr

import (
	"fmt"
	"net/http"

	"github.com/lateralusd/sochecker/checker"
)

type FLCheck struct {
	link string
	err  error
}

func (fl *FLCheck) Check(client *http.Client, name string) bool {
	url := fmt.Sprintf("https://www.flickr.com/people/%s/", name)
	fl.link = url
	resp, err := client.Get(url)
	if err != nil {
		fl.err = err
		return false
	}
	defer resp.Body.Close()

	return resp.StatusCode == 200
}

func (fl *FLCheck) Info() string {
	return "Check the username on flickr"
}

func (fl *FLCheck) Link() string {
	return fl.link
}

func (fl *FLCheck) Error() error {
	return fl.err
}

func init() {
	checker.Register("flickr", &FLCheck{})
}
