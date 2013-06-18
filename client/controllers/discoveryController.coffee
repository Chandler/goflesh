define ["ember"], (Em) ->
  DiscoveryController = Ember.ObjectController.extend
    orgs: (->
      string = @get 'filterString'
      if string == ""
        return @get 'model'
      else
        @get('model').filter (org) ->
          !!(org.get('name').indexOf(string) != -1)
    ).property 'filterString'
    updateFilter: (arg) ->
      @set 'filterString', arg