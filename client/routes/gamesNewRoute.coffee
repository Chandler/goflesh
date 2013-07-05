define ["ember", "GameModel"], (Em, GameModel) ->
  GamesRoute = Ember.Route.extend
    model: ->
      GameModel
