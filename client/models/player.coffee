define ["ember-data"], (DS) ->
  Player = DS.Model.extend
    game: DS.belongsTo 'Em.App.Game'
    user: DS.belongsTo 'Em.App.User'
    get_user: (->
      @get('user')
    ).property()
  Player.toString = -> 
    "Player"

  Player


