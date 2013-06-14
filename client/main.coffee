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

ember_namespace = [
  "OrganizationsShowController",
  "OrganizationsShowRoute",
  "OrganizationsNewController",
  "OrganizationsNewRoute",
  "DiscoveryController",
  "DiscoveryRoute",
  "DiscoveryView",
  "ListItemView",
]
    
require ["underscore", "app", "ember-data"].concat(ember_namespace), (_, App, DS, ember_namespace...) ->
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
    adapter: JSONAPIAdapter.create({ namespace: 'api' })

  #dynamically set all the ember objects on to the App"
  _.map _.zip(@ember_namespace, ember_namespace) , (pair) ->
    App.set(pair[0], pair[1])

  Em.App = App

  


