#requireJS bootstraper.
require.config
  baseUrl: "public/js"
  packages: [
      {
          name: "handlebars",
          location: "lib",
          main: "handlebars.js"
      },
      {
          name: "ember",
          location: "lib",
          main: "ember.js"
      },
      {
          name: "ember-data",
          location: "lib",
          main: "ember-data.js"
      }
  ]

require ["jquery", "app", "gameModel", "templates"], ($, App, GameModel, Templates) ->

  App.Router.map ->
    this.route 'discovery'

  App.IndexRoute = Ember.Route.extend(redirect: ->
    @transitionTo "discovery"
  )

  App.ApplicationView = Ember.View.extend
    template: 'willy wonka'

  App.ApplicationController = Ember.Controller.extend()

  App.DiscoveryRouter = Ember.Route.extend
    model: -> 
      GameModel.find()
    setupController:  (controller, model) ->
      controller.set('content', model)
    renderTemplate: ->
      render('disovery');
  
  App.DiscoveryController = Em.ObjectController.extend
    message: 'This is the discovery page' 

  App.DiscoveryView = Ember.View.extend
    template: (name, data) ->
      Templates["discovery"]({message: "hey"})

  view = App.DiscoveryView.create
    controller: 'discovery'

  view.append("#app")
