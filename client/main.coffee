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
      deps: ['jquery']
      exports: 'this["Ember"]["TEMPLATES"]'
    
require ["jquery", "app", "gameModel", "templates", "ember"], ($, App, GameModel, Templates, Em) ->
  debugger
  App.Router.map ->
    this.route 'discovery'

  App.IndexRoute = Ember.Route.extend(redirect: ->
    @transitionTo 'discovery'
  )

  App.ApplicationController = Ember.Controller.extend
    message: "this is the APPLICATION CONTROLLER"

  App.DiscoveryRoute = Ember.Route.extend
    setupController: (controller) ->
      controller.set('message', 'sup')
    # model: ->
    #   GameModel.find()
    renderTemplate: ->
      this.render

  # App.DiscoveryController = Em.ObjectController.extend
  #   message: 'This is the discovery page'

  # App.DiscoveryView = Ember.View.extend
  #   template: Ember.Handlebars.compile('Hello {{message}}')

  # view = App.DiscoveryView.create
  #   controller: App.DiscoveryController

  # view.append("#app")
