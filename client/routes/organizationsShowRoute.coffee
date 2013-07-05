define ["ember", "OrganizationModel"], (Em, OrganizationModel) ->
  OrganizationsShowRoute = Ember.Route.extend
    model: (params) ->
      OrganizationModel.find(params.organization_id)

