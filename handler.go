func homeHandler(w http.ResponseWriter, r *http.Request) {
	err := rnd.Template(w, http.StatusOK, []string{"static/home.tpl"}, nil)
	checkErr(err)
}