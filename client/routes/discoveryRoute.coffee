define ["ember", "Organization"], (Em, Organization) ->
  DiscoveryRoute = Ember.Route.extend
    model: ->
      Organization.find()
      
    setupController: (controller, model) ->
      @controller.set('filterString', '')
