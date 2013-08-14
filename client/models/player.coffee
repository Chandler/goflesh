define ["ember-data"], (DS) ->
  Player = DS.Model.extend
    game: DS.belongsTo 'Em.App.Game'
    
  Player.toString = -> 
    "Player"

  Player


