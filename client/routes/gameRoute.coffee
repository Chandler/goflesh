define ["ember", "Game"], (Em, Game) ->
  GameRoute = Em.Route.extend
    model: (params) ->
      Game.find(params.organization_id)

    setupController: (controller, model) ->
      @_super arguments...
      @controllerFor('games').set 'selectedGame', model
