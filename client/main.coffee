#requireJS bootstraper.
require.config
  baseUrl: "/public/js"
  packages: [
      {
          name: "handlebars"
          location: "lib"
          main: "handlebars.js"
      },
      {
          name: "ember"
          location: "lib"
          main: "new-ember.js"
      },
      {
          name: "ember-data"
          location: "lib"
          main: "new-ember-data.js"
      }
      {
          name: "templates"
          location: "."
          main: "templates.js"
      }
  ],
  shim:
    "templates":
      exports: 'this["Ember"]["TEMPLATES"]'

    
require [
  "app",
  "OrganizationsShowController",
  "OrganizationsShowRoute",
  "organizationsController",
  "discoveryController",
  "discoveryRoute",
  "discoveryView",
  "listItemView",
  "ember-data",
], (App, OrganizationsShowController, OrganizationsShowRoute, OrganizationsController, DiscoveryController, DiscoveryRoute, DiscoveryView, ListItemView, DS) ->
  #this is where everything gets attached to our App

  App.Router = Em.Router.extend
    enableLogging: true
    location: 'history'

  #http://darthdeus.github.io/blog/2013/02/01/ember-dot-js-router-and-template-naming-convention/
  App.Router.map ->
    @route 'discovery'
    @resource 'organizations', path: "/orgs", ->
      @route 'show', path: ":id", ->
      @route 'new'
    @route 'signup'
  
  App.IndexRoute = Ember.Route.extend(redirect: ->
    @transitionTo 'discovery'
  )

  JSONAPIAdapter = DS.RESTAdapter.extend()

  App.Store = DS.Store.extend
    adapter: JSONAPIAdapter.create()

  App.ApplicationController = Ember.Controller.extend
    message: "this is the application template"

  App.set('ListItemView', ListItemView)
  App.set('DiscoveryView', DiscoveryView)
  App.set('DiscoveryRoute', DiscoveryRoute)
  App.set('DiscoveryController', DiscoveryController)
  App.set('OrganizationsShowRoute', OrganizationsShowRoute)
  App.set('OrganizationsController', OrganizationsController)
  App.set('OrganizationsShowController', OrganizationsShowController)

  Em.App = App

  


