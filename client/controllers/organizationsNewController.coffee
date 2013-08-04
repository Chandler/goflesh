define ["NewController"], (NewController) ->
  OrganizationsNewController = NewController.extend
    submitFields: ['name', 'slug', 'location']
    name: ''
    slug: ''
    location: ''