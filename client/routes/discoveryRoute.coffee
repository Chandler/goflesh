App.DiscoveryRoute = Ember.Route.extend
  model: ->
    App.Organization.find()
    
  setupController: (controller, model) ->
    @controller.set('filterString', '')
