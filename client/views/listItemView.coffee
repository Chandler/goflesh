App.ListItemView = Ember.View.extend
  templateName: "listItem"
  didInsertElement: ->
    this.$().hide()
    this.$().fadeIn(100)
