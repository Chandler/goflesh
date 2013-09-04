#this isn't fantastic 


#models for each specific type of event
App.TagAttribute = DS.Model.extend
  tagger: DS.belongsTo 'App.Player' 
  taggee: DS.belongsTo 'App.Player'
  event:  DS.belongsTo 'App.PlayerEvent'
App.TagAttribute.toString = -> 
  "TagAttribute"


#mixin for attributes shared by each event feed
EventMixin = Ember.Mixin.create
  type: DS.attr 'string'
  tag:  DS.belongsTo 'App.TagAttribute', embedded:'always'

#models for each type of event feed (player/game/organization)
App.PlayerEvent = DS.Model.extend(EventMixin)
App.PlayerEvent.toString = -> 
  "PlayerEvent"


  
#   [
#   {
#     "type": "tag",
#     "tag": {
#       "id": 1,
#       "tagger_id": 1,
#       "taggee_id": 11,
#       "claimed": "2013-08-29T18:46:05.066734Z",
#       "created": "2013-08-30T01:46:05.067672Z",
#       "updated": "2013-08-30T01:46:05.067672Z"
#     }
#   }
# ]