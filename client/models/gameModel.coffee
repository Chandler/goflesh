define ["ember", "templates", "ember-data"], (Em, Templates, Ed) ->
  debugger
  Store = DS.Store.extend ->
    revision: 11
    adapter: 'DS.asdf'
    
  Game = DS.Model.extend

  Game.FIXTURES = [
    id: 1
    id: 2
    id: 3
  ]