define ["ember", "ember-data", "NewController"], (Em, DS, NewController) ->
  OrganizationsNewController = Em.ObjectController.extend
    name: '',
    slug: '',
    location: '',
    create: ->
      this.clearErrors()
      if this.name != ''
        model = this.get('model')
        record = model.createRecord
          name: this.name
          slug: this.slug
          location: this.location
        record.transaction.commit()
        record.becameError =  =>
          @set 'errors', 'SERVER ERROR'
        record.didCreate = =>
          @transitionToRoute('discovery');
      else
        @set 'errors', 'Empty Fields'
    errors: null,
    clearErrors: ->
      @set 'errors', null
    errorMessages: (->
      @get 'errors'
    ).property 'errors' 