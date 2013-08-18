App.Player = DS.Model.extend
  game: DS.belongsTo 'App.Game'
  user: DS.belongsTo 'App.User'

App.Player.toString = -> 
  "Player"


