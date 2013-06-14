/*
Simple tools to make basic API endpoints
*/
package controllers

import (
	"fmt"
	"github.com/robfig/revel"
	"reflect"
	"strings"
)

var (
	c = new(revel.Controller)
)

// Expose a list view for a given model, with zeroed blacklisted field names
func GetList(model interface{}, blacklist []string) revel.Result {
	name := GetObjectName(model)

	template := `
    SELECT *
    FROM "%s"
    `
	query := fmt.Sprintf(template, name)
	err := Dbm.Db.Ping()
	if err != nil {
		return c.RenderError(err)
	}
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

// get an object name using introspection
// for example, models.Organization -> organization
func GetObjectName(obj interface{}) string {
	fullName := reflect.TypeOf(obj).String()
	name := strings.ToLower(strings.SplitN(fullName, ".", 2)[1])
	return name
}
