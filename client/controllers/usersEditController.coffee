define ["ember", "ember-data"], (Em, DS) ->
  UsersEditController = Em.ObjectController.extend
    editUser: ->
      this.clearErrors()
      if this.first_name != ''
        record = @get('model')
        record.setProperties
          first_name: @get("first_name")
          last_name: @get("last_name")
          email: @get("email")
          screen_name: @get("screen_name")
          password: @get("password")
        record.transaction.commit()
        record.becameError =  =>
          console.log('noo')
          @set 'errors', 'SERVER ERROR'
        record.didUpdate = =>
          console.log('yay')
          @transitionToRoute('users.show', record);
      else
        @set 'errors', 'Empty Fields'
    errors: null,
    clearErrors: ->
      @set 'errors', null
    errorMessages: (->
      @get 'errors'
    ).property 'errors' 

