package main

import (
	"fmt"
	"testing"

	"github.com/Orololuwa/collect_am-api/src/driver"
	"github.com/Orololuwa/collect_am-api/src/handlers"
	"github.com/go-chi/chi/v5"
)



func TestRoutes(t *testing.T){
	sql := driver.CreateTestDBInstance()
	conn := driver.DB{SQL: sql}
	h := handlers.NewTestHandlers(&testApp)

	mux := routes(&testApp, h, &conn)

	switch v := mux.(type) {
	case *chi.Mux:
		// do nothing
	default:
		t.Errorf(fmt.Sprintf("type is not *chi.Mux, but is %T", v))
	}
}

