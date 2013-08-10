define ["ember", "Organization"], (Em, Organization) ->
  OrganizationsShowRoute = Em.Route.extend
    model: (params) ->
      Organization.find(params.organization_id)
