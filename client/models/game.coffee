define ["ember-data"], (DS) ->
  Game = DS.Model.extend
    name: DS.attr 'string'
    slug: DS.attr 'string'
    organziation: DS.belongsTo 'Em.App.Organization'

  Game.toString = -> 
    "Game"

  Game