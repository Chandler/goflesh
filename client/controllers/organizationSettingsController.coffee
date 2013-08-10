define ["ember", "BaseController"], (Em, BaseController) ->
  OrganizationSettingsController = BaseController.extend
    needs: 'organization'
    organization: null
    organizationBinding: 'controllers.organization'
    editOrg: ->
      this.clearErrors()
      if @get('organization.name') != ''
        record = @get('organization')
        record.setProperties
          name: @get("organization.name")
        record.transaction.commit()
        record.becameError =  =>
          @set 'errors', 'SERVER ERROR'
        record.didUpdate = =>
          @transitionTo('organization.home', record);
      else
        @set 'errors', 'Empty Field'