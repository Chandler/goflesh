App.DiscoveryRoute = Ember.Route.extend
  model: ->
    App.Organization.find()
  setupController: (controller, model) ->
    @controller.set('content', model)
    @controller.set('filterString', '')
