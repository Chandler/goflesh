App.PlayerEvent = DS.Model.extend
  name: DS.attr 'string'
  slug: DS.attr 'string'
  organization: DS.belongsTo 'App.Organization'
  players: DS.hasMany 'App.Player'
  running_start_time: DS.attr 'string'
App.Game.toString = -> 
  "Game"