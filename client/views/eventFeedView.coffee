#A container view for the event feed
App.EventFeedView = Ember.View.extend
  templateName: 'eventFeed/container'

#View for each row in the list.
App.EventFeedRowView = Ember.ListItemView.extend
  templateName: "eventTypes/playerTaggedPlayer"

#The actual Ember.ListView
App.EventFeedListView = Ember.ListView.extend
  eventTemplate: 'hey'
  height: 1400,
  rowHeight: 50,
  itemViewClass: App.EventFeedRowView

