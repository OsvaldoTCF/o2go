define(function (require) {
  var app = require('durandal/app'),
      ko = require('knockout'),
      bootstrap = require('bootstrap'),
      jquery = require('jquery'),
      router = require('plugins/router');

  var emailsViewModel = function(){
      var self = this;
      self.emails = ko.observableArray();
      self.supplier = ko.observable();
      self.supid = ko.observable();
      self.emlsURI = '/emailfornecedor';
  
      self.ajax = function(uri, method, data) {
          var request = {
              url: uri,
              type: method,
              cache: false,
              dataType: 'json',
              contentType: "application/json",
              accepts: "application/json",
              crossDomain: true,
              data: JSON.stringify(data),
              error: function(jqXHR) {
                  console.log("ajax error " + jqXHR.status);
              }
          };
          return $.ajax(request);
      }
      
      param = router.activeInstruction().params[0]
      
      if (typeof (param) !== 'undefined') {
          self.ajax(self.emlsURI+"?supplier="+param,'GET').done(function(data) {
              self.supplier(data.Supplier);
              self.supid(param);
              for (var i = 0; i < data.Length; i++) {                        
                  self.emails.push({
                      id: ko.observable(data.Emails[i].ID),
                      email: ko.observable(data.Emails[i].Email),
                      supp: ko.observable(data.Emails[i].SupplierID),
                      saved: ko.observable(true)
                  });  
              }
          });
      }               
      
      self.beginAdd = function() {
          self.emails.push({
              id: ko.observable(0),
              email: ko.observable(""),
              supp: ko.observable(self.supid()),
              saved: ko.observable(false)
          });
      }
      
      self.add = function(email){
          var data = {email:email.email(),supplier:email.supp()};
          self.doAdd(email,data);
      }
      
      self.cancel = function(email){
          self.emails.remove(email);
      }
      
      self.doAdd= function(email, data){
          self.ajax(self.emlsURI, 'POST', data).done(function(data) {
              self.emails.remove(email);
              self.emails.push({
                  id: ko.observable(data.ID),
                  email: ko.observable(data.Email),
                  supp: ko.observable(data.SupplierID),
                  saved: ko.observable(true)
              });
              app.showMessage("Email Added");
          });          
      }
  
      self.beginEdit = function(email) {
          var eml = ko.observable(email);
          var newmail = ko.observable(email.email());
          self.edit(eml,{email:newmail()})
      }
  
      self.edit = function(newmail) {               
          self.ajax(self.emlsURI+'/'+newmail.id(), 'PUT', {email:newmail.email()}).done(function(res) {
              self.updateEmail(newmail, res);
              app.showMessage("Email Changed");
          });
      }
      
      self.updateEmail = function(email, newEmail) {
          var i = self.emails.indexOf(email);
          self.emails()[i].id(newEmail.ID);
          self.emails()[i].email(newEmail.Email);
          self.emails()[i].supp(newEmail.SupplierID);
          self.emails()[i].saved(true);
      }
  
      self.remove = function(email) {
          self.ajax(self.emlsURI+'/'+email.id(),'DELETE').done(function(){
              self.emails.remove(email);
          });                    
      }             
  }
           
  return{
    emailsViewModel
  };
});
