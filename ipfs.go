package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	shell "github.com/ipfs/go-ipfs-api"
)

func ipfsReq() {
	// 	projectId := "2LS90Pe7Z19QqTwvqNSK07VBwqa"
	// 	projectSecret := "2adf483245dca8535a1dc316f45e8609"

	// 	sh := shell.NewShellWithClient("https://ipfs.infura.io:5001", NewClient(projectId, projectSecret))
	sh := shell.NewShell("localhost:5001")
	cid, err := sh.Add(strings.NewReader("Infura IPFS - Desoul.io Test"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Data successfully stored in IPFS: %v\n", cid)
}

// NewClient creates an http.Client that automatically perform basic auth on each request.
func NewClient(projectId, projectSecret string) *http.Client {
	return &http.Client{
		Transport: authTransport{
			RoundTripper:  http.DefaultTransport,
			ProjectId:     projectId,
			ProjectSecret: projectSecret,
		},
	}
}

// authTransport decorates each request with a basic auth header.
type authTransport struct {
	http.RoundTripper
	ProjectId     string
	ProjectSecret string
}

func (t authTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	r.SetBasicAuth(t.ProjectId, t.ProjectSecret)
	return t.RoundTripper.RoundTrip(r)
}
