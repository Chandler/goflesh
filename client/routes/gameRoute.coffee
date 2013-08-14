define ["ember", "Game", "Player"], (Em, Game, Player) ->
  GameRoute = Em.Route.extend
    model: (params) ->
      Game.find(params.game_id)

    setupController: (controller, model) ->
      @_super arguments...
      @controllerFor('games').set 'selectedGame', model
      @controllerFor('gameHome').set('gridModel', Player.find({game_id:  model.id}))
