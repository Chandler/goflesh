
  # editableRecordFields: ['name', 'slug', 'location']
  # name: ''
  # slug: ''
  # location: ''

App.OrganizationHomeController = Ember.Controller.extend
  needs: 'organization'
  organization: null
  organizationBinding: 'controllers.organization'
  contentBinding: 'organization.games'

App.OrganizationsController = Ember.ObjectController.extend
  selectedOrganization: null

App.OrganizationSettingsController = BaseController.extend
  editableRecordFields: ['name', 'slug']
  needs: 'organization'
  organizationBinding: 'controllers.organization'
  contentBinding: 'controllers.organization.content'