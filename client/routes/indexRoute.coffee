define ["ember"], (Em) ->
  IndexRoute = Em.Route.extend(redirect: ->
    @transitionTo 'discovery'
  )