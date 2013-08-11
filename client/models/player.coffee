define ["ember-data"], (DS) ->
  Player = DS.Model.extend
  game: DS.belongsTo 'Em.App.Game'
  user: DS.belongsTo 'Em.App.User'
  
  Player.toString = -> 
    "Player"

  Player