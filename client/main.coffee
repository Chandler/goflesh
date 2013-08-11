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

#plugins that need to run once to attach themselves to their parents.
require(['jquery-cookie','handlebars_helpers'])

#require parts of our framework not on the ember namespace
#we shouldn't need this, but things break without it
#tracking here: https://github.com/Chandler/flesh/issues/5
require(['NewController', 'BaseController'])

ember_namespace = [
  "Auth",
  "Router",
  "Store",
  
  #models
  "Game",
  "User",
  "Organization",
  "PasswordReset",
  
  #routes
  "DiscoveryRoute",
  "GameRoute",
  "GamesNewRoute",
  "UserRoute",
  "UsersNewRoute",
  "OrganizationRoute",
  "OrganizationsNewRoute",
  "PasswordResetRoute",
  "IndexRoute",
  
  
  #controllers
  "ApplicationController"
  "LoginController",
  "PasswordResetController",
  "SendPasswordResetController",
  "DiscoveryController",
  "OrganizationsController"
  "OrganizationSettingsController"
  "OrganizationsNewController",
  "UsersController",
  "UsersNewController",
  "UserSettingsController",
  "GamesController",
  "GamesNewController",
  "GameSettingsController",
  
  #views
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

  Em.App = App

  


