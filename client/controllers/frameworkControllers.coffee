
# Custom controller stack
#
# BaseController
# NewController

BaseMixin = Ember.Mixin.create
  signedIn: ->
    App.Auth.signedIn

  errors: null

  #list of fields on the record which you are going to expose in the 
  #template for editing/saving
  editableRecordFields: null,
  
  clearErrors: ->
    @set 'errors', null
  
  errorMessages: (->
    @get 'errors'
  ).property 'errors'

  currentUser: (->
    App.Auth.get('user')
  ).property()

  signOut:  ->
    App.Auth.destroySession()

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

  edit: ->
    @get('store').get('defaultTransaction').commit()

  save: ->
    @clearErrors()
    if @fieldsPopulated()
      console.log "Begining save record"
      record = @recordToSave()
      @get('store').get('defaultTransaction').commit()
      record.on 'becameError', =>
        @set 'errors', 'SERVER ERROR'
      record.on 'didCreate', =>\
        @successTransition()
    else
      @set 'errors', "Empty Fields"

BaseController = Ember.Controller.extend(BaseMixin)
BaseObjectController = Ember.ObjectController.extend(BaseMixin)


NewController = BaseController.extend
  recordToSave: ->
    @get('model').createRecord(@getRecordProperties())

  #we have no model object the values are saved right on the controller
  getRecordProperties: ->
    @getProperties(@editableRecordFields)

