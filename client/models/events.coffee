App.PlayerEvent = DS.Model.extend
  type: DS.attr 'string'
App.PlayerEvent.toString = -> 
  "PlayerEvent"