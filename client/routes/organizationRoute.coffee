App.OrganizationRoute = Ember.Route.extend
  model: (params) ->
    App.Organization.find(params.organization_id)

  #possibly not needed?
  # setupController: (controller, model) ->
  #   @_super arguments...
  #   @controllerFor('organizations').set 'selectedModel', model
