/*
Simple tools to make basic API endpoints
*/
package controllers

import (
	"fmt"
	"github.com/robfig/revel"
	"io"
	"io/ioutil"
	"reflect"
	"strings"
)

var (
	c = new(revel.Controller)
)

type createHelper func([]byte) ([]interface{}, error)

func CreateList(createModelList createHelper, body io.Reader) revel.Result {
	data, err := ioutil.ReadAll(body)
	if err != nil {
		return c.RenderError(err)
	}

	interfaces, err := createModelList(data)
	if err != nil {
		return c.RenderError(err)
	}

	// do the bulk insert
	err = Dbm.Insert(interfaces...)
	if err != nil {
		return c.RenderError(err)
	}

	// Return a copy of the data with id's set
	return c.RenderJson(interfaces)
}

// Expose a list view for a given model, with zeroed blacklisted field names
func GetList(model interface{}, blacklist []string) revel.Result {
	name := GetObjectName(model)

	template := `
    SELECT *
    FROM "%s"
    `
	query := fmt.Sprintf(template, name)
	result, err := Dbm.Select(model, query)
	if err != nil {
		return c.RenderError(err)
	}
	for _, item := range result {
		ZeroOutBlacklist(item, blacklist)
	}

	out := make(map[string]interface{})
	out[name+"s"] = result

	return c.RenderJson(out)
}

func ZeroOutBlacklist(item interface{}, blacklist []string) {
	concreteItem := reflect.ValueOf(item).Elem()
	for _, toBlack := range blacklist {
		val := concreteItem.FieldByName(toBlack)
		zero := reflect.Zero(val.Type())
		val.Set(zero)
	}
}

// Expose a list view for a given model, with zeroed blacklisted field names
func GetById(model interface{}, blacklist []string, id int) revel.Result {
	name := GetObjectName(model)

	result, err := Dbm.Get(model, id)
	if err != nil {
		return c.RenderError(err)
	}
	ZeroOutBlacklist(result, blacklist)

	out := make(map[string][]interface{})
	out[name] = []interface{}{result}

	return c.RenderJson(out)
}

// get an object name using introspection
// for example, models.Organization -> organization
func GetObjectName(obj interface{}) string {
	fullName := reflect.TypeOf(obj).String()
	name := strings.ToLower(strings.SplitN(fullName, ".", 2)[1])
	return name
}
