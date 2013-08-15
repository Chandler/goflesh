define ["ember", "Game", "Player", "User"], (Em, Game, Player, User) ->
  GameRoute = Em.Route.extend
    model: (params) ->
      Game.find(params.game_id)

    setupController: (controller, model) ->
      @_super arguments...
      @controllerFor('games').set 'selectedGame', model
      @controllerFor('gameHome').set('players', Player.find({game_id:  model.id}))
      @controllerFor('gameHome').set('organization', model.get('organization'))
