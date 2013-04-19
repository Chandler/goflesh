define ["ember", "templates", "ember-data"], (Em, Templates, DS) ->

  Game = DS.Model.extend
    name: DS.attr 'number'

  Game.FIXTURES = [
    id: 1
    id: 2
    id: 3
  ]

  Game