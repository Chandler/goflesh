define ["ember", "Game", "Player", "Organization"], (Em, Game, Player, Organization) ->
  GameRoute = Em.Route.extend
    model: (params) ->
      Game.find(params.game_id)

    setupController: (controller, model) ->
      @_super arguments...
      @controllerFor('games').set 'selectedGame', model
      console.log model.get('organization')
      @controllerFor('gameHome').set('orgs', model.get('organization'))
      @controllerFor('gameHome').set('gridModel', model.get('organization.members'))
