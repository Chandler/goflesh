define ["ember", "ember-data"], (Em, DS) ->
  OrganizationsShowController = Ember.ObjectController.extend
    orgname: (->
        console.log "ohh"
      ).property('test')