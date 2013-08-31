App.PlayerEvent = DS.Model.extend
  type: DS.attr 'string'

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