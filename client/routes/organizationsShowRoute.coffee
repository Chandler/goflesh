define ["ember", "OrganizationModel"], (Em, OrganizationModel) ->
  OrganizationsShowRoute = Ember.Route.extend
    model: (params) ->
      OrganizationModel

    setupController: (controller, model) ->
      console.log(model)
      console.log(controller)

