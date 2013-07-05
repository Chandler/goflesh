define ["ember", "ember-data"], (Em, DS) ->
  OrganizationsShowController = Ember.ObjectController.extend
    name: (->
        console.log "ohh"
      ).property('test')