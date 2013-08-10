define ["ember", "Organization"], (Em, Organization) ->
  OrganizationRoute = Em.Route.extend
    model: (params) ->
      model = Organization.find(params.organization_id)
      @controllerFor('organizations').set 'selectedModel', model
