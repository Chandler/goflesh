define ["ember"], (Em) ->
  DiscoveryController = Ember.ObjectController.extend
    organizations: (->
      string = @get 'filterString'
      if string == ""
        return @get 'content'
      else
        @get('content').filter (org) ->
          !!(org.get('name').indexOf(string) != -1)
    ).property 'filterString'
    updateFilter: (arg) ->
      @set 'filterString', arg