define(function (require) {
  var app = require('durandal/app'),
      ko = require('knockout'),
      bootstrap = require('bootstrap');

  function SponsorsViewModel() {
      var self = this;
      self.sponsors = ko.observableArray();
      self.sposURI = '/patrocinadores';
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
      self.ajax(self.sposURI,'GET').done(function(data) {
          for (var i = 0; i < data.Length; i++) {                        
              self.suppliers.push({
                  id: ko.observable(data.Suppliers[i].ID),
                  name: ko.observable(data.Suppliers[i].Name),
                  phone: ko.observable(data.Suppliers[i].PhoneExt),
                  note: ko.observable(data.Suppliers[i].Note)
              });  
          }
      });
      self.beginAdd = function() {
          $('#add').modal('show');
      }
      self.add = function(sponsor){
          self.ajax(self.sposURI, 'POST', supplier).done(function(data) {
              self.suppliers.push({
                  id: ko.observable(data.ID),
                  name: ko.observable(data.Name),
                  phone: ko.observable(data.PhoneExt),
                  note: ko.observable(data.Note)
              });
          });
      }
      self.beginEdit = function(sponsor) {
          editSponsorViewModel.setSponsor(sponsor);
          $('#edit').modal('show');
      }
      self.edit = function(sponsor, data) {
          self.ajax(self.sposURI+'/'+sponsor.id(), 'PUT', data).done(function(res) {
              self.updateSponsor(sponsor, res);
          });
      }
      self.updateSponsor = function(sponsor, newSponsor) {
          var i = self.sponsors.indexOf(sponsor);
          self.sponsors()[i].id(newSponsor.ID);
          self.sponsors()[i].name(newSponsor.Name);
          self.sponsors()[i].phone(newSponsor.PhoneExt);
          self.sponsors()[i].note(newSponsor.Note);
      }
      self.remove = function(sponsor) {
          self.ajax(self.sposURI+'/'+sponsor.id(),'DELETE').done(function(){
              self.suppliers.remove(sponsor);
          });                    
      }              
  }
  
  function AddSponsorViewModel() {
      var self = this;
      self.name = ko.observable();
      self.phone = ko.observable();
      self.note = ko.observable();
      self.addSponsor = function() {
          $('#add').modal('hide');
          sponsorsViewModel.add({
              name: self.name(),
              note: self.note(),
              phone: self.phone()
          });
          self.name("");
          self.phone("");
          self.note("");
      }
  }
  
  function EditSponsorViewModel() {
      var self = this;
      self.name = ko.observable();
      self.phone = ko.observable();
      self.note = ko.observable();
      self.setSponsor = function(sponsor) {
          self.sponsor = sponsor;
          self.name(sponsor.name());
          self.phone(sponsor.phone());
          self.note(sponsor.note());
          $('edit').modal('show');
      }
      self.editSponsor = function() {
          $('#edit').modal('hide');
          sponsorrsViewModel.edit(self.sponsor, {
              name: self.name(),
              phone: self.phone() ,
              note: self.note()
          });
      }
  }
  
  var sponsorsViewModel = new SponsorsViewModel(); 
  var addSponsorViewModel = new AddSponsorViewModel();
  var editSponsorViewModel = new EditSponsorViewModel();
            
  return {
     sponsorsViewModel, 
     addSponsorViewModel, 
     editSponsorViewModel
   };
});
