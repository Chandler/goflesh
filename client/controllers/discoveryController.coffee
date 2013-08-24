App.DiscoveryController = Ember.Controller.extend
  orgs: (->
    string = @get 'filterString'
    if string == ""
      return []
    else
      @get('content').filter (org) ->
        !!(org.get('name').indexOf(string) != -1)
  ).property 'filterString'
  updateFilter: (arg) ->
    @set 'filterString', arg