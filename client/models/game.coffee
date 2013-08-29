App.Game = DS.Model.extend
  name: DS.attr 'string'
  slug: DS.attr 'string'
  organization: DS.belongsTo 'App.Organization'
  running_start_time: DS.attr 'string'
  players: DS.hasMany 'App.Player'
  myMethod: (->
    @get('players.length') > 0
  ).property()

  isCurrentUserGame: (->
   false
  ).property('players')


App.Game.toString = -> 
  "Game"