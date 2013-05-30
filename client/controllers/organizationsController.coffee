define ["ember", "ember-data"], (Em, DS) ->
  OrganizationsController = Ember.ObjectController.extend

    go: ->
      model = this.get('model')
      record = model.createRecord
        name: "aaaaa"
        slug: "aslug"
      console.log record.get('name')
      record.transaction.commit()
      console.log model.find(1)