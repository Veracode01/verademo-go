package routes

import (
	"net/http"
	"verademo-go/src-app/controllers"

	"github.com/gorilla/mux"
)

type Page struct {
	Title string
	Body  []byte
}

/**
The routes file will create a Mux router, and then define all the
handlefunctions which will be run when a specified URL is received.
- Then, the handler functions can pass funcitonality to controllers,
so the controllers should have similar structure and be able to process a login, etc.

*/

// func viewHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Print("view")
// 	title := r.URL.Path[len("/view/"):]
// 	p, err := LoadPage(title)
// 	if err != nil {
// 		p = &Page{Title: title}
// 	}
// 	Render(w, "view", p)
// }

// func editHandler(w http.ResponseWriter, r *http.Request) {
// 	title := r.URL.Path[len("/edit/"):]
// 	p, err := LoadPage(title)
// 	if err != nil {
// 		p = &Page{Title: title}
// 	}
// 	Render(w, "edit", p)
// 	// t, _ := template.ParseFiles("edit.html")
// 	// t.Execute(w, p)
// }

func feedHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		controllers.ShowFeed(w, r)
	} else if r.Method == "POST" {

	}
}

func toolsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		controllers.ShowTools(w, r)
	} else if r.Method == "POST" {
		controllers.ProcessTools(w, r)
	}
}

/*
Handler function used by router for register page.
*/
func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		controllers.ShowRegister(w, r)
	} else if r.Method == "POST" {
		controllers.ProcessRegister(w, r)
	}
}

/*
Creates a router to listen for requests
Creates a session store
*/
func Routes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/feed", feedHandler)
	router.HandleFunc("/tools", toolsHandler)
	router.HandleFunc("/register", registerHandler)
	router.PathPrefix("/resources/").Handler(http.StripPrefix("/resources/", http.FileServer(http.Dir("resources"))))

	return router
}
