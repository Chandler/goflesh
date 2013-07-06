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

ember_namespace = [
  "Router",
  "Store",
  "IndexRoute",
  "GamesShowRoute",
  "GamesNewRoute",
  "GamesNewController",
  "GamesEditRoute",
  "GamesEditController",
  "Organization",
  "Game",
  "OrganizationsNewController",
  "OrganizationsNewRoute",
  "OrganizationsEditRoute",
  "OrganizationsEditController"
  "DiscoveryController",
  "DiscoveryRoute",
  "DiscoveryView",
  "ListItemView",
  "UsersNewController",
  "UsersNewRoute",
  "UsersShowRoute",
  "UsersEditRoute",
  "UsersEditController",
]
    
require ["underscore", "app"].concat(ember_namespace), (_, App, ember_namespace...) ->
  #This is the requirejs-ember secret sauce. Dyanamically set all ember objects to the ember App namespace.
  _.map _.zip(@ember_namespace, ember_namespace) , (pair) ->
    App.set(pair[0], pair[1])

  Em.App = App

  


