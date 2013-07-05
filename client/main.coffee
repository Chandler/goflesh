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
require(['jquery-cookie'])

ember_namespace = [
  "Auth",
  "Router",
  "Store",
  "IndexRoute",
  "GamesShowRoute",
  "OrganizationModel",
  "GameModel",
  "UserModel",
  "OrganizationsShowController",
  "OrganizationsShowRoute",
  "OrganizationsNewController",
  "OrganizationsNewRoute",
  "DiscoveryController",
  "DiscoveryRoute",
  "DiscoveryView",
  "ListItemView",
  "UsersNewController",
  "UsersNewRoute",
  "UsersShowController",
  "UsersShowRoute",
]
    
require ["underscore", "app"].concat(ember_namespace), (_, App, ember_namespace...) ->
  #This is the requirejs-ember secret sauce. Dyanamically set all ember objects to the ember App namespace.
  _.map _.zip(@ember_namespace, ember_namespace) , (pair) ->
    App.set(pair[0], pair[1])

  Em.App = App

  


