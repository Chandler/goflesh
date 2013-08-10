define ["ember", "Organization"], (Em, Organization) ->
  DiscoveryRoute = Em.Route.extend
    model: ->
      Organization.find()
      
    setupController: (controller, model) ->
      @controller.set('filterString', '')
