package todo

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"sync"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/examples/todo/schema"
	"github.com/pkg/errors"
)

/*
	runs a todo graphql app
	each is a HTTP Post request to the port used with a body
	Get single todo: "{todo(id:\"b\"){id,text,done}}"
	Create new todo: "mutation+_{createTodo(text:\"My+new+todo\"){id,text,done}}"
	Update todo: "mutation+_{updateTodo(id:\"a\",done:true){id,text,done}}"
	Load todo list: "{todoList{id,text,done}}"

	here are the curls for these if the port is 8080
	curl -X POST -d '{"query":"mutation _{updateTodo(id:\"b\",done:true){id,text,done}}", "operationName":"Mutation"}' 'http://localhost:8080/graphql'
	curl -X POST -d '{"query":"mutation _{createTodo(text:\"My new todo\"){id,text,done}}", "operationName":"Mutation"}' 'http://localhost:8080/graphql'
	curl -X POST -d '{"query":"{todo(id:\"b\"){id,text,done}}"}' 'http://localhost:8080/graphql'
	curl -X POST -d '{"query":"{todoList{id,text,done}}"}' 'http://localhost:8080/graphql'
*/

// NewTodoServer will return a todo server
func NewTodoServer(port string) *TodoApp {
	return &TodoApp{port: port}
}

type TodoApp struct {
	port   string
	server *http.Server
	mutex  sync.Mutex
}

// Start will start the function, this is async. Returns a channel for async errors.
func (t *TodoApp) Start(ctx context.Context) (<-chan error, error) {
	t.setup()
	s, errs, err := startServer(ctx, t.port)
	if err != nil {
		return nil, err
	}
	errChan := make(chan error)
	go func() {
		select {
		case er := <-errs:
			errChan <- er
			t.Kill(ctx)
			break
		case <-ctx.Done():
			t.Kill(ctx)
			break
		}
	}()
	t.server = s
	return errChan, nil
}

// Kill the server, if it is running. This is thread safe.
func (t *TodoApp) Kill(ctx context.Context) error {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	if t.server != nil {
		err := t.server.Shutdown(ctx)
		t.server = nil
		return err
	}
	return nil
}

func (t *TodoApp) setup() {
	todo1 := schema.Todo{ID: "a", Text: "A todo not to forget", Done: false}
	todo2 := schema.Todo{ID: "b", Text: "This is the most important", Done: false}
	todo3 := schema.Todo{ID: "c", Text: "Please do this or else", Done: false}
	schema.TodoList = append(schema.TodoList, todo1, todo2, todo3)
}

type TodoHandler struct{}

type request struct {
	Query         string            `json:"query,omitempty"`
	OperationName string            `json:"operationName,omitempty"`
	Variables     map[string]string `json:"variables,omitempty"`
}

func (t *TodoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.Write([]byte(err.Error()))
		}
		r := request{}
		err = json.Unmarshal(body, &r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		} else {
			result := executeQuery(r.Query, schema.TodoSchema)
			json.NewEncoder(w).Encode(result)
		}
	} else {
		w.Write([]byte("NOT FOUND"))
	}
}

func startServer(ctx context.Context, port string) (*http.Server, <-chan error, error) {
	http.Handle("/graphql", &TodoHandler{})
	s := http.Server{
		Handler: &TodoHandler{},
		BaseContext: func(l net.Listener) context.Context {
			return ctx
		},
		Addr: fmt.Sprintf(":%s", port),
	}
	errs := make(chan error)
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			if !strings.Contains(err.Error(), http.ErrServerClosed.Error()) {
				errs <- errors.Wrap(err, "error with todo app")
			}
		}
	}()
	return &s, errs, nil
}

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}
