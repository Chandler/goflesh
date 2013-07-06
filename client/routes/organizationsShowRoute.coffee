define ["ember", "Organization"], (Em, Organization) ->
  OrganizationsShowRoute = Ember.Route.extend
    model: (params) ->
      Organization.find(params.organization_id)
    # setupController: (controller, model) ->
    #   first = model.get('games')
    #   console.log first
    #   controller.set('stuff', first)
