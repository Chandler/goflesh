define ["ember", "templates", "ember-data"], (Em, Templates, DS) ->
  Game = DS.Model.extend
    name: DS.attr 'string'
    slug: DS.attr 'string'
    organziation: DS.hasMany 'Em.App.OrganizationModel'

  Game.toString = -> 
    "Game"

  Game