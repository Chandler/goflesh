define ["ember", "Game"], (Em, Game) ->
  GamesShowRoute = Ember.Route.extend
    model: (params) ->
      Game.find(params.game_id)

