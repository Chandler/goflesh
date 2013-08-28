
# Custom controller stack
#
# BaseController
# NewController

BaseController = Ember.Controller.extend

  errors: null

  #list of fields on the record which you are going to expose in the 
  #template for editing/saving
  editableRecordFields: null,
  
  clearErrors: ->
    @set 'errors', null
  
  errorMessages: (->
    @get 'errors'
  ).property 'errors'

  clearMessages: ->
    @set 'messages', null

  fieldsPopulated: ->
    for k,v of @getRecordProperties()
      return false if !v
    true

  successTransition: ->
    @transitionToRoute('discovery');

  #returns an object with values from textFields on the page
  #useful for checking to see if the properties you care about (submitFields) 
  #are empty or not
  getRecordProperties: ->
    @recordToSave().getProperties(@editableRecordFields)

  recordToSave: ->
    @get('content')

  save:(update) ->
    @clearErrors()
    if @fieldsPopulated()
      console.log "Begining save record"
      record = @recordToSave()
      @get('store').get('defaultTransaction').commit()
      record.on 'becameError', =>
        @set 'errors', 'SERVER ERROR'
      record.on 'didUpdate', =>
        @successTransition()
    else
      @set 'errors', "Empty Fields"
  

NewController = BaseController.extend
  #@override
  recordToSave: ->
    @get('model').createRecord(@getRecordProperties())

  #we have no model object the values are saved right on the controller
  getRecordProperties: ->
    @getProperties(@editableRecordFields)

