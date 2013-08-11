define ["ember"], (Em) ->
  BaseController = Em.Controller.extend
    errors: null
    messages: null

    clearErrors: ->
      @set 'errors', null
    errorMessages: (->
      @get 'errors'
    ).property 'errors'

    clearMessages: ->
      @set 'messages', null

    successMessages: (->
      @get 'messages'
      ).property 'messages'
    
    fieldsPopulated: ->
      for k,v of @recordProperties
        return false if v == ''
      true
