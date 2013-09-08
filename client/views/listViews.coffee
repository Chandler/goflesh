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
  height: 1000, # change to bigger on Players List Page
  rowHeight: 80,
  adjustLayout: (new_width, new_height) -> 
    @set('width', new_width)
    @set('height', new_height)
  itemViewClass: App.PlayerRowView


App.EventRowView = Ember.ListItemView.extend
  templateName: (->
    rowTemplate = @get('context.type')
    "eventList/" + rowTemplate
  ).property()


App.EventListView = Ember.ListView.extend
  height: 1000,
  rowHeight: 80,
  adjustLayout: (new_width, new_height) -> 
    @set('width', new_width)
    @set('height', new_height)
  itemViewClass: App.EventRowView


