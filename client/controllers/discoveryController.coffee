define ["ember"], (Em) ->
  DiscoveryController = Em.ObjectController.extend
    orgs: (->
      string = @get 'filterString'
      if string == ""
        return []
      else
        @get('model').filter (org) ->
          !!(org.get('name').indexOf(string) != -1)
    ).property 'filterString'
    updateFilter: (arg) ->
      @set 'filterString', arg