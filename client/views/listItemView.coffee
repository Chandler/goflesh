define ["ember"], (Em) ->
  ListItemView = Em.View.extend
    templateName: "listItem"
    didInsertElement: ->
      this.$().hide()
      this.$().fadeIn(100)
