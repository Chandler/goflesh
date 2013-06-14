define ["ember", "OrganizationModel"], (Em, OrganizationModel) ->
  DiscoveryRoute = Ember.Route.extend
    model: ->
      OrganizationModel
