define ["ember", "gameModel"], (Em, GameModel) ->
  DiscoveryRoute = Ember.Route.extend
    model: ->
      GameModel.find()
    setupController: (controller, model) ->
      controller.set('message', model)
