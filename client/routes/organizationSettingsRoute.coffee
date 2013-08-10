# define ["ember", "Organization", "User"], (Em, Organization ->
#   OrganizationSettingsRoute = Em.Route.extend
#     model: (params) ->
#       Organization.find(params.organization_id)
