
App.OrganizationHomeController = Ember.Controller.extend
  needs: 'organization'
  organization: null
  organizationBinding: 'controllers.organization'
  contentBinding: 'organization.games'
  joinGame: (game) ->
    App.Player.createRecord
      game: game
      user: App.Auth.get('user')
    @get('store').get('defaultTransaction').commit()

# App.OrganizationsController = Ember.ObjectController.extend
#   selectedOrganization: null

App.OrganizationSettingsController = BaseController.extend
  editableRecordFields: ['name', 'slug']
  needs: 'organization'
  organizationBinding: 'controllers.organization'
  contentBinding: 'controllers.organization.content'