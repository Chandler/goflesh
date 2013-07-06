define ["ember", "ember-data"], (Em, DS) ->
  GamesShowController = Ember.ObjectController.extend(setupController: (GamesShowController, Game) ->
  		GamesShowController.set "Game", Game
  	)