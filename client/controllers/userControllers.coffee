App.UserSettingsController = BaseController.extend
  needs: 'user'
  user: null
  userBinding: 'controllers.user'
  contentBinding: 'user'
  edit: ->
    @get('store').get('defaultTransaction').commit()
    @transitionToRoute('user.home')

App.UsersNewController = NewController.extend
  #last minute hack for registration will fix
  requiredFields: ['first_name', 'last_name', 'email', 'screen_name', 'password']
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
  ).property('content.id')

  currentPlayer: (->
    player = @get('players').get('lastObject')
    player
  ).property('players.@each')

App.UsersController = Ember.ObjectController.extend
  selectedUser: null

App.UserHomeController = Ember.Controller.extend
  needs: 'user'
  user: null
  userBinding: 'controllers.user'


