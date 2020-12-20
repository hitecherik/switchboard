package walker

import (
	"fmt"

	"github.com/pelletier/go-toml"
)

type Redirect struct {
	Host        string
	Path        string
	Destination string
}

func Walk(tree *toml.Tree) ([]Redirect, error) {
	globalRedirects := make([]Redirect, 0)
	hostRedirects := make([]Redirect, 0)

	for _, key := range tree.Keys() {
		if destination, ok := tree.Get(key).(string); ok {
			globalRedirects = append(globalRedirects, Redirect{
				Path:        key,
				Destination: destination,
			})
		} else {
			fmt.Printf("Reading host %v\n", key)
			subtree := tree.GetPath([]string{key}).(*toml.Tree)

			var err error
			hostRedirects, err = walkWithHost(subtree, key, hostRedirects)
			if err != nil {
				return nil, err
			}
		}
	}

	return append(hostRedirects, globalRedirects...), nil
}

func walkWithHost(tree *toml.Tree, host string, redirects []Redirect) ([]Redirect, error) {
	for _, path := range tree.Keys() {
		destination, ok := tree.Get(path).(string)

		if !ok {
			return nil, fmt.Errorf("key \"%v\" in host \"%v\" is not a string", path, host)
		}

		redirects = append(redirects, Redirect{
			Host:        host,
			Path:        path,
			Destination: destination,
		})
	}

	return redirects, nil
}
