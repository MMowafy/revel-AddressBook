function showContactList(){
		$("#contact-list-container").show();
		$("#add-contact-container").hide();
		$("#view-details-container").hide();
		}
	function showAddContact(){
		$("#contact-list-container").hide();
		$("#view-details-container").hide();
		$("#add-contact-container").show();
		}
	function showViewDetails(){
	    $("#contact-list-container").hide();
		$("#add-contact-container").hide();
		$("#view-details-container").show();
	}
	function AddToContacts() {
		$.ajax({
			url:"/index/addNewContact",
			method:"POST",
			data:$("#addcontact-form").serialize(),
			success: function(data) {
				var contact=data;
				$("#view-contact-body").append("<tr id='contact-row-" + contact.PK + "'><td>"+contact.ContactName+"</td> <td><button class='btn btn-default' onclick='ViewDetails("+contact.PK+")'>View</button></td><td><button class='btn btn-danger' onclick='DeleteContact(" + contact.PK + ",1)'>Delete</button></td><td><a class='open-AddPhoneModal btn btn-danger' data-toggle='modal' href='#add-phone-modal' data-id="+contact.PK+">Add Number</a></td></tr>");
				}
			})
		}
	function AddNumber (){
	    $.ajax({
			url:"/index/addNumberToContact",
			method:"POST",
			data:$("#addnumber-form").serialize(),
			success: function(data) {
				var contact=JSON.parse(data);
				if(!contact) return;
				$("#view-contact-body").append("<tr id='contact-row-" + contact.PK + "'><td>"+contact.ContactName+"</td> <td><button class='btn btn-default' onclick='ViewDetails("+contact.PK+")'>View</button></td><td><button class='btn btn-danger' onclick='DeleteContact(" + contact.PK + ",1)'>Delete</button></td><td><a class='open-AddPhoneModal btn btn-danger' data-toggle='modal' href='#add-phone-modal' data-id="+contact.PK+">Add Number</a></td></tr>");
				}
			})
	}
	function SortContacts (columnName) {
	    $.ajax({
			url:"/index/SortContacts?sortBy="+columnName,
			method:"GET",
			success: function(rawData) {
				var searchResults=$("#view-contact-body");
				searchResults.empty();
				rawData.forEach(function(result){
					var row = $("#view-contact-body").append("<tr id='contact-row-" + result.PK + "'><td>"+result.ContactName+"</td> <td><button class='btn btn-default' onclick='ViewDetails("+result.PK+")'>View</button></td><td><button class='btn btn-danger' onclick='DeleteContact(" + result.PK + ",1)'>Delete</button></td><td><a class='open-AddPhoneModal btn btn-danger' data-toggle='modal' href='#add-phone-modal' data-id="+result.PK+">Add Number</a></td></tr>");
					searchResults.append(row);
				    });
				}
			})
	}
	function DeleteContact(pk,number) {
        $.ajax({
          method: "POST",
          url: "/index/DeleteContact?pk="+pk+"&number="+number+"",
          success: function() {
            $("#contact-row-" + pk).remove();
            $("#details-row-" + pk).remove();
          }
        });
      }
     function ViewDetails(pk) {
     $.ajax({
          method: "GET",
          url: "/index/viewContactDetails?pk=" + pk,
          success: function(rawData) {
                showViewDetails();
				var searchResults=$("#view-details-body");
				searchResults.empty();
				rawData.forEach(function(result){
					var row = $("#view-details-body").append("<tr id='details-row-" + result.PK + "'><td>"+result.ContactName+"</td> <td>"+result.Phone+"</td><td>"+result.Email+"</td><td>"+result.Address+"</td><td>"+result.Nationality+"</td><td><button class='btn btn-danger' onclick='DeleteContact(" + result.PK + ",2)'>Delete</button></td></tr>");
					searchResults.append(row);
				    });
				}
        })
     }
$(document).on("click", ".open-AddPhoneModal", function () {
     var myContactId = $(this).data('id');
     var myContactName = $(this).data('nametitle');
     $(".modal-body #contactID").val( myContactId );
     $(".modal-header #contactnameTitle").text( "New Number for "+myContactName+"");
    $('#add-phone-modal').modal('show');
});