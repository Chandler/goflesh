App.Organization = DS.Model.extend
  name: DS.attr 'string'
  slug: DS.attr 'string'
  location: DS.attr 'string'
  users: DS.hasMany 'App.User'
  games: DS.hasMany 'App.Game',
    inverse: 'organization'

App.Organization.toString = -> 
  "Organization"