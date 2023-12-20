package misc

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
)

const port = "37899"

type EmployeeInfo struct {
	Name    string   `json:"Name"`
	Title   string   `json:"Title"`
	Reports []string `json:"Reports"`
}

type HTTPServer struct {
	Server       *http.Server
	EmployeeInfo map[string]EmployeeInfo
}

func NewHTTPServer() *HTTPServer {
	s := HTTPServer{
		Server: &http.Server{},
		EmployeeInfo: map[string]EmployeeInfo{
			"A123456789": {
				Name:    "Flynn Mackie",
				Title:   "Senior VP of Engineering",
				Reports: []string{"A100000001", "A100000002"},
			},
			"A100000001": {
				Name:    "Wesley Thomas",
				Title:   "VP of Design",
				Reports: []string{"A100000003"},
			},
			"A100000002": {
				Name:    "Nina Chiswick",
				Title:   "VP of Engineering",
				Reports: []string{"A100000005"},
			},
			"A100000003": {
				Name:    "Randall Cosmo",
				Title:   "Director of Design",
				Reports: []string{"A100000004"},
			},
			"A100000004": {
				Name:    "Brenda Plager",
				Title:   "Senior Designer",
				Reports: []string{},
			},
			"A100000005": {
				Name:    "Tommy Quinn",
				Title:   "Director of Engineering",
				Reports: []string{"A100000006", "A100000008"},
			},
			"A100000006": {
				Name:    "Jake Farmer",
				Title:   "Frontend Manager",
				Reports: []string{"A100000007"},
			},
			"A100000007": {
				Name:    "Liam Freeman",
				Title:   "Junior Code Monkey",
				Reports: []string{},
			},
			"A100000008": {
				Name:    "Sheila Dunbar",
				Title:   "Backend Manager",
				Reports: []string{"A100000009"},
			},
			"A100000009": {
				Name:    "Peter Young",
				Title:   "Senior Code Cowboy",
				Reports: []string{},
			},
		},
	}
	mux := http.NewServeMux()
	employeeIdUrl := "/employeedb/api/employee/"
	mux.HandleFunc(employeeIdUrl, func(w http.ResponseWriter, r *http.Request) {
		//fmt.Printf("Recieved request at %s\n", r.URL.Path)
		employeeId := strings.TrimPrefix(r.URL.Path, employeeIdUrl)
		//fmt.Printf("Fetched employee data %v\n", employeeId)
		jsonData, _ := json.Marshal(s.EmployeeInfo[employeeId])
		//fmt.Printf("Returning data(%v) back..\n", jsonData)
		w.Write(jsonData)
	})
	s.Server.Handler = mux
	s.Server.Addr = ":" + port
	return &s
}

func (h *HTTPServer) Start(ch chan error) {
	go func() {
		err := h.Server.ListenAndServe()
		if err != nil {
			ch <- err
			return
		}
	}()
	ch <- nil
}

func (h *HTTPServer) Stop(ctx context.Context) {
	h.Server.Shutdown(ctx)
}
