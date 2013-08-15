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
      },
      {
          name: "templates"
          location: "."
          main: "templates.js"
      },
      {
          name: "ember-grid"
          location: "lib"
          main: "ember-grid.js"
      }
  ],
  shim:
    "templates":
      exports: 'this["Ember"]["TEMPLATES"]'

#plugins that need to run once to attach themselves to their parents.
require(['jquery-cookie','handlebars_helpers', 'ember-grid'])

#require parts of our framework not on the ember namespace
#we shouldn't need this, but things break without it
#tracking here: https://github.com/Chandler/flesh/issues/5
require(['NewController', 'BaseController'])

#Any file you create that needs be on the ember namespace must be listed here
ember_namespace = [
  "Auth",
  "Router",
  "Store",
  
  # Models
  "Player",
  "Member",
  "Game",
  "User",
  "Organization",
  
  # Routes
  "DiscoveryRoute",
  
  "GameRoute",
  "GamesNewRoute",
  
  "UserRoute",
  "UsersNewRoute",
  
  "OrganizationRoute",
  "OrganizationsNewRoute",
  "IndexRoute",

  # Controllers
  "ApplicationController"
  "LoginController",
  "DiscoveryController",

  "OrganizationHomeController",
  "OrganizationSettingsController",
  "OrganizationsController",
  "OrganizationsNewController",
  
  "UsersController",
  "UsersNewController",
  "UserHomeController",
  "UserSettingsController",
  
  "GamesController",
  "GamesNewController",
  "GameSettingsController",
  "GameHomeController",

  # Views
  "RegisterKillView",
  "TimeSeriesView",
  "PlayerTableRowView",
  "PlayerListTableView",  
  "ApplicationView",
  "LoginView",
  "DiscoveryView",
  "ListItemView",
  "GameItemView"

]
    
require ["underscore", "app"].concat(ember_namespace), (_, App, ember_namespace...) ->
  #This is the requirejs-ember secret sauce. Dyanamically set all ember objects to the ember App namespace.
  _.map _.zip(@ember_namespace, ember_namespace) , (pair) ->
    App.set(pair[0], pair[1])

  window.App = App
  Em.App = App

  


