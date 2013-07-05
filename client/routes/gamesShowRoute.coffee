define ["ember", "GameModel"], (Em, GameModel) ->
  GamesShowRoute = Ember.Route.extend
    model: (params) ->
      GameModel.find(params.game_id)

