define ["ember", "ember-data"], (Em, DS) ->
  BaseController = Ember.ObjectController.extend
    errors: null,
    clearErrors: ->
      @set 'errors', null
    errorMessages: (->
      @get 'errors'
    ).property 'errors' 
