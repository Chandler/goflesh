define ["ember", "ember-data", "NewController"], (Em, DS, NewController) ->
  OrganizationsNewController = NewController.extend
    fields:
      name: '',
      slug: '',
      location: '',