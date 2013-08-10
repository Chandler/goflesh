define ["ember", "Organization"], (Em, Organization) ->
  OrganizationRoute = Em.Route.extend
    model: (params) ->
      Organization.find(params.organization_id)

    setupController: (controller, model) ->
      @_super arguments...
      @controllerFor('organizations').set 'selectedModel', model
