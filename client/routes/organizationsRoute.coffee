define ["ember", "organizationModel"], (Em, OrganizationModel) ->
  OrganizationsRoute = Ember.Route.extend
    model: ->
      OrganizationModel
    setupController: (controller, model) ->
      this.controllerFor('discovery').set('model', model)
