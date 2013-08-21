App.GamesNewRoute = Ember.Route.extend
  model: ->
    App.Game

App.GameRoute = Ember.Route.extend
  model: (params) ->
    App.Game.find(params.game_id)

  setupController: (controller, model) ->
    @_super arguments...
    @controllerFor('games').set 'selectedGame', model

    
  