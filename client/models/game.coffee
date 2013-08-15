define ["ember-data"], (DS) ->
  Game = DS.Model.extend
    name: DS.attr 'string'
    slug: DS.attr 'string'
    organization: DS.belongsTo 'Em.App.Organization'
    players: DS.hasMany 'Em.App.Player'
  
  Game.toString = -> 
    "Game"

  Game