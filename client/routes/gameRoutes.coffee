App.GamesNewRoute = Ember.Route.extend
  model: ->
    App.Game

App.GameRoute = Ember.Route.extend
  model: (params) ->
    App.Game.find(params.game_id)

  setupController: (controller, model) ->
    event = App.PlayerEvent.find({ids: [1,2,3]})
    console.log event
    @controllerFor('gameHome').set 'events', event
    @_super arguments...

    
  