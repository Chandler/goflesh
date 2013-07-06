define ["ember", "OrganizationModel"], (Em, OrganizationModel) ->
  OrganizationsShowRoute = Ember.Route.extend
    model: (params) ->
      a = OrganizationModel.find(params.organization_id)
      console.log a
      a