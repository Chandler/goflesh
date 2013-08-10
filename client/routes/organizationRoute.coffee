define ["ember", "Organization", "User"], (Em, Organization, User) ->
  OrganizationRoute = Em.Route.extend
    model: (params) ->
      console.log "organizationRoute"
      Organization.find(params.organization_id)
    # redirect: ->
    #   #set default tab for organization profile
    #   this.transitionTo('organization.home');
    
