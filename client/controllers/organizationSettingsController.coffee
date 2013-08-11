define ["ember", "BaseController", "ember-data"], (Em, BaseController) ->
  OrganizationSettingsController = BaseController.extend
    needs: 'organization'
    organization: null
    organizationBinding: 'controllers.organization'
    edit: ->
      this.clearErrors()
      if @get('organization.name') != ''
        record = @get('organization.content')
        @get('store').get('defaultTransaction').commit()
        record.on 'becameError', =>
          @set 'errors', 'SERVER ERROR'
        record.on 'didUpdate', =>
          @transitionTo('organization.home', record);
      else
        @set 'errors', 'Empty Field'