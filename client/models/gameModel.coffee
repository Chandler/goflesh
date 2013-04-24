define ["ember", "templates", "ember-data"], (Em, Templates, DS) ->

  Game = DS.Model.extend
    name: DS.attr 'string'

  Game.FIXTURES = [
    {
      id: 1,
      name: 'joe'
    },
    {
      id: 2
      name: 'kevin'
    }
  ]

  Game