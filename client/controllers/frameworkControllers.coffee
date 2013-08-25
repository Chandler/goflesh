
# Custom controller stack
#
# BaseController
# NewController

BaseController = Ember.Controller.extend
  upload: ->
    token = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJwcml2YXRlX3Rva2VuIjoiYjdhZGFkNjI2YWI0NTc0ZmMwYTE5M2MyOWQ0YjBiMTkxYjQxYWNlOWQwMzUzNTU5YzhhMzRhMDhkOWRiNGI2ZCJ9.V0JQ5FygLaWiXn1yzN-gdmywkki6V1V8r2y1TqSMPUU"
    client = new AvatarsIO(prompt('Enter your public token'));
    uploader = client.create($('#upload'))
    uploader.setAllowedExtensions ["png", "jpg"]
    uploader.setIdentifier "4"
    
    uploader.on "complete", (url) ->
      console.log url
      
    uploader.on "error", (err) ->
      console.log "error happened during request"

  
  errors: null

  #list of fields on the record which you are going to expose in the 
  #template for editing/saving
  editableRecordFields: null,
  
  clearErrors: ->
    @set 'errors', null
  
  errorMessages: (->
    @get 'errors'
  ).property 'errors'


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

