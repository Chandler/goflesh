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
        record.becameError =  ->
          this.set 'errors', 'SERVER ERROR'
      else
        this.set 'errors', 'Empty Fields'
    errors: null,
    clearErrors: ->
      this.set 'errors', null
    errorMessages: (->
      this.get 'errors'
    ).property 'errors' 
