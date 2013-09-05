App.UserSettingsController = BaseController.extend
  needs: 'user'
  user: null
  userBinding: 'controllers.user'
  contentBinding: 'user'
  edit: ->
    @_super
    @transitionTo 'user.home'

App.UsersNewController = NewController.extend
  editableRecordFields: ['first_name', 'last_name', 'email', 'screen_name', 'password', 'phone']
  first_name: '',
  last_name: '',
  email: '',
  screen_name: '',
  password: ''
  phone: ''

App.UserController = BaseObjectController.extend
  userIsCurrentUser: (->
    id = App.Auth.get('user.id')
    return (id && id == @get('content.id'))
  ).property()

App.UsersController = Ember.ObjectController.extend
  selectedUser: null

App.UserHomeController = Ember.Controller.extend
  needs: 'user'
  user: null
  userBinding: 'controllers.user'


