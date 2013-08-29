App.Player = DS.Model.extend
  status: DS.attr 'string'
  game: DS.belongsTo 'App.Game'
  user: DS.belongsTo 'App.User'

  isHuman: (->
    @get('status') == 'human'
  ).property()

  isZombie: (->
    @get('status') == 'zombie'
  ).property()

  isStaved: (->
    @get('status') == 'starved'    
  ).property()

App.Player.toString = -> 
  "Player"


