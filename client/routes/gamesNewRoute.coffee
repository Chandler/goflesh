define ["ember", "Game"], (Em, Game) ->
  GamesRoute = Em.Route.extend
    model: ->
      Game
