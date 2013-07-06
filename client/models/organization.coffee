define ["ember", "templates", "ember-data"], (Em, Templates, DS) ->

  Organization = DS.Model.extend
    name: DS.attr 'string'
    slug: DS.attr 'string'
    location: DS.attr 'string'
    users: DS.hasMany 'Em.App.User'
    games: DS.hasMany 'Em.App.Game'


  Organization.toString = -> 
    "Organization"

  Organization