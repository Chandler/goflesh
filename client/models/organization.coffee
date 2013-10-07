App.Organization = DS.Model.extend
  name: DS.attr 'string'
  slug: DS.attr 'string'
  description: DS.attr 'string'
  location: DS.attr 'string'
  users: DS.hasMany 'App.User'
  games: DS.hasMany 'App.Game',
    inverse: 'organization'

  idaho: (->
    return @get('id') == "2"
  ).property()
  
    
App.Organization.toString = -> 
  "Organization"