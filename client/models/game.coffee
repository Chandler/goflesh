App.Game = DS.Model.extend
  name: DS.attr 'string'
  slug: DS.attr 'string'
  organization: DS.belongsTo 'App.Organization'
  running_start_time: DS.attr 'string'
  players: DS.hasMany 'App.Player'
  
  isCurrentUserGame: (->
    user = App.Auth.get('user')
    @get('players').findProperty('user', user)
  ).property('players.@each.user')

App.Game.toString = -> 
  "Game"