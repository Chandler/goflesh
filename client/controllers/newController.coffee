define ["ember", "ember-data"], (Em, DS) ->
  NewController = Ember.ObjectController.extend
    errors: null,
    clearErrors: ->
      @set 'errors', null
    errorMessages: (->
      @get 'errors'
    ).property 'errors' 
