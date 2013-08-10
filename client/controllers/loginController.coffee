define ["ember"], (Em) ->
  LoginController = Em.ObjectController.extend
    email: ''
    password: ''
    login: (arg) ->
      @clearErrors()
      Em.App.Auth.signIn
        data:
          email: this.email
          password: this.password
    errors: null,
    clearErrors: ->
      @set 'errors', null
    errorMessages: (->
      @get 'errors'
    ).property 'errors' 
