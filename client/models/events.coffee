App.PlayerEvent = DS.Model.extend
  name: DS.attr 'string'
App.PlayerEvent.toString = -> 
  "PlayerEvent"