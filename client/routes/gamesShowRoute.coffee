define ["ember", "Game"], (Em, Game) ->
  GamesShowRoute = Em.Route.extend
    model: (params) ->
      Game.find(params.game_id)

