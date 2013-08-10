define ["ember", "Organization"], (Em, Organization) ->
  DiscoveryRoute = Em.Route.extend
    model: ->
      Organization
