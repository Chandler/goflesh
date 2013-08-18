App.Organization = DS.Model.extend
  name: DS.attr 'string'
  slug: DS.attr 'string'
  location: DS.attr 'string'
  games: DS.hasMany 'App.Game',
    inverse: 'organization'

App.Organization.toString = -> 
  "Organization"