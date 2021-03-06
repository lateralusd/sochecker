package olx

import (
	"fmt"
	"net/http"

	"github.com/lateralusd/sochecker/checker"
)

type OLXCheck struct {
	link string
	err  error
}

func (ol *OLXCheck) Check(client *http.Client, name string) bool {
	url := fmt.Sprintf("https://api.olx.ba/profil/%s/", name)
	ol.link = url
	resp, err := client.Get(url)
	if err != nil {
		ol.err = err
		return false
	}
	defer resp.Body.Close()

	return resp.StatusCode == 200
}

func (ol *OLXCheck) Info() string {
	return "Check the username on olx.ba"
}

func (ol *OLXCheck) Link() string {
	return ol.link
}

func (ol *OLXCheck) Error() error {
	return ol.err
}

func init() {
	checker.Register("olx", &OLXCheck{})
}
