define ["ember"], (Em) ->
  BaseController = Em.Controller.extend
    errors: null
    
    clearErrors: ->
      @set 'errors', null
    errorMessages: (->
      @get 'errors'
    ).property 'errors'

    fieldsPopulated: ->
      for k,v of @recordProperties
        return false if v == ''
      true
