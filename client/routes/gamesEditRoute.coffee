define ["ember", "Game"], (Em, GameM) ->
  GamesShowRoute = Em.Route.extend
    model: (params) ->
      Game.find(params.game_id)

