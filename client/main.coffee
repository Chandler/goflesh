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

require ["jquery", "app"], ($, App) ->

  App.Router.map ->
    this.route 'discovery'

  # DiscoveryRoute = Em.Route.Extend
  #   model: ->
  #     return GameModel.find() 

  DiscoveryController = Em.ArrayController.extend

