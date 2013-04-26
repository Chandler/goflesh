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
          main: "ember.js"
      },
      {
          name: "ember-data"
          location: "lib"
          main: "ember-data.js"
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
  "DiscoveryController",
  "discoveryRoute",
  "discoveryView",
  "ListItemView",
  "ember-data"
], (App, DiscoveryController, DiscoveryRoute, DiscoveryView, ListItemView, DS) ->
  #this is where everything gets attached to our App

  App.Router = Em.Router.extend
    enableLogging: true
    location: 'history'

  App.Router.map ->
    this.route 'discovery'

  App.IndexRoute = Ember.Route.extend(redirect: ->
    @transitionTo 'discovery'
  )

  App.Store = DS.Store.extend
    revision: 11,
    adapter: DS.FixtureAdapter.create()

  App.ApplicationController = Ember.Controller.extend
    message: "this is the application template"

  App.set('ListItemView', ListItemView)
  App.set('DiscoveryView', DiscoveryView)
  App.set('DiscoveryRoute', DiscoveryRoute)
  App.set('DiscoveryController', DiscoveryController)

  Em.App = App

  


