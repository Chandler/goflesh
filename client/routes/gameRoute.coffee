define ["ember", "Game", "Organization"], (Em, Game, Organization) ->
  GameRoute = Em.Route.extend
    model: (params) ->
      Game.find(params.game_id)

    setupController: (controller, model) ->
      @_super arguments...
      @controllerFor('games').set 'selectedGame', model
      @controllerFor('gameHome').set('gridModel', Organization.find({game_id:  model.id}))
      @controllerFor('gameHome').set('chartValues', [1,2,1,3])