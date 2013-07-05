define ["ember", "templates", "ember-data", "OrganizationModel"], (Em, Templates, DS, OrganizationModel) ->
  Game = DS.Model.extend
    name: DS.attr 'string'
    slug: DS.attr 'string'

  Game.toString = -> 
    "Game"

  Game