App.OrganizationsNewRoute = Ember.Route.extend
  model: ->
    App.Organization

App.OrganizationRoute = Ember.Route.extend
  events:
    leaveGame: (game) ->
      players = game.get('players')
      createdPlayer = App.Player.createRecord
        game: game
        user: App.Auth.get('user')
      @get('store').get('defaultTransaction').commit()
      players.addObject(createdPlayer)
    joinGame: (game) ->
      players = game.get('players')
      createdPlayer = App.Player.createRecord
        game: game
        user: App.Auth.get('user')
      @get('store').get('defaultTransaction').commit()
      players.addObject(createdPlayer)
  model: (params) ->
    App.Organization.find(params.organization_id)
