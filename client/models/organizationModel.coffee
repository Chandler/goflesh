define ["ember", "templates", "ember-data"], (Em, Templates, DS) ->

  Organization = DS.Model.extend
    name: DS.attr 'string'
    slug: DS.attr 'string'
    location: DS.attr 'string'
    games: DS.hasMany 'Em.App.GameModel'
    users: DS.hasMany 'Em.App.UserModel'

  Organization.toString = -> 
    "Organization"

  Organization