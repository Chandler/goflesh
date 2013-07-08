define ["ember", "ember-data"], (Em, DS) ->
  BaseController = Em.ObjectController.extend
    errors: null,
    clearErrors: ->
      @set 'errors', null
    errorMessages: (->
      @get 'errors'
    ).property 'errors' 
