define ["ember", "Organization"], (Em, Organization) ->
  DiscoveryRoute = Ember.Route.extend
    model: ->
      Organization
