define ["ember", "templates", "ember-data"], (Em, Templates, DS) ->

  Store = DS.Store.extend
    revision: 11
    adapter: 'DS.fixtureAdapter'
    
  Game = DS.Model.extend()

  Game.FIXTURES = [
    id: 1
    id: 2
    id: 3
  ]