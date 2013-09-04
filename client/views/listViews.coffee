#
# Container views for customizing the display of lists on different pages
#

# Game page container
App.GamePageListView = Ember.View.extend
  templateName: 'lists/gamePageList'

# Containers for user and org profiles
App.OrgProfileListView = Ember.View.extend
  templateName: 'lists/profileList'


#
# Player List
#
App.PlayerRowView = Ember.ListItemView.extend
  templateName: "playerList/playerListRow"


App.PlayerListView = Ember.ListView.extend
  height: 1400,
  rowHeight: 50,
  itemViewClass: App.PlayerRowView


#
# Custom Mixins for each supported event type
#


# JoinGameEventMixin = Ember.Mixin.create
#   setupTemplateData: ->
#     @set 'player', App.Player.find(2)


# TagEventMixin = Ember.Mixin.create
#   setupTemplateData: ->
#     debugger
#     @set 'tagger', App.Player.find(2)
#     @set 'taggee', App.Player.find(2)


#
# Event Feed
#
App.EventRowView = Ember.ListItemView.extend
  templateName: (->
    # rowMixins = {
    #   tag: TagEventMixin,
    #   joingame: JoinGameEventMixin
    # }

    # rowMixins[@get('context.type')].apply(@)
    
    # @setupTemplateData()
    rowTemplate = @get('context.type')
    "eventList/" + rowTemplate
  ).property()


App.EventListView = Ember.ListView.extend
  height: 1400,
  rowHeight: 50,
  itemViewClass: (->
    App.EventRowView
  ).property()


