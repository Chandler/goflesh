App.OrganizationsNewRoute = Ember.Route.extend
  model: ->
    App.Organization

App.OrganizationRoute = Ember.Route.extend
  model: (params) ->
    App.Organization.find(params.organization_id)