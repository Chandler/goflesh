#this is on the refactor chopping block as soon as we upgrade to ember 1.0.0
#just look away

#models for each specific type of event
App.TagAttribute = DS.Model.extend
  tagger: DS.belongsTo 'App.Player' 
  taggee: DS.belongsTo 'App.Player'
  event:  DS.belongsTo 'App.PlayerEvent'
App.TagAttribute.toString = -> 
  "TagAttribute"

App.JoinedAttribute = DS.Model.extend
  player: DS.belongsTo 'App.Player'
  event:  DS.belongsTo 'App.PlayerEvent'
App.JoinedAttribute.toString = -> 
  "JoinedAttribute"


#mixin for attributes shared by each event feed
EventMixin = Ember.Mixin.create
  type: DS.attr 'string'
  tag:  DS.belongsTo 'App.TagAttribute', embedded:'always'
  joined:  DS.belongsTo 'App.JoinedAttribute', embedded:'always'

#models for each type of event feed (player/game/organization)
App.PlayerEvent = DS.Model.extend(EventMixin)
App.PlayerEvent.toString = -> 
  "PlayerEvent"


#models for each specific type of event
App.GameTagAttribute = DS.Model.extend
  tagger: DS.belongsTo 'App.Player' 
  taggee: DS.belongsTo 'App.Player'
  event:  DS.belongsTo 'App.GameEvent'
App.GameTagAttribute.toString = -> 
  "GameTagAttribute"


#mixin for attributes shared by each event feed
GameEventMixin = Ember.Mixin.create
  type: DS.attr 'string'
  tag:  DS.belongsTo 'App.GameTagAttribute', embedded:'always'
  player:  DS.belongsTo 'App.Player', embedded:'always'


App.GameEvent = DS.Model.extend(GameEventMixin)
App.GameEvent.toString = -> 
  "GameEvent"


  