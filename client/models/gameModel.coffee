define ["ember", "templates", "ember-data"], (Em, Templates, DS) ->
  Game = DS.Model.extend
    name: DS.attr 'string'
    slug: DS.attr 'string'
  Game.toString = -> 
    "Game"

  Game