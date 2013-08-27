App.EventFeed = DS.Model.extend
  name: DS.attr 'string'
  slug: DS.attr 'string'
  organization: DS.belongsTo 'App.Organization'
  players: DS.hasMany 'App.Player'

App.EventFeed.toString = -> 
  "EventFeed"