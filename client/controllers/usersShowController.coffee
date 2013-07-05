define ["ember", "ember-data"], (Em, DS) ->
  UsersShowController = Ember.ObjectController.extend(setupController: (UsersShowController, User) ->
  		UsersShowController.set "User", User
  	)