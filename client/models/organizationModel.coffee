define ["ember", "templates", "ember-data"], (Em, Templates, DS) ->

  Organization = DS.Model.extend
    name: DS.attr 'string'
    slug: DS.attr 'string'
    games: DS.hasMany 'Em.App.GameModel'

  Organization.toString = -> 
    "Organization"

  Organization