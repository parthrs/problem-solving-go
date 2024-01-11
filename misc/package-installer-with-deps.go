package misc

/*
You've been asked to install a package on a set of target servers.
For each package there are a set of dependencies specified, those should
get installed first. Install all the packages without installing the
dependencies each time.

Idea:
- The dependencies will be specified for each package
- We need to keep track of what deps are installed on each server

Implementation:
- This implementation recursively calls self for installation
- We declarare a package struct that "ties" a package (name)
  and its deps together.
*/

import (
	"fmt"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

type Package struct {
	Name string
	Deps []*Package
}

func NewPackage(n string, d []*Package) *Package {
	return &Package{
		Name: n,
		Deps: d,
	}
}

var (
	InstallationStatus = make(map[string]string)
	InstallationURL    = "http://127.0.0.1:37899/addpackage"
)

// InstallApiCall makes a GET request to the server with query
// parameters - package name and a server/node name
func InstallApiCall(packageName, nodeName string) error {
	c := http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("GET", InstallationURL+fmt.Sprintf("?node=%s&package=%s", nodeName, packageName), nil)
	if err != nil {
		return err
	}

	resp, err := c.Do(req)
	if err != nil {
		return err
	}

	if statusCode := resp.StatusCode; statusCode >= 200 && statusCode <= 299 {
		return nil
	} else {
		return fmt.Errorf("recieved statusCode: %d", statusCode)
	}
}

// InstallPackage recursively calls self to install a package
// and its dependencies. If any of the deps fail to install
// the installation of the entire package is abandoned.
func InstallPackage(p *Package, nodeName string) (err error) {
	if _, lookup := InstallationStatus[p.Name]; lookup {
		log.Info().Str("Package", p.Name).Msg("Package already installed. Returning success..")
		return
	}

	log.Info().Str("Package", p.Name).Msg("Ensuring dependencies are installed for package")
	for _, pack := range p.Deps {
		log.Info().Str("Package", p.Name).Str("Dependency", pack.Name).Msg("Calling installation for dependency for package")
		err = InstallPackage(pack, nodeName)
		if err != nil {
			return
		}
	}

	log.Info().Str("Package", p.Name).Msg("Installing package")
	err = InstallApiCall(p.Name, nodeName)
	if err != nil {
		log.Info().Str("Package", p.Name).Str("Error", err.Error()).Msg("Failed installing package, returning")
		return
	}
	log.Info().Str("Package", p.Name).Msg("Updating install map with success for package")
	InstallationStatus[p.Name] = "Installed"
	return
}
