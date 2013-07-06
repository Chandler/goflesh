define ["ember", "Game"], (Em, GameM) ->
  GamesShowRoute = Ember.Route.extend
    model: (params) ->
      Game.find(params.game_id)

