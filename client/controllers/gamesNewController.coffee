define ["NewController"], (NewController) ->
  GamesNewController = NewController.extend
    recordProperties: ['name', 'slug']
    name: '',
    slug: '',
