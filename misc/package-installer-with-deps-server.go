package misc

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/rs/zerolog/log"
)

/*
This implements the server side of the package installer.
*/

// Node represents a server identified by a name which
// has a list of packages installed. To ensure duplicate
// packages are not installed is a burden on the client.
// We use a slice (list) here and not a set so as to verify
// that the client managed dependencies well and not dupes
// were installed.
type Node struct {
	ServerName string
	Packages   []string
}

// PackageServer represents the web server that "interacts"
// with a set of servers. In reality its just a way to
// emulate which server was initialized and the packages
// that were installed on that server/Node.
type PackageServer struct {
	HTTPServer *http.Server
	Nodes      map[string]*Node
	Port       string
}

func NewNode(name string) *Node {
	return &Node{
		ServerName: name,
		Packages:   []string{},
	}
}

func NewPackageServer() *PackageServer {
	mux := http.NewServeMux()

	p := PackageServer{
		HTTPServer: &http.Server{},
		Nodes:      map[string]*Node{},
		Port:       "37899",
	}
	// For addnode and getpackages we plan to get the
	// param from the URL itself, hence the trailing "/"
	mux.HandleFunc("/addnode/", p.addNewNode)
	mux.HandleFunc("/getpackages/", p.getNodePackages)
	// For add package we fetch using the query params
	mux.HandleFunc("/addpackage", p.addPackage)

	p.HTTPServer.Addr = ":" + p.Port
	p.HTTPServer.Handler = mux

	return &p
}

// Adds a new Node on the package server
func (s *PackageServer) addNewNode(w http.ResponseWriter, r *http.Request) {
	nodeName := strings.TrimPrefix(r.URL.Path, "/addnode/")
	log.Debug().Str("nodeName", nodeName).Msg("Adding new node")
	s.Nodes[nodeName] = NewNode(nodeName)
	w.WriteHeader(http.StatusOK)
}

// Adds a package to an existing node on the package server
func (s *PackageServer) addPackage(w http.ResponseWriter, r *http.Request) {
	nodeName := r.URL.Query().Get("node")
	packageName := r.URL.Query().Get("package")
	log.Debug().Str("nodeName", nodeName).Str("packageName", packageName).Msg("Installing package on to node")
	if node, found := s.Nodes[nodeName]; !found {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("No such node"))
	} else {
		node.Packages = append(node.Packages, packageName)
		w.WriteHeader(http.StatusOK)
	}
}

// Lists the packages that were installed for a node
func (s *PackageServer) getNodePackages(w http.ResponseWriter, r *http.Request) {
	nodeName := strings.TrimPrefix(r.URL.Path, "/getpackages/")
	log.Debug().Str("nodeName", nodeName).Msg("Retrieving packages for node")
	packages := s.Nodes[nodeName].Packages
	response, err := json.Marshal(packages)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

// Start the package server
func (s *PackageServer) Start() error {
	log.Debug().Str("Port", s.Port).Msg("Starting server..")
	var err error
	go func() {
		err = s.HTTPServer.ListenAndServe()
	}()
	if err != nil {
		log.Err(err).Msg("Error starting the server..")
		return err
	}
	return nil
}

// Stop the package server
func (s *PackageServer) Stop(ctx context.Context) error {
	log.Debug().Msg("Stopping server..")
	err := s.HTTPServer.Shutdown(ctx)
	if err != nil {
		log.Err(err).Msg("Error starting the server..")
		return err
	}
	return nil
}
