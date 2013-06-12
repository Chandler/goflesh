define ["ember", "ember-data"], (Em, DS) ->
  OrganizationsController = Ember.ObjectController.extend
    orgname: '',
    slug: '',
    location: '',
    createOrg: ->
      this.clearErrors()
      if this.orgname != ''
        model = this.get('model')
        record = model.createRecord
          name: this.orgname
          slug: this.slug
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
