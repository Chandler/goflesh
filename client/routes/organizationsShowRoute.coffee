define ["ember", "Organization", "User"], (Em, Organization, User) ->
  OrganizationsShowRoute = Ember.Route.extend
    model: (params) ->
      Organization.find(params.organization_id)

