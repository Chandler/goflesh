#requireJS bootstraper.
require.config
  baseUrl: "public/js"
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
  "organizationsController",
  "organizationsRoute"
  "discoveryController",
  "discoveryRoute",
  "discoveryView",
  "listItemView",
  "ember-data"
], (App, OrganizationsController, OrganizationsRoute, DiscoveryController, DiscoveryRoute, DiscoveryView, ListItemView, DS) ->
  #this is where everything gets attached to our App

  App.Router = Em.Router.extend
    enableLogging: true
    location: 'history'

  App.Router.map ->
    this.route 'discovery'
    this.route 'organizations', path: "/orgs"
    this.route 'signup'
  App.IndexRoute = Ember.Route.extend(redirect: ->
    @transitionTo 'discovery'
  )

  App.Store = DS.Store.extend
    revision: 12,
    adapter: DS.RESTAdapter.create()

  App.ApplicationController = Ember.Controller.extend
    message: "this is the application template"

  App.set('ListItemView', ListItemView)
  App.set('DiscoveryView', DiscoveryView)
  App.set('DiscoveryRoute', DiscoveryRoute)
  App.set('DiscoveryController', DiscoveryController)
  App.set('OrganizationsController', OrganizationsController)
  App.set('OrganizationsRoute', OrganizationsRoute)

  Em.App = App

  


