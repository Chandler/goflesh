App.Organization = DS.Model.extend
  name: DS.attr 'string'
  slug: DS.attr 'string'
  location: DS.attr 'string'
  members: DS.hasMany 'App.Member'
  games: DS.hasMany 'App.Game',
    inverse: 'organization'

App.Organization.toString = -> 
  "Organization"