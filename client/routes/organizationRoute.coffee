App.OrganizationRoute = Ember.Route.extend
  model: (params) ->
    App.Organization.find(params.organization_id)

  setupController: (controller, model) ->
    @_super arguments...
    @controllerFor('organizations').set 'selectedOrganization', model
