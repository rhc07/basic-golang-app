package render

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

// template.FuncMap{} is empty. A map of functions that you can use in a template
var functions = template.FuncMap{}

func RenderTemplate(w http.ResponseWriter, tmpl string) {

	tc, err := CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)
	// ignoring an error here
	_ = t.Execute(buf, nil)

	_, err = buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser", err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	// Glob just returns the names of all the files matching a particular pattern
	pages, err := filepath.Glob("./templates/*.page.html")

	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		// this function fetches the name of the html page
		name := filepath.Base(page)
		fmt.Println("Page is currently", page)
		// ts is equal to template set
		// functions is the variable at the top of the page
		// ParseFiles(page) is passing the functions through the selected page.
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		}
		// If there is more than one match, parse the template set(ts) through the page using the ParseGlob function
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}
	return myCache, nil

}
