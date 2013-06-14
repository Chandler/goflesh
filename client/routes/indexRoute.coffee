define ["ember"], (Em) ->
  IndexRoute = Ember.Route.extend(redirect: ->
    @transitionTo 'discovery'
  )