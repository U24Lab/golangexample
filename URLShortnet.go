//create two files a JSON and YAML file with Shortcut and URL Mapping

package main

import (
	"fmt"
	"net/http"
	s "strings"
)

type urlShortCut struct {
	shortcut string
	urls     string
}

func main() {

	jsonUrlMap := map[string]string{

		"site": "https://www.ui5cn.com",
		"blog": "https://www.blog.ui5cn.com",
	}
	jsonUrlMap["pricing"] = "https://www.ui5cn.com/pages/pricing"

	//http.HandleFunc("/", navigate)
	getURL(jsonUrlMap)
	http.ListenAndServe(":8080", nil)

}

func getURL(mapURL map[string]string) {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		userLink := r.URL.Path
		userURL := s.Split(userLink, "/")[1]
		if userURL == "" {
			fmt.Fprintf(w, "This is Home")
		} else if _, ok := mapURL[userURL]; ok {
			http.Redirect(w, r, mapURL[userURL], http.StatusFound)
		} else {
			fmt.Fprintf(w, "URL Not Found")
		}

		fmt.Println(userURL)

	})

}

func redirect(userUrl string){
	

// func navigate(w http.ResponseWriter, r *http.Request) {

// 	userLink := r.URL.Path
// 	fmt.Println(userLink)
// 	//http.Redirect(w, r, newUrl, http.StatusSeeOther)
// 	fmt.Fprintf(w, "Test Working")

// }