require.config({
  baseUrl: "public/js",
  packages: [
    {
      name: "handlebars",
      location: "lib",
      main: "handlebars.js"
    }, {
      name: "ember",
      location: "lib",
      main: "ember.js"
    }, {
      name: "ember-data",
      location: "lib",
      main: "ember-data.js"
    }
  ]
});

require(["jquery", "app", "gameModel", "templates"], function($, App, GameModel, Templates) {
  var view;

  App.Router.map(function() {
    return this.route('discovery');
  });
  App.IndexRoute = Ember.Route.extend({
    redirect: function() {
      return this.transitionTo("discovery");
    }
  });
  App.ApplicationView = Ember.View.extend({
    template: 'willy wonka'
  });
  App.ApplicationController = Ember.Controller.extend();
  App.DiscoveryRouter = Ember.Route.extend({
    model: function() {
      return GameModel.find();
    },
    setupController: function(controller, model) {
      return controller.set('content', model);
    },
    renderTemplate: function() {
      return render('disovery');
    }
  });
  App.DiscoveryController = Em.ObjectController.extend({
    message: 'This is the discovery page'
  });
  App.DiscoveryView = Ember.View.extend({
    template: function(name, data) {
      return Templates["discovery"]({
        message: "hey"
      });
    }
  });
  view = App.DiscoveryView.create({
    controller: 'discovery'
  });
  return view.append("#app");
});
