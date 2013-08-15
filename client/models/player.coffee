define ["ember-data"], (DS) ->
  Player = DS.Model.extend
    game: DS.belongsTo 'App.Game'
    user: DS.belongsTo 'Em.App.User'

  Player.toString = -> 
    "Player"

  Player


