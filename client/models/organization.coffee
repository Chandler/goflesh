define ["ember-data"], (DS) ->

  Organization = DS.Model.extend
    name: DS.attr 'string'
    slug: DS.attr 'string'
    location: DS.attr 'string'
    members: DS.hasMany 'Em.App.Member'
    games: DS.hasMany 'Em.App.Game'
    
  Organization.toString = -> 
    "Organization"

  Organization  
