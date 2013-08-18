NewController = BaseController.extend
  createTransition: ->
    @transitionToRoute('discovery');
  create: ->
    @clearErrors()
    @recordProperties = @getProperties(@submitFields)
    if @fieldsPopulated()
      model = @get('model')
      record = model.createRecord(@recordProperties)
      record.transaction.commit()
      record.becameError =  =>
        @set 'errors', 'SERVER ERROR'
      record.didCreate = =>
        @createTransition()
    else
      @set 'errors', 'Empty Fields'