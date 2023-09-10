package framework

import (
	"net/url"
)

var (
	DevInternalURL  *url.URL
	ProdInternalURL *url.URL
)

func init() {
	var err error

	DevInternalURL, err = new(url.URL).Parse("http://localhost:20041/")
	if err != nil {
		panic(err.Error())
	}

	ProdInternalURL, err = new(url.URL).Parse("https://agora-forum-service-internal-tgsqc5aoiq-od.a.run.app/")
	if err != nil {
		panic(err.Error())
	}
}
