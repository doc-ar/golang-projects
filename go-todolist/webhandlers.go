package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func newTodoHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %d", err)
	}
	todo := r.FormValue("todo")

	if todo != "" {
		addTodo(todo)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
	// renderTodos(w, r)
}

func deleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %d", err)
	}

	id := r.FormValue("deleteId")
	id_int, err := strconv.Atoi(id)
	if err != nil {
		return
	}
	deleteTodo(id_int)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func renderTodos(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
<!doctype html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    <title></title>
    <link href="/static/style2.css" rel="stylesheet" />
  </head>
  <body>
    <div id="main" class="container">
      <div id="input_box" class="container">
        <form method="POST" action="/newtodo">
          <input
            name="todo"
            type="text"
            placeholder="What to do?"
            value=""
            id="input_box"
          />
          <input
            type="submit"
            name="submit"
            value="Add"
            id="submit_button"
          />
        </form>
      </div>
    <div id="items_list" class="container">
      <table>
        <tr>
          <th>#</th>
          <th>Name</th>
          <th>Delete</th>
        </tr>
    `)

	for i := range len(todos) {
		fmt.Fprintf(w, "<tr><td>%d</td><td>%s</td><td><form action=\"/deletetodo\" method=\"POST\"><input type=\"hidden\" name=\"deleteId\" value=\"%d\"><button type=\"submit\">&#x2715;</button></form></td></tr>\n", todos[i].id+1, todos[i].value, todos[i].id)
	}

	fmt.Fprintf(w, `
      <table>
      </div>
    </div>
  </body>
</html>
    `)
}
