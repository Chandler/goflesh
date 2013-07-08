define ["ember", "Organization", "User"], (Em, Organization, User) ->
  OrganizationsShowRoute = Em.Route.extend
    model: (params) ->
      Organization.find(params.organization_id)

