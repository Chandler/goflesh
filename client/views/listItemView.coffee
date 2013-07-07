define ["ember"], (Em) ->
  ListItemView = Em.View.extend
    templateName: "listItem"
    member_count: ->
      return 50
    didInsertElement: ->
      this.$().hide()
      this.$().fadeIn(100)
