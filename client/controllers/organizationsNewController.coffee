define ["ember", "NewController"], (Em, NewController) ->
  OrganizationsNewController = NewController.extend
    recordProperties: ['name', 'slug', 'location']
    name: ''
    slug: ''
    location: ''