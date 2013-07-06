define ["ember", "ember-data"], (Em, DS) ->
  OrganizationsShowController = Ember.ObjectController.extend(setupController: (OrganizationsShowController, Organization) ->
  		OrganizationsShowController.set "Organization", Organization
  	)