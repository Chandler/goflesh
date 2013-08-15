define ["ember-data"], (DS) ->
  Game = DS.Model.extend
    name: DS.attr 'string'
    slug: DS.attr 'string'
    players: DS.hasMany 'Em.App.Player'
  
  Game.toString = -> 
    "Game"

  Game