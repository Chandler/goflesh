define ["ember", "OrganizationModel"], (Em, OrganizationModel) ->
  OrganizationsShowRoute = Ember.Route.extend
    model: (params) ->
      OrganizationModel.find(params.id)

    setupController: (controller, model) ->
      console.log(model)
      console.log(controller)

