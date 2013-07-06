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
          main: "ember.js"
      },
      {
          name: "ember-auth"
          location: "lib"
          main: "ember-auth.js"

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

#jquery-cookie needs to run once to attach itself to jquery.
require(['jquery-cookie', 'handlebars_helpers'])

ember_namespace = [
  "Auth",
  "Router",
  "Store",
  
  #models
  "Game",
  "User",
  "Organization",
  
  #routes
  "GamesEditRoute",
  "GamesNewRoute",
  "GamesShowRoute",
  "IndexRoute",
  "UsersEditRoute",
  "UsersNewRoute",
  "UsersShowRoute",
  "DiscoveryRoute",
  "OrganizationsNewRoute",
  "OrganizationsShowRoute",
  "OrganizationsEditRoute",
  
  #controllers
  "ApplicationController"
  "LoginController",
  "DiscoveryController",
  "OrganizationsNewController",
  "OrganizationsEditController"
  "UsersNewController",
  "UsersEditController",
  "GamesNewController",
  "GamesEditController",
  
  #views
  "ApplicationView",
  "LoginView",
  "DiscoveryView",
  "ListItemView",

]
    
require ["underscore", "app"].concat(ember_namespace), (_, App, ember_namespace...) ->
  #This is the requirejs-ember secret sauce. Dyanamically set all ember objects to the ember App namespace.
  _.map _.zip(@ember_namespace, ember_namespace) , (pair) ->
    App.set(pair[0], pair[1])

  Em.App = App

  


