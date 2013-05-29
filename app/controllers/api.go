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
	// get the model name using introspection
	// for example, models.Organization -> Organization
	fullName := reflect.TypeOf(model).String()
	name := strings.ToLower(strings.SplitN(fullName, ".", 2)[1])

	template := `
    SELECT *
    FROM "%s"
    `
	query := fmt.Sprintf(template, name)

	result, err := dbm.Select(model, query)
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
