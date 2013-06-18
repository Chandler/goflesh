define ["ember", "OrganizationModel"], (Em, OrganizationModel) ->
  OrganizationsShowRoute = Ember.Route.extend
    model: (params) ->
      debugger
      OrganizationModel.find(params.id)
