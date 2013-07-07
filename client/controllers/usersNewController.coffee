define ["ember", "ember-data", "NewController"], (Em, DS, NewController) ->
  UsersNewController = NewController.extend
    first_name: '',
    last_name: '',
    email: '',
    screen_name: '',
    password: ''
    createUser: ->
      this.clearErrors()
      if this.first_name != ''
        model = this.get('model')
        record = model.createRecord
          first_name: this.first_name
          last_name: this.last_name
          email: this.email
          screen_name: this.screen_name
          password: this.password
        record.transaction.commit()
        record.becameError =  =>
          @set 'errors', 'SERVER ERROR'
        record.didCreate = =>
          @transitionToRoute('users.show', record.id);
      else
        @set 'errors', 'Empty Fields'
    errors: null,
    clearErrors: ->
      @set 'errors', null
    errorMessages: (->
      @get 'errors'
    ).property 'errors' 
