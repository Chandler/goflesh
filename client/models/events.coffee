#this is on the refactor chopping block as soon as we upgrade to ember 1.0.0
#just look away

#models for each specific type of event
App.TagAttribute = DS.Model.extend
  tagger: DS.belongsTo 'App.Player' 
  taggee: DS.belongsTo 'App.Player'
  event:  DS.belongsTo 'App.Event'
App.TagAttribute.toString = -> 
  "TagAttribute"

#models for each type of event feed (player/game/organization)
App.Event = DS.Model.extend
  type: DS.attr 'string'
  tag:  DS.belongsTo 'App.TagAttribute', embedded:'always'
  player:  DS.belongsTo 'App.Player', embedded:'always'
App.Event.toString = -> 
  "Event"

  