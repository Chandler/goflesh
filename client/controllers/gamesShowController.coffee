define ["ember", "ember-data"], (Em, DS) ->
  GamessShowController = Ember.ObjectController.extend(setupController: (GamessShowController, Game) ->
  		UsersShowController.set "Game", Game
  	)