define ["BaseController"], (BaseController) ->
  NewController = Em.ObjectController.extend
    create: ->
      @clearErrors()
      if @name != ''
        model = @get('model')
        record = model.createRecord(@getProperties(@recordProperties))
        record.transaction.commit()
        record.becameError =  =>
          @set 'errors', 'SERVER ERROR'
        record.didCreate = =>
          @transitionToRoute('discovery');
      else
        @set 'errors', 'Empty Fields'