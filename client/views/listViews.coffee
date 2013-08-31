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
# Event Feed
#
App.EventRowView = Ember.ListItemView.extend

  templateName: (=>
    rowTemplate = "tag"
    "eventList/" + rowTemplate
  ).property()

App.EventListView = Ember.ListView.extend
  height: 1400,
  rowHeight: 50,
  itemViewClass: (->
    App.EventRowView
  ).property()


