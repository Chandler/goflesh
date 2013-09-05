BaseMixin = Ember.Mixin.create
  signedIn: ->
    App.Auth.signedIn

  currentUser: (->
    App.Auth.get('user')
  ).property()

  signOut:  ->
    App.Auth.destroySession()
    location.reload()

  errors: null

  clearErrors: ->
    @set 'errors', null
  
  errorMessages: (->
    @get 'errors'
  ).property 'errors'

  #list of fields on the record which you are going to expose in the 
  #template for editing/saving
  editableRecordFields: null,
  requiredFields: null,
  
  fieldsPopulated: ->
    for k,v of @getRequiredRecordProperties()
      return false if !v
    true

  successTransition: ->
    @transitionToRoute('discovery');


  getRequiredRecordProperties: ->
    fields = @requiredFields || @editableRecordFields
    @recordToSave().getProperties(fields)


  #returns an object with values from textFields on the page
  #useful for checking to see if the properties you care about (submitFields) 
  #are empty or not
  getRecordProperties: ->
    @recordToSave().getProperties(@editableRecordFields)

  recordToSave: ->
    @get('content')

  #assumes edit forms were bound to model, need to commit the changes
  edit: ->
    @get('store').get('defaultTransaction').commit()

  save: ->
    @clearErrors()
    if @fieldsPopulated()
      console.log "Begining save record"
      record = @recordToSave()
      record.save()
      record.one 'didError', =>
        @set 'errors', e.get('errors')
      record.one 'becameError', =>
        @set 'errors', e.get('errors')
      record.one 'becameInvalid', (e) =>
        @set 'errors', e.get('errors')
      record.one 'didCreate', =>        
        @successTransition()
    else
      @set 'errors', "Please fill in all fields"

NewMixin = Ember.Mixin.create
  successTransition: ->
    @transitionToRoute('login');
  recordToSave: ->
    @get('model').createRecord(@getRecordProperties())

  #we have no model object the values are saved right on the controller
  getRecordProperties: ->
    @getProperties(@editableRecordFields)

BaseController = Ember.Controller.extend(BaseMixin)

BaseObjectController = Ember.ObjectController.extend(BaseMixin)

NewController = BaseController.extend(NewMixin)


