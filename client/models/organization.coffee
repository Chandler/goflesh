App.Organization = DS.Model.extend
  name: DS.attr 'string'
  slug: DS.attr 'string'
  location: DS.attr 'string'
  users: DS.hasMany 'App.User'
  games: DS.hasMany 'App.Game',
    inverse: 'organization'
  myMethod: (->
    @get('games.length') > 0
  ).property()
App.Organization.toString = -> 
  "Organization"