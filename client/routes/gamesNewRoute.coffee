define ["ember", "Game"], (Em, Game) ->
  GamesRoute = Ember.Route.extend
    model: ->
      Game
