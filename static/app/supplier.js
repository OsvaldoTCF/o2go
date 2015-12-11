define(function (require) {
  var app = require('durandal/app'),
      ko = require('knockout'),
      bootstrap = require('bootstrap');

  var suppliersViewModel = function() {
      var self = this;
      self.suppliers = ko.observableArray();
      self.supsURI = '/fornecedores';
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
      self.ajax(self.supsURI,'GET').done(function(data) {
          for (var i = 0; i < data.Length; i++) {                        
              self.suppliers.push({
                  id: ko.observable(data.Suppliers[i].ID),
                  name: ko.observable(data.Suppliers[i].Name),
                  phone: ko.observable(data.Suppliers[i].PhoneNumber),
                  note: ko.observable(data.Suppliers[i].Note),
                  emails: ko.observableArray(data.Suppliers[i].Emails)
              });  
          }
      });
      self.beginAdd = function() {
          $('#add').modal('show');
      }
      self.add = function(supplier){
          self.ajax(self.supsURI, 'POST', supplier).done(function(data) {
              self.suppliers.push({
                  id: ko.observable(data.ID),
                  name: ko.observable(data.Name),
                  phone: ko.observable(data.PhoneNumber),
                  note: ko.observable(data.Note)
              });
          });
      }
      self.beginEdit = function(supplier) {
          editSupplierViewModel.setSupplier(supplier);
          $('#edit').modal('show');
      }
      self.edit = function(supplier, data) {
          self.ajax(self.supsURI+'/'+supplier.id(), 'PUT', data).done(function(res) {
              self.updateSupplier(supplier, res);
          });
      }
      self.updateSupplier = function(supplier, newSupplier) {
          var i = self.suppliers.indexOf(supplier);
          self.suppliers()[i].id(newSupplier.ID);
          self.suppliers()[i].name(newSupplier.Name);
          self.suppliers()[i].phone(newSupplier.PhoneNumber);
          self.suppliers()[i].note(newSupplier.Note);
      }
      self.remove = function(supplier) {
          self.ajax(self.supsURI+'/'+supplier.id(),'DELETE').done(function(){
              self.suppliers.remove(supplier);
          });                    
      }              
  }
  
  var addSupplierViewModel = function() {
      var self = this;
      self.name = ko.observable();
      self.phone = ko.observable();
      self.note = ko.observable();
      self.addSupplier = function() {
          $('#add').modal('hide');
          suppliersViewModel.add({
              name: self.name(),
              note: self.note(),
              phone: self.phone()
          });
          self.name("");
          self.phone("");
          self.note("");
      }
  }
  
  var editSupplierViewModel = function() {
      var self = this;
      self.name = ko.observable();
      self.phone = ko.observable();
      self.note = ko.observable();
      self.setSupplier = function(supplier) {
          self.supplier = supplier;
          self.name(supplier.name());
          self.phone(supplier.phone());
          self.note(supplier.note());
          $('edit').modal('show');
      }
      self.editSupplier = function() {
          $('#edit').modal('hide');
          suppliersViewModel.edit(self.supplier, {
              name: self.name(),
              phone: self.phone() ,
              note: self.note()
          });
      }
  }
            
  return {
     suppliersViewModel, 
     addSupplierViewModel, 
     editSupplierViewModel
   };
});
