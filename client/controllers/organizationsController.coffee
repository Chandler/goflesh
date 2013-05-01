define ["ember"], (Em) ->
  OrganizationsController = Ember.ObjectController.extend

    go: ->
      model = this.get('model')
      