define ["ember"], (Em) ->
  BaseController = Em.ObjectController.extend
    errors: null,
    clearErrors: ->
      @set 'errors', null
    errorMessages: (->
      @get 'errors'
    ).property 'errors' 
