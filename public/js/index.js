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
				$("#view-contact-body").append("<tr id='contact-row-" + contact.PartitionNumber + "'><td>"+contact.ContactName+"</td> <td><button class='btn btn-default' onclick='ViewDetails(\""+contact.ContactName+"\",\""+contact.PartitionNumber+"\")'>View</button></td><td><button class='btn btn-danger' onclick='DeleteContact(\""+contact.ContactName+"\",1,\""+contact.PartitionNumber+"\")'>Delete</button></td><td><a class='open-AddPhoneModal btn btn-danger' data-toggle='modal' href='#add-phone-modal' data-id=\""+contact.ContactName+"\">Add Number</a></td></tr>");

				}
			})
		}
	function AddNumber (){
	    $.ajax({
			url:"/index/addNumberToContact",
			method:"POST",
			data:$("#addnumber-form").serialize(),
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
					var row = $("#view-contact-body").append("<tr id='contact-row-" + result.PartitionNumber + "'><td>"+result.ContactName+"</td> <td><button class='btn btn-default' onclick='ViewDetails("+result.ContactName+"," + result.PartitionNumber + ")'>View</button></td><td><button class='btn btn-danger' onclick='DeleteContact(" + result.ContactName + ",1," + result.PartitionNumber + ")'>Delete</button></td><td><a class='open-AddPhoneModal btn btn-danger' data-toggle='modal' href='#add-phone-modal' data-id="+result.ContactName+">Add Number</a></td></tr>");
					searchResults.append(row);
				    });
				}
			})
	}
	function DeleteContact(contactname,number,id) {
        $.ajax({
          method: "POST",
          url: "/index/DeleteContact?contactname="+contactname+"&number="+number+"&id="+id+"",
          success: function() {
            $("#contact-row-" + id).remove();
            $("#details-row-" + id).remove();
          }
        });
      }
     function ViewDetails(contactname,partitionnumber) {
     $.ajax({
          method: "GET",
          url: "/index/viewContactDetails?contactname="+contactname+"&partitionnumber="+partitionnumber+"",
          success: function(rawData) {
                showViewDetails();
				var searchResults=$("#view-details-body");
				searchResults.empty();
				rawData.forEach(function(result){
					var row = $("#view-details-body").append("<tr id='details-row-" + result.PhoneID + "'><td>"+result.ContactName+"</td> <td>"+result.Phone+"</td><td>"+result.Email+"</td><td>"+result.Address+"</td><td>"+result.Nationality+"</td><td><button class='btn btn-danger' onclick='DeleteContact(\""+result.ContactName+"\",2,\""+result.PhoneID+"\")'>Delete</button></td></tr>");
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