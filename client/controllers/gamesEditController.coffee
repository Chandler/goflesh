define ["ember"], (Em) ->
  GamesEditController = Em.ObjectController.extend
    editGame: ->
      this.clearErrors()
      if this.name != ''
        record = @get('model')
        record.setProperties
          name: @get("name")
          slug: @get("slug")
        record.transaction.commit()
        record.becameError =  =>
          @set 'errors', 'SERVER ERROR'
        record.didUpdate = =>
          @transitionToRoute('game.show', record);
      else
        @set 'errors', 'Empty Fields'
    errors: null,
    clearErrors: ->
      @set 'errors', null
    errorMessages: (->
      @get 'errors'
    ).property 'errors' 