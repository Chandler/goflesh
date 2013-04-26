define ["ember", "OrganizationModel"], (Em, OrganizationModel) ->
  DiscoveryRoute = Ember.Route.extend
    model: ->
      OrganizationModel
    setupController: (controller, model) ->
      this.controllerFor('discovery').set('content', model.find())
