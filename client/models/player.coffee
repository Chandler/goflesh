App.Player = DS.Model.extend
  status: DS.attr 'string'
  game: DS.belongsTo 'App.Game'
  user: DS.belongsTo 'App.User'
  human_code: DS.attr 'string'

  isHuman: (->
    @get('status') == 'human'
  ).property('status')

  isZombie: (->
    @get('status') == 'zombie'
  ).property('status')

  isStarved: (->
    @get('status') == 'starved'    
  ).property('status')

App.Player.toString = -> 
  "Player"


