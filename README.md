# README

Chapi exposes a generated api client for the Clubhouse (www.clubhouse.io) API. It provides an authorized client constructor `chapi.NewClient(<ch api key>)`

The client is generated from ./clubhouse.swagger.v3.json (sourced from https://clubhouse.io/api/rest/v3/clubhouse.swagger.json) using `go-swagger` (https://goswagger.io/).

The current impl was generated using `go-swagger v0.25.0 (f032690aab0634d97e2861a708d8fd9365ba77d2)`
