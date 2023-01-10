package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
)

func main() {
	// Get the repository URL and authentication method from the command line arguments
	url := os.Args[1]
	authMethod := os.Args[2]

	var auth transport.AuthMethod

	// Set the appropriate authentication method
	if authMethod == "ssh" {
		privateKey, err := ioutil.ReadFile(os.Getenv("HOME") + "/.ssh/id_rsa")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		signer, err := ssh.NewSignerFromKey(privateKey)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		auth = &ssh.PublicKeys{User: "git", Signer: signer}
	} else if authMethod == "token" {
		token := os.Getenv("GITHUB_TOKEN")
		if token == "" {
			fmt.Println("GITHUB_TOKEN environment variable is not set")
			os.Exit(1)
		}
		auth = &http.BasicAuth{
			Username: "x-access-token",
			Password: token,
		}
	} else {
		fmt.Printf("Invalid authentication method: %s\n", authMethod)
		os.Exit(1)
	}

	// Clone the repository
	_, err := git.PlainClone(os.TempDir()+"/repo", false, &git.CloneOptions{
		URL:           url,
		Auth:          auth,
		ReferenceName: plumbing.ReferenceName("refs/heads/master"),
		SingleBranch:  true,
	})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
