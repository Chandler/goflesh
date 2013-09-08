App.GamesNewRoute = Ember.Route.extend
  model: ->
    App.Game

App.GameRoute = Ember.Route.extend
  events:
    joinGame: (game) ->
      players = game.get('players')
      createdPlayer = App.Player.createRecord
        game: game
        user: App.Auth.get('user')
      @get('store').get('defaultTransaction').commit()
      players.addObject(createdPlayer)
  
  model: (params) ->
    App.Game.find(params.game_id)

  setupController: (controller, model) ->
    events = App.GameEvent.find({ids: [model.get('id'),1]})
    @controllerFor('gameHome').set 'events', events
    @_super arguments...

    
  