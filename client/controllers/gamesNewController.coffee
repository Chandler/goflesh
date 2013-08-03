define ["ember", "ember-data", "NewController"], (Em, DS, NewController) ->
  GamesNewController = NewController.extend
    recordProperties: ['name', 'slug']
    name: '',
    slug: '',
