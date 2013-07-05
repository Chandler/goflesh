define ["ember"], (Em) ->
  ApplicationView = Em.View.extend
    didInsertElement: ->
     console.log "application view rendered"