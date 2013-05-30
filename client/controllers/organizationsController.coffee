define ["ember", "ember-data"], (Em, DS) ->
  OrganizationsController = Ember.ObjectController.extend
    orgname: '',
    slug: '',
    location: '',
    go: ->
      if this.orgname != ''
        model = this.get('model')
        record = model.createRecord
          name: this.orgname
          slug: this.slug
        record.transaction.commit()
        
    hasErrors: (->
      
    ).property 'errorMessages' 
