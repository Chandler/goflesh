define ["ember", "OrganizationModel"], (Em, OrganizationModel) ->
  DiscoveryRoute = Ember.Route.extend
    model: ->
      OrganizationModel
      
    setupController: (controller, model) ->
      @controllerFor('discovery').set('content', model.find())
      @controllerFor('discovery').set('filterString', '')
