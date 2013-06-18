define ["ember", "OrganizationModel"], (Em, OrganizationModel) ->
  DiscoveryRoute = Ember.Route.extend
    model: ->
      OrganizationModel.find()
      
    setupController: (controller, model) ->
      @controller.set('filterString', '')
