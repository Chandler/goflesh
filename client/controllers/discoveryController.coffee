define ["ember"], (Em) ->
  DiscoveryController = Ember.ObjectController.extend
    organizations: (->
      console.log 'organizations'
      return this.get('orgs')
    ).property("orgs.@each")
    something: ->
      neworgs = this.get('model').filter (org) ->
        if (!org.get('show'))
          true
      console.log('something')
      this.set('orgs', [])