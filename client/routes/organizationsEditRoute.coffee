define ["ember", "Organization"], (Em, Organization) ->
  OrganizationsShowRoute = Ember.Route.extend
    model: (params) ->
      Organization.find(params.organization_id)
