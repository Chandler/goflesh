App.PlayerListView = Ember.ListView.extend({
  height: 200,
  height: 500,
  rowHeight: 50,
  itemViewClass: Ember.ListItemView.extend({templateName: "avatar"})
});