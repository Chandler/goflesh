define ["ember", "Game"], (Em, Game) ->
  GameRoute = Em.Route.extend
    model: (params) ->
      Game.find(params.game_id)

    setupController: (controller, model) ->
      @_super arguments...
      @controllerFor('games').set 'selectedGame', model
      # console.log model.get('organization')
      # @controllerFor('gameHome').set('orgs', model.get('organization'))
      # @controllerFor('gameHome').set('players', model.get('players'))
