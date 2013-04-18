define ["ember", "templates", "ember-data", "app"], (Em, Templates, DS, App) ->

  App.Store = DS.Store.extend
    revision: 11,
    adapter: DS.FixtureAdapter.create()
    
  App.Game = DS.Model.extend
    name: DS.attr 'number'

  App.Game.FIXTURES = [
    id: 1
    id: 2
    id: 3
  ]

  App.Game