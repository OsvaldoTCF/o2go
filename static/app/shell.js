define(function (require) {
  var router = require('plugins/router'),
      jquery = require('jquery');
 
  return {
     router: router,
     activate: function () {
       router.map([
         { route: '', title:'Home', moduleId: 'home', nav: true },
         { route: 'supplier', title:'Supplier', moduleId: 'supplier', nav: true },
         { route: 'sponsor', title:'Sponsor', moduleId: 'sponsor', nav: true },
         { route: 'emailsupplier(/:id)', title:'Email', moduleId:'email_supplier', nav :false},
         { route: 'emailsponsor(/:id)', title:'Email', moduleId:'email_sponsor', nav :false}
       ]).buildNavigationModel();
       
     router.activeItem.settings.areSameItem = function (currentItem, newItem, currentActivationData, newActivationData) {    
        return false;
     };
 
       return router.activate();
     }
   };
});
