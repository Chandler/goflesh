define ["NewController"], (NewController) ->
  GamesNewController = NewController.extend
    submitFields: ['name', 'slug']
    name: '',
    slug: '',
