define ["ember"], (Em) ->
  DiscoveryController = Ember.ObjectController.extend
    organizations: (->
      string = this.get 'filterString'
      if string == ""
        return this.get 'content'
      else
        this.get('content').filter (org) ->
          !!(org.get('name').indexOf(string) != -1)
    ).property 'filterString'
    updateFilter: (arg) ->
      this.set 'filterString', arg