this["JST"] = this["JST"] || {};

this["JST"]["application"] = Handlebars.template(function (Handlebars,depth0,helpers,partials,data) {
  this.compilerInfo = [2,'>= 1.0.0-rc.3'];
helpers = helpers || Handlebars.helpers; data = data || {};
  var stack1, functionType="function", escapeExpression=this.escapeExpression;


  if (stack1 = helpers.sup) { stack1 = stack1.call(depth0, {hash:{},data:data}); }
  else { stack1 = depth0.sup; stack1 = typeof stack1 === functionType ? stack1.apply(depth0) : stack1; }
  return escapeExpression(stack1);
  });

this["JST"]["home"] = Handlebars.template(function (Handlebars,depth0,helpers,partials,data) {
  this.compilerInfo = [2,'>= 1.0.0-rc.3'];
helpers = helpers || Handlebars.helpers; data = data || {};
  var stack1, functionType="function", escapeExpression=this.escapeExpression;


  if (stack1 = helpers.hey) { stack1 = stack1.call(depth0, {hash:{},data:data}); }
  else { stack1 = depth0.hey; stack1 = typeof stack1 === functionType ? stack1.apply(depth0) : stack1; }
  return escapeExpression(stack1);
  });