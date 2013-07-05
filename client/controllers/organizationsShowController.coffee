define ["ember", "ember-data"], (Em, DS) ->
  OrganizationsShowController = Ember.ObjectController.extend
    orgname: (->
        debugger
        console.log "ohh"
      ).property('test')