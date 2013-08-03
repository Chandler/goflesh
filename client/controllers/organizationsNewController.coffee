define ["NewController"], (NewController) ->
  OrganizationsNewController = NewController.extend
    recordProperties: ['name', 'slug', 'location']
    name: ''
    slug: ''
    location: ''