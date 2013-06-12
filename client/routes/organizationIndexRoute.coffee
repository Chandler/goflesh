define ["ember", "OrganizationModel"], (Em, OrganizationModel) ->
  OrganizationsShowRoute = Ember.Route.extend
    model: (params) ->
      console.log(params)

    setupController: (controller, model) ->
      console.log('')

