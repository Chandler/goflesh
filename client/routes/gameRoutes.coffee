App.GamesNewRoute = Ember.Route.extend
  model: ->
    App.Game

App.GameRoute = Ember.Route.extend
  model: (params) ->
    App.Game.find(params.game_id)

  setupController: (controller, model) ->
    events = App.PlayerEvent.find({ids: [1,2,3]})
    @controllerFor('gameHome').set 'events', events
    @_super arguments...

    
  